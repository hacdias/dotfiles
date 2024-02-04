# Create a new directory and enter it
function mkd() {
	mkdir -p "$@" && cd "$_";
}

# Create a data URL from a file
function dataurl() {
	local mimeType=$(file -b --mime-type "$1");
	if [[ $mimeType == text/* ]]; then
		mimeType="${mimeType};charset=utf-8";
	fi
	echo "data:${mimeType};base64,$(openssl base64 -in "$1" | tr -d '\n')";
}

# `o` with no arguments opens the current directory, otherwise opens the given
# location
function o() {
	if [ $# -eq 0 ]; then
		open .;
	else
		open "$@";
	fi;
}

function get_dns() {
	networksetup -getdnsservers Wi-Fi
}

function clear_dns() {
	networksetup -setdnsservers Wi-Fi "Empty"
}

function cloudflare_dns() {
	networksetup -setdnsservers Wi-Fi 1.1.1.1 1.0.0.1 2606:4700:4700::1111 2606:4700:4700::1001
}

function fdiff() {
	git diff --name-only --relative --diff-filter=d | xargs bat --diff
}

function traceparent() {
	# see spec: https://www.w3.org/TR/trace-context
	# version-format   = trace-id "-" parent-id "-" trace-flags
	# trace-id         = 32HEXDIGLC  ; 16 bytes array identifier. All zeroes forbidden
	# parent-id        = 16HEXDIGLC  ; 8 bytes array identifier. All zeroes forbidden
	# trace-flags      = 2HEXDIGLC   ; 8 bit flags. Currently, only one bit is used. See below for detail
	version="00" # fixed in spec at 00
	trace_id="$(cat /dev/urandom | gtr -dc 'a-f0-9' | fold -w 32 | head -n 1)"
	parent_id="00$(cat /dev/urandom | gtr -dc 'a-f0-9' | fold -w 14 | head -n 1)"
	tace_flag="01"   # sampled
	trace_parent="$version-$trace_id-$parent_id-$tace_flag"
	echo $trace_parent
}