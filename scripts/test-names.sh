#!/bin/sh
# Checks the order in which mutter picks a name. The script creates users, so
# run it only in a throwaway container:
#
#   docker run --rm -v "$PWD":/src -w /src golang:1.26 scripts/test-names.sh
set -eu

CGO_ENABLED=0 go build -buildvcs=false -o /usr/local/bin/mutter .

useradd -m -c "Ada Lovelace,Room 1,555-1234,555-5678" ada
useradd -m -c "& Hopper" grace
useradd -m -c ",Room 2" alan
useradd -m -c "" linus
useradd -m -c "   " blank
useradd -m -c "Margaret Hamilton" margaret
runuser -u margaret -- env HOME=/home/margaret git config --global user.name "Maggie H"

fail=0

# check WANT USER [NAME=VALUE...] mutter [ARGS...]
check() {
	want=$1 user=$2
	shift 2
	got=$(cd /tmp && runuser -u "$user" -- env -i HOME="/home/$user" PATH=/usr/local/bin:/usr/bin:/bin "$@" | cut -d' ' -f1)
	if [ "$got" = "$want" ]; then
		echo "ok   $user: $*: $got"
	else
		echo "FAIL $user: $*: got $got, want $want"
		fail=1
	fi
}

check Ada ada mutter
check Grace grace mutter
check alan alan mutter
check linus linus mutter
check blank blank mutter
check Maggie margaret mutter
check Mo margaret MUTTER_NAME=Mo mutter
check Flag margaret MUTTER_NAME=Mo mutter -name Flag

exit $fail
