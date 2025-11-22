function Invoke-EchoServer {

    $TcpPort = if ($env:TCP_PORT) { $env:TCP_PORT } else { 20001 }
    $UdpPort = if ($env:UDP_PORT) { $env:UDP_PORT } else { 20002 }

    $HostLogDir = Join-Path $PSScriptRoot "data"
    if (!(Test-Path $HostLogDir)) {
        New-Item -ItemType Directory -Path $HostLogDir | Out-Null
    }

    $ContainerLogFile = "/app/data/echo.log"

    Write-Host "Host Log Directory: $HostLogDir"
    Write-Host "Container Log File: $ContainerLogFile"
    Write-Host "Using TCP Port: $TcpPort"
    Write-Host "Using UDP Port: $UdpPort"

    $cid = docker run -d `
        -p 0:20001 `
        -p 0:20002/udp `
        -v "${HostLogDir}:/app/data" `
        echoservertest `
        -tcp-port $TcpPort `
        -udp-port $UdpPort `
        -log-file $ContainerLogFile

    Write-Host "Container started: $cid"
    Write-Host "Assigned ports:"
    docker port $cid
}

Invoke-EchoServer
