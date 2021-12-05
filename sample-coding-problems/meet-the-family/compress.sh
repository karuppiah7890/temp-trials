set -o errexit
set -o nounset
set -o pipefail
# set -o xtrace

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

ZIP_PATH="Family.zip"

echo "Removing any existing ${ZIP_PATH}"

rm -rfv ${ZIP_PATH}

echo "Compressing ${MY_DIR} into ${ZIP_PATH} excluding binaries, extra markdown files and scripts"

pushd ${MY_DIR}

zip -r ${ZIP_PATH} . -x STORY.md compress.sh .gitignore meet-the-family

popd

echo "Done!"
