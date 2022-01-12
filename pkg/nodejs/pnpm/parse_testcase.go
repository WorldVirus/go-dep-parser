package pnpm

import (
	"github.com/aquasecurity/go-dep-parser/pkg/types"
)

var (
	// docker run --name node --rm -it node:12-alpine sh
	// npm install -g pnpm
	// pnpm add promise jquery
	// pnpm list -depth 10 | grep -E -o "\S+\s+[0-9]+(\.[0-9]+)+$" | awk '{printf("{\""$1"\", \""$2"\",\"\"},\n")}' | sort -u
	// have to hardcode value of depth flag because of there is a problem with output dependencies  https://github.com/pnpm/pnpm/issues/3583
	pnpmNormal = []types.Library{
		{"jquery", "3.6.0", ""},
		{"promise", "8.1.0", ""},
		{"asap", "2.0.6", ""},
	}

	// ... and
	// pnpm add react redux
	// pnpm list -depth 10 | grep -E -o "\S+\s+[0-9]+(\.[0-9]+)+$" | awk '{printf("{\""$1"\", \""$2"\",\"\"},\n")}' | sort -u
	pnpmReact = []types.Library{
		{"jquery", "3.6.0", ""},
		{"promise", "8.1.0", ""},
		{"asap", "2.0.6", ""},
		{"react", "17.0.2", ""},
		{"loose-envify", "1.4.0", ""},
		{"js-tokens", "4.0.0", ""},
		{"object-assign", "4.1.1", ""},
		{"redux", "4.1.2", ""},
		{"@babel/runtime", "7.16.3", ""},
		{"regenerator-runtime", "0.13.9", ""},
	}

	// ... and
	// pnpm add mocha
	// pnpm list -depth 10 | grep -E -o "\S+\s+[0-9]+(\.[0-9]+)+$" | awk '{printf("{\""$1"\", \""$2"\",\"\"},\n")}' | sort -u
	pnpmMany = []types.Library{
		{"@babel/runtime", "7.16.3", ""},
		{"@ungap/promise-all-settled", "1.1.2", ""},
		{"ansi-colors", "4.1.1", ""},
		{"ansi-regex", "5.0.1", ""},
		{"ansi-styles", "4.3.0", ""},
		{"anymatch", "3.1.2", ""},
		{"argparse", "2.0.1", ""},
		{"asap", "2.0.6", ""},
		{"balanced-match", "1.0.2", ""},
		{"binary-extensions", "2.2.0", ""},
		{"brace-expansion", "1.1.11", ""},
		{"braces", "3.0.2", ""},
		{"browser-stdout", "1.3.1", ""},
		{"camelcase", "6.2.1", ""},
		{"chalk", "4.1.2", ""},
		{"chokidar", "3.5.2", ""},
		{"cliui", "7.0.4", ""},
		{"color-convert", "2.0.1", ""},
		{"color-name", "1.1.4", ""},
		{"concat-map", "0.0.1", ""},
		{"debug", "4.3.2", ""},
		{"decamelize", "4.0.0", ""},
		{"diff", "5.0.0", ""},
		{"emoji-regex", "8.0.0", ""},
		{"escalade", "3.1.1", ""},
		{"escape-string-regexp", "4.0.0", ""},
		{"fill-range", "7.0.1", ""},
		{"find-up", "5.0.0", ""},
		{"flat", "5.0.2", ""},
		{"fs.realpath", "1.0.0", ""},
		{"fsevents", "2.3.2", ""},
		{"get-caller-file", "2.0.5", ""},
		{"glob", "7.1.7", ""},
		{"glob-parent", "5.1.2", ""},
		{"growl", "1.10.5", ""},
		{"has-flag", "4.0.0", ""},
		{"he", "1.2.0", ""},
		{"inflight", "1.0.6", ""},
		{"inherits", "2.0.4", ""},
		{"is-binary-path", "2.1.0", ""},
		{"is-extglob", "2.1.1", ""},
		{"is-fullwidth-code-point", "3.0.0", ""},
		{"is-glob", "4.0.3", ""},
		{"is-number", "7.0.0", ""},
		{"is-plain-obj", "2.1.0", ""},
		{"is-unicode-supported", "0.1.0", ""},
		{"isexe", "2.0.0", ""},
		{"jquery", "3.6.0", ""},
		{"js-tokens", "4.0.0", ""},
		{"js-yaml", "4.1.0", ""},
		{"locate-path", "6.0.0", ""},
		{"log-symbols", "4.1.0", ""},
		{"loose-envify", "1.4.0", ""},
		{"minimatch", "3.0.4", ""},
		{"mocha", "9.1.3", ""},
		{"ms", "2.1.2", ""},
		{"ms", "2.1.3", ""},
		{"nanoid", "3.1.25", ""},
		{"normalize-path", "3.0.0", ""},
		{"object-assign", "4.1.1", ""},
		{"once", "1.4.0", ""},
		{"p-limit", "3.1.0", ""},
		{"p-locate", "5.0.0", ""},
		{"path-exists", "4.0.0", ""},
		{"path-is-absolute", "1.0.1", ""},
		{"picomatch", "2.3.0", ""},
		{"promise", "8.1.0", ""},
		{"randombytes", "2.1.0", ""},
		{"react", "17.0.2", ""},
		{"readdirp", "3.6.0", ""},
		{"redux", "4.1.2", ""},
		{"regenerator-runtime", "0.13.9", ""},
		{"require-directory", "2.1.1", ""},
		{"safe-buffer", "5.2.1", ""},
		{"serialize-javascript", "6.0.0", ""},
		{"string-width", "4.2.3", ""},
		{"strip-ansi", "6.0.1", ""},
		{"strip-json-comments", "3.1.1", ""},
		{"supports-color", "7.2.0", ""},
		{"supports-color", "8.1.1", ""},
		{"to-regex-range", "5.0.1", ""},
		{"which", "2.0.2", ""},
		{"workerpool", "6.1.5", ""},
		{"wrap-ansi", "7.0.0", ""},
		{"wrappy", "1.0.2", ""},
		{"y18n", "5.0.8", ""},
		{"yargs", "16.2.0", ""},
		{"yargs-parser", "20.2.4", ""},
		{"yargs-unparser", "2.0.0", ""},
		{"yocto-queue", "0.1.0", ""},
	}

	// ... and
	// pnpm add lodash request chalk commander express async axios vue
	// pnpm list -depth 10 | grep -E -o "\S+\s+[0-9]+(\.[0-9]+)+$" | awk '{printf("{\""$1"\", \""$2"\",\"\"},\n")}' | sort -u
	pnpmManyV2 = []types.Library{
		{"@babel/runtime", "7.16.3", ""},
		{"@ungap/promise-all-settled", "1.1.2", ""},
		{"accepts", "1.3.7", ""},
		{"ajv", "6.12.6", ""},
		{"ansi-colors", "4.1.1", ""},
		{"ansi-regex", "5.0.1", ""},
		{"ansi-styles", "4.3.0", ""},
		{"anymatch", "3.1.2", ""},
		{"argparse", "2.0.1", ""},
		{"array-flatten", "1.1.1", ""},
		{"asap", "2.0.6", ""},
		{"asn1", "0.2.6", ""},
		{"assert-plus", "1.0.0", ""},
		{"async", "3.2.2", ""},
		{"asynckit", "0.4.0", ""},
		{"aws-sign2", "0.7.0", ""},
		{"aws4", "1.11.0", ""},
		{"axios", "0.24.0", ""},
		{"balanced-match", "1.0.2", ""},
		{"bcrypt-pbkdf", "1.0.2", ""},
		{"binary-extensions", "2.2.0", ""},
		{"body-parser", "1.19.0", ""},
		{"brace-expansion", "1.1.11", ""},
		{"braces", "3.0.2", ""},
		{"browser-stdout", "1.3.1", ""},
		{"bytes", "3.1.0", ""},
		{"camelcase", "6.2.1", ""},
		{"caseless", "0.12.0", ""},
		{"chalk", "4.1.2", ""},
		{"chalk", "5.0.0", ""},
		{"chokidar", "3.5.2", ""},
		{"cliui", "7.0.4", ""},
		{"color-convert", "2.0.1", ""},
		{"color-name", "1.1.4", ""},
		{"combined-stream", "1.0.8", ""},
		{"commander", "8.3.0", ""},
		{"concat-map", "0.0.1", ""},
		{"content-disposition", "0.5.3", ""},
		{"content-type", "1.0.4", ""},
		{"cookie", "0.4.0", ""},
		{"cookie-signature", "1.0.6", ""},
		{"core-util-is", "1.0.2", ""},
		{"dashdash", "1.14.1", ""},
		{"debug", "2.6.9", ""},
		{"debug", "4.3.2", ""},
		{"decamelize", "4.0.0", ""},
		{"delayed-stream", "1.0.0", ""},
		{"depd", "1.1.2", ""},
		{"destroy", "1.0.4", ""},
		{"diff", "5.0.0", ""},
		{"ecc-jsbn", "0.1.2", ""},
		{"ee-first", "1.1.1", ""},
		{"emoji-regex", "8.0.0", ""},
		{"encodeurl", "1.0.2", ""},
		{"escalade", "3.1.1", ""},
		{"escape-html", "1.0.3", ""},
		{"escape-string-regexp", "4.0.0", ""},
		{"etag", "1.8.1", ""},
		{"express", "4.17.1", ""},
		{"extend", "3.0.2", ""},
		{"extsprintf", "1.3.0", ""},
		{"fast-deep-equal", "3.1.3", ""},
		{"fast-json-stable-stringify", "2.1.0", ""},
		{"fill-range", "7.0.1", ""},
		{"finalhandler", "1.1.2", ""},
		{"find-up", "5.0.0", ""},
		{"flat", "5.0.2", ""},
		{"follow-redirects", "1.14.5", ""},
		{"forever-agent", "0.6.1", ""},
		{"form-data", "2.3.3", ""},
		{"forwarded", "0.2.0", ""},
		{"fresh", "0.5.2", ""},
		{"fs.realpath", "1.0.0", ""},
		{"fsevents", "2.3.2", ""},
		{"get-caller-file", "2.0.5", ""},
		{"getpass", "0.1.7", ""},
		{"glob", "7.1.7", ""},
		{"glob-parent", "5.1.2", ""},
		{"growl", "1.10.5", ""},
		{"har-schema", "2.0.0", ""},
		{"har-validator", "5.1.5", ""},
		{"has-flag", "4.0.0", ""},
		{"he", "1.2.0", ""},
		{"http-errors", "1.7.2", ""},
		{"http-errors", "1.7.3", ""},
		{"http-signature", "1.2.0", ""},
		{"iconv-lite", "0.4.24", ""},
		{"inflight", "1.0.6", ""},
		{"inherits", "2.0.3", ""},
		{"inherits", "2.0.4", ""},
		{"ipaddr.js", "1.9.1", ""},
		{"is-binary-path", "2.1.0", ""},
		{"is-extglob", "2.1.1", ""},
		{"is-fullwidth-code-point", "3.0.0", ""},
		{"is-glob", "4.0.3", ""},
		{"is-number", "7.0.0", ""},
		{"is-plain-obj", "2.1.0", ""},
		{"is-typedarray", "1.0.0", ""},
		{"is-unicode-supported", "0.1.0", ""},
		{"isexe", "2.0.0", ""},
		{"isstream", "0.1.2", ""},
		{"jquery", "3.6.0", ""},
		{"js-tokens", "4.0.0", ""},
		{"js-yaml", "4.1.0", ""},
		{"jsbn", "0.1.1", ""},
		{"json-schema", "0.4.0", ""},
		{"json-schema-traverse", "0.4.1", ""},
		{"json-stringify-safe", "5.0.1", ""},
		{"jsprim", "1.4.2", ""},
		{"locate-path", "6.0.0", ""},
		{"lodash", "4.17.21", ""},
		{"log-symbols", "4.1.0", ""},
		{"loose-envify", "1.4.0", ""},
		{"media-typer", "0.3.0", ""},
		{"merge-descriptors", "1.0.1", ""},
		{"methods", "1.1.2", ""},
		{"mime", "1.6.0", ""},
		{"mime-db", "1.51.0", ""},
		{"mime-types", "2.1.34", ""},
		{"minimatch", "3.0.4", ""},
		{"mocha", "9.1.3", ""},
		{"ms", "2.0.0", ""},
		{"ms", "2.1.1", ""},
		{"ms", "2.1.2", ""},
		{"ms", "2.1.3", ""},
		{"nanoid", "3.1.25", ""},
		{"negotiator", "0.6.2", ""},
		{"normalize-path", "3.0.0", ""},
		{"oauth-sign", "0.9.0", ""},
		{"object-assign", "4.1.1", ""},
		{"on-finished", "2.3.0", ""},
		{"once", "1.4.0", ""},
		{"p-limit", "3.1.0", ""},
		{"p-locate", "5.0.0", ""},
		{"parseurl", "1.3.3", ""},
		{"path-exists", "4.0.0", ""},
		{"path-is-absolute", "1.0.1", ""},
		{"path-to-regexp", "0.1.7", ""},
		{"performance-now", "2.1.0", ""},
		{"picomatch", "2.3.0", ""},
		{"promise", "8.1.0", ""},
		{"proxy-addr", "2.0.7", ""},
		{"psl", "1.8.0", ""},
		{"punycode", "2.1.1", ""},
		{"qs", "6.5.2", ""},
		{"qs", "6.7.0", ""},
		{"randombytes", "2.1.0", ""},
		{"range-parser", "1.2.1", ""},
		{"raw-body", "2.4.0", ""},
		{"react", "17.0.2", ""},
		{"readdirp", "3.6.0", ""},
		{"redux", "4.1.2", ""},
		{"regenerator-runtime", "0.13.9", ""},
		{"request", "2.88.2", ""},
		{"require-directory", "2.1.1", ""},
		{"safe-buffer", "5.1.2", ""},
		{"safe-buffer", "5.2.1", ""},
		{"safer-buffer", "2.1.2", ""},
		{"send", "0.17.1", ""},
		{"serialize-javascript", "6.0.0", ""},
		{"serve-static", "1.14.1", ""},
		{"setprototypeof", "1.1.1", ""},
		{"sshpk", "1.16.1", ""},
		{"statuses", "1.5.0", ""},
		{"string-width", "4.2.3", ""},
		{"strip-ansi", "6.0.1", ""},
		{"strip-json-comments", "3.1.1", ""},
		{"supports-color", "7.2.0", ""},
		{"supports-color", "8.1.1", ""},
		{"to-regex-range", "5.0.1", ""},
		{"toidentifier", "1.0.0", ""},
		{"tough-cookie", "2.5.0", ""},
		{"tunnel-agent", "0.6.0", ""},
		{"tweetnacl", "0.14.5", ""},
		{"type-is", "1.6.18", ""},
		{"unpipe", "1.0.0", ""},
		{"uri-js", "4.4.1", ""},
		{"utils-merge", "1.0.1", ""},
		{"uuid", "3.4.0", ""},
		{"vary", "1.1.2", ""},
		{"verror", "1.10.0", ""},
		{"vue", "2.6.14", ""},
		{"which", "2.0.2", ""},
		{"workerpool", "6.1.5", ""},
		{"wrap-ansi", "7.0.0", ""},
		{"wrappy", "1.0.2", ""},
		{"y18n", "5.0.8", ""},
		{"yargs", "16.2.0", ""},
		{"yargs-parser", "20.2.4", ""},
		{"yargs-unparser", "2.0.0", ""},
		{"yocto-queue", "0.1.0", ""},
	}
)