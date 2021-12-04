printf "INFO\tOperating System set to $INPUT_OS!\n"

function ubuntu() {
  sudo apt update
  DEBIAN_FRONTEND="noninteractive" sudo apt-get install autoconf automake cmake curl libtool make ninja-build patch python3-pip unzip virtualenv

  sudo wget -O /usr/local/bin/bazel https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-$([ $(uname -m) = "aarch64" ] && echo "arm64" || echo "amd64")
  sudo chmod +x /usr/local/bin/bazel
}

function fedora() {
  sudo dnf install dnf-plugins-core cmake libtool libstdc++ libstdc++-static libatomic ninja-build lld patch aspell-en
  sudo dnf copr enable vbatts/bazel
  sudo dnf install bazel3-$BAZEL_VERSION
}

function darwin() {
  brew install coreutils wget cmake libtool go automake ninja clang-format autoconf aspell
  brew install bazelisk
}

if [[ "$INPUT_OS" = *"ubuntu"* ]]; then
  ubuntu
  if [ $? -eq 1 ]; then
    printf "ERROR\tUnable to setup ubuntu environment\n"
      exit 1
  fi
elif [[ "$INPUT_OS" = *"macos"* ]]; then
  darwin 
  if [ $? -eq 1 ]; then
    printf "ERROR\tUnable to setup darwin environment\n"
      exit 1
  fi
else
  printf "ERROR\tOperating system not supported\n"
  exit 1
fi

if ! type "bazel" > /dev/null 2>&1; then
  printf "ERROR\tbazel not found\n"
  exit 1;
fi

