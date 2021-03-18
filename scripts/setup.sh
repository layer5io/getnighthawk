BAZEL_VERSION="3.7.2"

printf "INFO\tOperating System set to $INPUT_OS!\n"

function ubuntu() {
  sudo apt update
  DEBIAN_FRONTEND="noninteractive" sudo apt-get install -y libtool cmake automake autoconf make ninja-build curl unzip virtualenv 
  curl -fLO "https://github.com/bazelbuild/bazel/releases/download/${BAZEL_VERSION}/bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh"
  DEBIAN_FRONTEND="noninteractive" sudo apt-get install -y bazel=$BAZEL_VERSION
  sudo chmod +x "bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh"
  ./bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh --user
}

function fedora() {
  sudo dnf install dnf-plugins-core cmake libtool libstdc++ libstdc++-static libatomic ninja-build lld patch aspell-en
  sudo dnf copr enable vbatts/bazel
  sudo dnf install bazel3-$BAZEL_VERSION
}

function darwin() {
  sudo brew install coreutils wget cmake libtool automake ninja clang-format autoconf aspell
  curl -fLO "https://github.com/bazelbuild/bazel/releases/download/${BAZEL_VERSION}/bazel-${BAZEL_VERSION}-installer-darwin-x86_64.sh"
  sudo chmod +x "bazel-${BAZEL_VERSION}-installer-darwin-x86_64.sh"
  ./bazel-${BAZEL_VERSION}-installer-darwin-x86_64.sh --user
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

