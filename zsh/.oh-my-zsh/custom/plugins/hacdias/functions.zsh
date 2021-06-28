# Create a new directory and enter it
function mkd() {
	mkdir -p "$@" && cd "$_";
}

# Define some cd aliases and their completions.
typeset -A cd_aliases
cd_aliases=(
	[cdc]="$HOME/Code"
	[cdd]="$HOME/Documents"
	[cdmsce]="$HOME/Google Drive/MSCE"
	[cdt]="$HOME/Code/@tue"
)

for k in "${(@k)cd_aliases}"; do
	function $k() {
		cd "$cd_aliases[$0]/$1"
	}

	function _$k() {
		((CURRENT == 2)) && _files -/ -W $cd_aliases[${0:1}]
	}

	compdef _$k $k
done

# Create a data URL from a file
function dataurl() {
	local mimeType=$(file -b --mime-type "$1");
	if [[ $mimeType == text/* ]]; then
		mimeType="${mimeType};charset=utf-8";
	fi
	echo "data:${mimeType};base64,$(openssl base64 -in "$1" | tr -d '\n')";
}

# Normalize `open` across Linux, macOS, and Windows.
# This is needed to make the `o` function (see below) cross-platform.
if [ ! $(uname -s) = 'Darwin' ]; then
	if grep -q Microsoft /proc/version; then
		# Ubuntu on Windows using the Linux subsystem
		alias open='explorer.exe';
	else
		alias open='xdg-open';
	fi
fi

# `o` with no arguments opens the current directory, otherwise opens the given
# location
function o() {
	if [ $# -eq 0 ]; then
		open .;
	else
		open "$@";
	fi;
}

function load_fstar() {
	export PATH="$HOME/z3/bin:$PATH"
	eval $(opam env)
	fstar.exe --version
}
