BAZEL_VERSION="3.7.2"
OS=$INPUT_OS

printf "INFO\tOperating System set to $OS!\n"

function ubuntu() {
  apt update
  apt install -y curl gnupg
  curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg
  mv bazel.gpg /etc/apt/trusted.gpg.d/
  echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
  DEBIAN_FRONTEND="noninteractive" apt-get install -y libtool cmake automake autoconf make ninja-build curl unzip virtualenv bazel=$BAZEL_VERSION
}

function fedora() {
  dnf install dnf-plugins-core cmake libtool libstdc++ libstdc++-static libatomic ninja-build lld patch aspell-en
  dnf copr enable vbatts/bazel
  dnf install bazel3-$BAZEL_VERSION
}

function darwin() {
  brew install coreutils wget cmake libtool automake ninja clang-format autoconf aspell
  curl -fLO "https://github.com/bazelbuild/bazel/releases/download/${BAZEL_VERSION}/bazel-${BAZEL_VERSION}-installer-darwin-x86_64.sh"
  chmod +x "bazel-${BAZEL_VERSION}-installer-darwin-x86_64.sh"
  ./bazel-${BAZEL_VERSION}-installer-darwin-x86_64.sh --user
}

if [ "$OS" = *"ubuntu"* ]; then
  ubuntu
  if [ $? -eq 1 ]; then
    printf "ERROR\tUnable to setup ubuntu environment\n"
      exit 1
  fi
elif [ "$OS" = *"macos"* ]; then
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

