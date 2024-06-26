cu=`pwd`
rm -rf release/packages && echo "success deletion" || echo "not success"
mkdir -p release/packages/
os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64'
for os in $os_all; do
    for arch in $arch_all; do
      set GOOS=$os
      set GOARCH=$arch
      if [ $os = "windows" ]; then
        go build -o $os"_"$arch".exe" && echo "Success build for arch "$arch" and os "$os || echo "No problem"
        mv $os"_"$arch".exe" release/packages && echo "Move success" || echo "Move not success"
      else
        go build -o $os"_"$arch && echo "Success build for arch "$arch" and os "$os || echo "No problem"
        mv $os"_"$arch release/packages && echo "Move success" || echo "Move not success"
      fi
    done
done
echo "Success Build"
cd $cu