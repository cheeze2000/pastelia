module.exports = {
	"env": {
		"browser": true,
		"commonjs": true,
		"es2021": true,
		"vue/setup-compiler-macros": true
	},
	"extends": [
		"eslint:recommended",
		"plugin:vue/vue3-strongly-recommended"
	],
	"globals": {
		"hljs": "readonly"
	},
	"parserOptions": {
		"ecmaVersion": 13,
		"sourceType": "module"
	},
	"plugins": [
		"vue"
	],
	"rules": {
		"comma-dangle": ["error", "never"],
		"indent": ["error", "tab"],
		"linebreak-style": ["error", "unix"],
		"quotes": ["error", "double"],
		"semi": ["error", "always"],
		"sort-imports": ["error", { allowSeparatedGroups: true }],
		"vue/html-indent": ["error", "tab"],
		"vue/html-self-closing": "off",
		"vue/max-attributes-per-line": ["error", { singleline: 2, multiline: 2 }],
		"vue/multi-word-component-names": "off",
		"vue/require-default-prop": "off"
	}
};
