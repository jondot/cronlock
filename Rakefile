
include FileUtils

#

ARCHS = [
  ['linux-arm', ''],
  ['linux-amd64',''],
  ['linux-386',''],
  ['darwin-amd64','']
]

task :package do
  ARCHS.each do |arch, prefix|
    `mkdir build`
    `mkdir -p pkg/#{arch}`
    bin = go_build("build/cronlock", arch)

    cd("pkg/#{arch}") do
      cp "../../#{bin}", "cronlock"

      tarball = "cronlock-#{arch}.tar.gz"
      `tar -czf #{tarball} *`
      cp tarball, "../"
    end

    `rm -rf pkg/#{arch}`
    `rm -rf build`
  end
end


def go_build(label, arch)
    system({"GOOS" => arch.split('-')[0], 
          "GOARCH" => arch.split('-')[1]}, 
          "go build -o #{label}-#{arch}")
    "#{label}-#{arch}"
end

