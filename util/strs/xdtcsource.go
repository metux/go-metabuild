package strs

import (
	"fmt"
)

const maxWidth = 71

func XdtCSource(in string, src string, resname string) string {
	packed := PackXML(string(in))
	return fmt.Sprintf(
		"/* automatically generated from %s */\n"+
			"#ifdef __SUNPRO_C\n"+
			"#pragma align 4 (%s)\n"+
			"#endif\n"+
			"#ifdef __GNUC__\n"+
			"static const char %s[] __attribute__ ((__aligned__ (4))) =\n"+
			"#else\n"+
			"static const char %s[] =\n"+
			"#endif\n"+
			"{\n"+
			"%s};\n\nstatic const unsigned %s_length = %du;\n\n",
		src, resname, resname, resname, CLiteral(packed, "  ", maxWidth), resname, len(packed))
}
