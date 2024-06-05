param(
    [string]$version = "0.1.0"
)

$url = "https://github.com/ccc159/freeport/releases/download/v$version/freeport-$version-windows-amd64.zip"
$output = "$env:Temp\freeport-$version-windows-amd64.zip"
$installPath = "C:\Program Files\freeport"

# Create install directory if it doesn't exist
if (-Not (Test-Path -Path $installPath)) {
    New-Item -ItemType Directory -Force -Path $installPath
}

# Download the ZIP file
Write-Host "Downloading $url to $output"
Invoke-WebRequest -Uri $url -OutFile $output

# Extract the ZIP file
Write-Host "Extracting $output to $installPath"
Expand-Archive -Path $output -DestinationPath $installPath -Force

# Add the install path to the system PATH environment variable
$env:Path += ";$installPath"
[Environment]::SetEnvironmentVariable("Path", $env:Path, [System.EnvironmentVariableTarget]::Machine)

# Verify installation
$freeportPath = Join-Path -Path $installPath -ChildPath "freeport-windows-amd64.exe"
if (Test-Path -Path $freeportPath) {
    Write-Host "Freeport installed successfully!"
} else {
    Write-Host "Installation failed."
}
