<h2 align="center">pr0music</h2>
<p align="center">Nutzerbot für <a href="https://pr0gramm.com">pr0gramm.com</a> um Musik in Posts zu erkennen.</p>
<p align="center"><b>Achtung!</b> Diese Version nutzt <u>ACR Cloud</u> als Dienstleister!</p>

<!-- Tags - 1 -->
<p align="center">
    <a href="https://ci.s-c.systems/Pacerino/pr0music">
        <img src="https://ci.s-c.systems/api/badges/Pacerino/pr0music/status.svg?ref=refs/heads/acr_cloud" />
    </a>
</p>

<br />

## Development 

```bash
# Clone the repository 
git clone https://github.com/Pacerino/pr0music

# Switch the branch
git checkout -q acr_cloud

# Enter into the directory
cd pr0music

# Install the dependencies
go get ./...

# Copy the .env
cp .env.sample .env

# Fill out the .env
nano .env

# Start the Server
go run .
```