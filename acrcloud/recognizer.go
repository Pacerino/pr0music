package acrcloud

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lacrcloud_extr_tool -lpthread
#include <stdio.h>
#include <stdlib.h>
#include "dll_acr_extr_tool.h"
*/
import "C"

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
	"unsafe"
)

const (
	ACR_OPT_REC_AUDIO   string = "audio"
	ACR_OPT_REC_HUMMING string = "humming"
	ACR_OPT_REC_BOTH    string = "both"
)

const (
	ACR_ERR_CODE_OK         int = 0
	ACR_ERR_CODE_HTTP_ERR   int = 3000
	ACR_ERR_CODE_NO_RESULT  int = 1001
	ACR_ERR_CODE_JSON_ERR   int = 2002
	ACR_ERR_CODE_DECODE_ERR int = 2004
	ACR_ERR_CODE_MUTE_ERR   int = 2006
	ACR_ERR_CODE_GEN_FP_ERR int = 2008
	ACR_ERR_CODE_PARAM_ERR  int = 2009
)

type Recognizer struct {
	Host         string
	AccessKey    string
	AccessSecret string
	RecType      string
	TimeoutS     int
	HttpClient_  *http.Client
}

func NewRecognizer(configs map[string]string) *Recognizer {
	var result = new(Recognizer)
	result.Host = configs["host"]
	result.AccessKey = configs["access_key"]
	result.AccessSecret = configs["access_secret"]
	result.RecType = configs["recognize_type"]
	result.TimeoutS = 10

	result.HttpClient_ = &http.Client{
		Timeout: time.Duration(result.TimeoutS * int(time.Second)),
	}

	C.acr_init()
	return result
}

func (self *Recognizer) Post(url string, fieldParams map[string]string, fileParams map[string][]byte, timeoutS int) (string, int, error) {
	postDataBuffer := bytes.Buffer{}
	mpWriter := multipart.NewWriter(&postDataBuffer)

	for key, val := range fieldParams {
		_ = mpWriter.WriteField(key, val)
	}

	for key, val := range fileParams {
		fw, err := mpWriter.CreateFormFile(key, key)
		if err != nil {
			mpWriter.Close()
			return "", ACR_ERR_CODE_HTTP_ERR, fmt.Errorf("Create Form File Error: %v", err)
		}
		fw.Write(val)
	}

	mpWriter.Close()

	//hClient := &http.Client {
	//    Timeout: time.Duration(timeoutS * int(time.Second)),
	//}

	req, err := http.NewRequest("POST", url, &postDataBuffer)
	if err != nil {
		return "", ACR_ERR_CODE_HTTP_ERR, fmt.Errorf("NewRequest Error: %v", err)
	}
	req.Header.Set("Content-Type", mpWriter.FormDataContentType())

	response, err := self.HttpClient_.Do(req)
	if err != nil {
		return "", ACR_ERR_CODE_HTTP_ERR, fmt.Errorf("Http Client Do Error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", ACR_ERR_CODE_HTTP_ERR, fmt.Errorf("Http Response Status Code Is Not %d: %d", http.StatusOK, response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", ACR_ERR_CODE_HTTP_ERR, fmt.Errorf("Read From Http Response Error: %v", err)
	}

	return string(body), ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) GetSign(str string, key string) string {
	hmacHandler := hmac.New(sha1.New, []byte(key))
	hmacHandler.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hmacHandler.Sum(nil))
}

func (self *Recognizer) CreateHummingFingerprint(pcmData []byte) ([]byte, int, error) {
	if pcmData == nil || len(pcmData) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter pcmData is nil or len(pcmData) == 0")
	}

	var fp *C.char
	fpLenC := C.create_humming_fingerprint((*C.char)(unsafe.Pointer(&pcmData[0])), C.int(len(pcmData)), &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		return nil, ACR_ERR_CODE_GEN_FP_ERR, fmt.Errorf("Can not Create Humming Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) CreateAudioFingerprint(pcmData []byte) ([]byte, int, error) {
	if pcmData == nil || len(pcmData) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter pcmData is nil or len(pcmData) == 0")
	}

	var fp *C.char
	fpLenC := C.create_fingerprint((*C.char)(unsafe.Pointer(&pcmData[0])), C.int(len(pcmData)), 0, &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		return nil, ACR_ERR_CODE_MUTE_ERR, fmt.Errorf("Can not Create Audio Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) CreateHummingFingerprintByBuffer(fileBufferData []byte, startSeconds int, lenSeconds int) ([]byte, int, error) {
	if fileBufferData == nil || len(fileBufferData) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter fileBufferData is nil or len(fileBufferData) == 0")
	}

	var fp *C.char
	fpLenC := C.create_humming_fingerprint_by_filebuffer((*C.char)(unsafe.Pointer(&fileBufferData[0])), C.int(len(fileBufferData)), C.int(startSeconds), C.int(lenSeconds), &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		if fpLen == -1 {
			return nil, ACR_ERR_CODE_DECODE_ERR, fmt.Errorf("Can not Decode the Audio File.")
		}
		return nil, ACR_ERR_CODE_GEN_FP_ERR, fmt.Errorf("Can not Create Humming Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) CreateAudioFingerprintByBuffer(fileBufferData []byte, startSeconds int, lenSeconds int) ([]byte, int, error) {
	if fileBufferData == nil || len(fileBufferData) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter fileBufferData is nil or len(fileBufferData) == 0")
	}

	var fp *C.char
	fpLenC := C.create_fingerprint_by_filebuffer((*C.char)(unsafe.Pointer(&fileBufferData[0])), C.int(len(fileBufferData)), C.int(startSeconds), C.int(lenSeconds), 0, &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		if fpLen == -1 {
			return nil, ACR_ERR_CODE_DECODE_ERR, fmt.Errorf("Can not Decode the Audio File.")
		}
		return nil, ACR_ERR_CODE_MUTE_ERR, fmt.Errorf("Can not Create Audio Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) CreateAudioFingerprintByFpBuffer(fpBufferData []byte, startSeconds int, lenSeconds int) ([]byte, int, error) {
	if fpBufferData == nil || len(fpBufferData) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter fileBufferData is nil or len(fileBufferData) == 0")
	}

	var fp *C.char
	fpLenC := C.create_fingerprint_by_fpbuffer((*C.char)(unsafe.Pointer(&fpBufferData[0])), C.int(len(fpBufferData)), C.int(startSeconds), C.int(lenSeconds), &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		if fpLen == -1 {
			return nil, ACR_ERR_CODE_DECODE_ERR, fmt.Errorf("Can not Decode the Audio File.")
		}
		return nil, ACR_ERR_CODE_MUTE_ERR, fmt.Errorf("Can not Create Audio Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) CreateHummingFingerprintByFile(filePath string, startSeconds int, lenSeconds int) ([]byte, int, error) {
	if len(filePath) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter len(filePath) == 0")
	}

	var fp *C.char
	fpLenC := C.create_humming_fingerprint_by_file(C.CString(filePath), C.int(startSeconds), C.int(lenSeconds), &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		if fpLen == -1 {
			return nil, ACR_ERR_CODE_DECODE_ERR, fmt.Errorf("Can not Decode the Audio File.")
		}
		return nil, ACR_ERR_CODE_GEN_FP_ERR, fmt.Errorf("Can not Create Humming Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) CreateAudioFingerprintByFile(filePath string, startSeconds int, lenSeconds int) ([]byte, int, error) {
	if len(filePath) == 0 {
		return nil, ACR_ERR_CODE_PARAM_ERR, fmt.Errorf("Parameter len(filePath) == 0")
	}

	var fp *C.char
	fpLenC := C.create_fingerprint_by_file(C.CString(filePath), C.int(startSeconds), C.int(lenSeconds), 0, &fp)
	fpLen := int(fpLenC)
	if fpLen <= 0 {
		if fpLen == -1 {
			return nil, ACR_ERR_CODE_DECODE_ERR, fmt.Errorf("Can not Decode the Audio File.")
		}
		return nil, ACR_ERR_CODE_MUTE_ERR, fmt.Errorf("Can not Create Audio Fingerprint")
	}

	fpBytes := C.GoBytes(unsafe.Pointer(fp), C.int(fpLen))
	C.acr_free(fp)

	return fpBytes, ACR_ERR_CODE_OK, nil
}

func (self *Recognizer) GetDurationMsByFile(filePath string) (int, error) {
	if len(filePath) == 0 {
		return 0, fmt.Errorf("Parameter pcmData is nil or len(pcmData) == 0")
	}

	duration := C.get_duration_ms_by_file(C.CString(filePath))

	return int(duration), nil
}

func (self *Recognizer) GetDurationMsByFpBuffer(fpBufferData []byte) (int, error) {
	if len(fpBufferData) == 0 {
		return 0, fmt.Errorf("Parameter pcmData is nil or len(pcmData) == 0")
	}

	duration := C.get_duration_ms_by_fpbuffer((*C.char)(unsafe.Pointer(&fpBufferData[0])), C.int(len(fpBufferData)))

	return int(duration), nil
}

func (self *Recognizer) DoRecognize(audioFp []byte, humFp []byte, userParams map[string]string) (string, int, error) {
	qurl := "https://" + self.Host + "/v1/identify"
	http_method := "POST"
	http_uri := "/v1/identify"
	data_type := "fingerprint"
	signature_version := "1"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	string_to_sign := http_method + "\n" + http_uri + "\n" + self.AccessKey + "\n" + data_type + "\n" + signature_version + "\n" + timestamp
	sign := self.GetSign(string_to_sign, self.AccessSecret)

	if audioFp == nil && humFp == nil {
		return "", ACR_ERR_CODE_GEN_FP_ERR, fmt.Errorf("Can not Create Fingerprint")
	}

	field_params := map[string]string{
		"access_key":        self.AccessKey,
		"timestamp":         timestamp,
		"signature":         sign,
		"data_type":         data_type,
		"signature_version": signature_version,
	}

	if userParams != nil {
		for key, val := range userParams {
			field_params[key] = val
		}
	}

	file_params := map[string][]byte{}
	if audioFp != nil && len(audioFp) != 0 {
		file_params["sample"] = audioFp
		field_params["sample_bytes"] = strconv.Itoa(len(audioFp))
	}
	if humFp != nil && len(humFp) != 0 {
		file_params["sample_hum"] = humFp
		field_params["sample_hum_bytes"] = strconv.Itoa(len(humFp))
	}

	result, retCode, err := self.Post(qurl, field_params, file_params, self.TimeoutS)
	return result, retCode, err
}

func (self *Recognizer) GenErrRes(code int, msg string) string {
	res := make(map[string]interface{})
	res["status"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	jsons, errs := json.Marshal(res)
	if errs != nil {
		return fmt.Sprintf("{\"status\":{\"msg\":\"Http Error\", \"code\":%d}}", ACR_ERR_CODE_JSON_ERR)
	}
	return string(jsons)
}

/*
 *    This function support most of audio / video files.
 *
 *    Audio: mp3, wav, m4a, flac, aac, amr, ape, ogg ...
 *    Video: mp4, mkv, wmv, flv, ts, avi ...
 *
 *    @param data: file_path query buffer
 *    @param startSeconds: skip (start_seconds) seconds from from the beginning of (data)
 *    @param lenSeconds: use rec_length seconds data to recongize
 *    @param userParams: some User-defined fields.
 *    @return result metainfos
 */
func (self *Recognizer) RecognizeByFileBuffer(data []byte, startSeconds int, lenSeconds int, userParams map[string]string) (string, int, error) {
	var humFp []byte
	var humRetCode int
	var humRetMsg error
	var audioFp []byte
	var audioRetCode int
	var audioRetMsg error
	if self.RecType == ACR_OPT_REC_HUMMING || self.RecType == ACR_OPT_REC_BOTH {
		humFp, humRetCode, humRetMsg = self.CreateHummingFingerprintByBuffer(data, startSeconds, lenSeconds)
	}
	if self.RecType == ACR_OPT_REC_AUDIO || self.RecType == ACR_OPT_REC_BOTH {
		audioFp, audioRetCode, audioRetMsg = self.CreateAudioFingerprintByBuffer(data, startSeconds, lenSeconds)
	}

	if humFp == nil && audioFp == nil {
		if audioRetMsg != nil {
			return "", audioRetCode, audioRetMsg
		} else {
			return "", humRetCode, humRetMsg
		}
	}

	result, retCode, err := self.DoRecognize(audioFp, humFp, userParams)
	if retCode != ACR_ERR_CODE_OK {
		return "", retCode, err
	}
	return result, ACR_ERR_CODE_OK, nil
}

/*
 *    This function support ACRCloud DB Fingerprint.
 *
 *    @param data: file_path query fingerprint buffer
 *    @param startSeconds: skip (start_seconds) seconds from from the beginning of (data)
 *    @param lenSeconds: use rec_length seconds data to recongize
 *    @param userParams: some User-defined fields.
 *    @return result metainfos
 */
func (self *Recognizer) RecognizeByFpBuffer(data []byte, startSeconds int, lenSeconds int, userParams map[string]string) string {
	var audioFp []byte
	var audioRetCode int
	var audioRetMsg error

	if self.RecType == ACR_OPT_REC_AUDIO || self.RecType == ACR_OPT_REC_BOTH {
		audioFp, audioRetCode, audioRetMsg = self.CreateAudioFingerprintByFpBuffer(data, startSeconds, lenSeconds)
	}

	if audioFp == nil {
		if audioRetMsg != nil {
			return self.GenErrRes(audioRetCode, fmt.Sprintf("%s", audioRetMsg))
		}
	}

	result, retCode, err := self.DoRecognize(audioFp, nil, userParams)
	if retCode != ACR_ERR_CODE_OK {
		return self.GenErrRes(retCode, fmt.Sprintf("%s", err))
	}
	return result
}

/*
 *    This function support most of audio / video files.
 *
 *    Audio: mp3, wav, m4a, flac, aac, amr, ape, ogg ...
 *    Video: mp4, mkv, wmv, flv, ts, avi ...
 *
 *    @param data: file_path
 *    @param startSeconds: skip (start_seconds) seconds from from the beginning of (data)
 *    @param lenSeconds: use rec_length seconds data to recongize
 *    @param userParams: some User-defined fields.
 *    @return result metainfos
 */
func (self *Recognizer) RecognizeByFile(filePath string, startSeconds int, lenSeconds int, userParams map[string]string) string {
	var humFp []byte
	var humRetCode int
	var humRetMsg error
	var audioFp []byte
	var audioRetCode int
	var audioRetMsg error
	if self.RecType == ACR_OPT_REC_HUMMING || self.RecType == ACR_OPT_REC_BOTH {
		humFp, humRetCode, humRetMsg = self.CreateHummingFingerprintByFile(filePath, startSeconds, lenSeconds)
	}
	if self.RecType == ACR_OPT_REC_AUDIO || self.RecType == ACR_OPT_REC_BOTH {
		audioFp, audioRetCode, audioRetMsg = self.CreateAudioFingerprintByFile(filePath, startSeconds, lenSeconds)
	}

	if humFp == nil && audioFp == nil {
		if audioRetMsg != nil {
			return self.GenErrRes(audioRetCode, fmt.Sprintf("%s", audioRetMsg))
		} else {
			return self.GenErrRes(humRetCode, fmt.Sprintf("%s", humRetMsg))
		}
	}

	result, retCode, err := self.DoRecognize(audioFp, humFp, userParams)
	if retCode != ACR_ERR_CODE_OK {
		return self.GenErrRes(retCode, fmt.Sprintf("%s", err))
	}
	return result
}

// Only support Microsoft PCM, 16 bit, mono 8000 Hz
func (self *Recognizer) Recognize(data []byte, userParams map[string]string) string {
	var humFp []byte
	var humRetCode int
	var humRetMsg error
	var audioFp []byte
	var audioRetCode int
	var audioRetMsg error
	if self.RecType == ACR_OPT_REC_HUMMING || self.RecType == ACR_OPT_REC_BOTH {
		humFp, humRetCode, humRetMsg = self.CreateHummingFingerprint(data)
	}
	if self.RecType == ACR_OPT_REC_AUDIO || self.RecType == ACR_OPT_REC_BOTH {
		audioFp, audioRetCode, audioRetMsg = self.CreateAudioFingerprint(data)
	}

	if humFp == nil && audioFp == nil {
		if audioRetMsg != nil {
			return self.GenErrRes(audioRetCode, fmt.Sprintf("%s", audioRetMsg))
		} else {
			return self.GenErrRes(humRetCode, fmt.Sprintf("%s", humRetMsg))
		}
	}

	result, retCode, err := self.DoRecognize(audioFp, humFp, userParams)
	if retCode != ACR_ERR_CODE_OK {
		return self.GenErrRes(retCode, fmt.Sprintf("%s", err))
	}
	return result
}
