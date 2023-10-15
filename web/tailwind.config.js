/** @type {import('tailwindcss').Config} */
export default {
	content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
	theme: {
		extend: {
			colors: {
				black: '#242424',
				white: '#fcfcfc',
				accent: '#9ee8f7',
			},
		},
	},
	plugins: [],
};
