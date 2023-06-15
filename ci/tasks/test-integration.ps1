trap {
  write-error $_
  exit 1
}

$env:GOPATH = Join-Path -Path $PWD "gopath"
$env:PATH = $env:GOPATH + "/bin;C:/go/bin;" + $env:PATH

cd $env:GOPATH/src/github.com/shono09835/bosh-cli

powershell.exe bin/install-go.ps1

go.exe run github.com/onsi/ginkgo/ginkgo -race -trace integration

if ($LastExitCode -ne 0) {
  Write-Error $_
  exit 1
}
