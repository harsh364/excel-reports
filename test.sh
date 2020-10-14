#! /bin/bash

echo

# Just Once for installation
# go get -u golang.org/x/lint
# go get -u github.com/paddyforan/combinedcoverage
# go get -u github.com/wadey/gocovmerge


dir=$1
testDir="tests"
if [[ -z $1 ]]; then
  dir="./..."
fi

echo "********************** Go-Vet **********************"
go vet $(go list $dir | grep -v '/vendor/')
echo
echo

echo "********************** Go-Lint **********************"
# go list ./... | grep -v /vendor/ | xargs -L1 golint
golint $dir
echo
echo

echo "********************** Go-Test **********************"
excluded=("/$testDir" "/reportstatus" "/cmd")

rm -rf tests_coverage
mkdir -p tests_coverage

for pkg in $(go list $dir | grep -v '/vendor/'); do
  for i in "${excluded[@]}"; do
    if [[ "$pkg" == *"$i"* ]] ; then
      continue 2
    fi
  done

  # echo "################### testing ${pkg} ###################"
  pkgTest=${pkg/example.com\/emailreports/.\/$testDir}
	go test -race -covermode=atomic -coverprofile=tests_coverage/${pkg//\//.}.out -coverpkg $pkg $pkgTest
  # echo
done

echo "********************** Total Coverage **********************"
totalCoverage=$(combinedcoverage $(find tests_coverage/*.out | grep -v '/vendor/'))
echo "                       $totalCoverage                       "
echo "********************** Total Coverage **********************"
echo

echo $totalCoverage > test_coverage.txt

gocovmerge $(find tests_coverage/*.out | grep -v '/vendor/') > tests_coverage/total.out

go tool cover -html=tests_coverage/total.out -o test_coverage.html