#! /bin/sh

## execute me in from geeqie source root

srcroot="."
po_dir="${srcroot}/po"
locales_map="${po_dir}/locales.txt"
po_list=`find ${po_dir} -name "*.po"`
output="${po_dir}/translators"
targetdir="./"

mkdir -p "$targetdir"

cat "${po_dir}/LINGUAS" | while read -r locale
do
	printf "\n"
	awk '$1 == "'"$locale"'" {print $0}' "${locales_map}"
	awk '$0 ~/Translators:/ {
		while (1) {
			getline $0
		if ($0 == "#") {
			exit
			}
		else {
			print substr($0, 3)
			}
		}
		print $0
	}' "${po_dir}/${locale}.po"

done > "$output"
printf "\n\0" >> "$output"
