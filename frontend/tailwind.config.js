module.exports = {
    content: ["./src/**/*.{js, jsx, ts, tsx}", "./src/*.html"],
    darkMode: "class",
    theme: {
        extend: {
            width: {
                '100': '25rem',
                '104': '26rem',
                '108': '27rem',
                '112': '28rem',
                '116': '29rem',
                '120': '30rem',
                '124': '31rem',
                '128': '32rem',
                '132': '33rem',
                '136': '34rem',
                '140': '35rem',
                '144': '36rem',
                '148': '37rem',
                '152': '38rem',
                '156': '39rem',
                '160': '40rem'
            },
            boxShadow: {
                "frame": "0 2.5px 3.5px -2.5px rgba(0, 0, 0, 0.15)"
            },
            transitionProperty: {
                "opacity": "opacity",
                "transform": "transform",
                "width-opacity": "width, opacity",
                "transform-opacity": "transform, opacity",
                "height-opacity": "height, opacity"
            },
            borderRadius: {
                "container": "0.25rem"
            },
            animation: {
                "spin-slow": "spin 2s linear infinite",
            },
            borderWidth: {
                1: "1px",
                3: "3px",
            },
            colors: {
                "textcol": "rgb(var(--color-textcol))",
                "border-def": "rgb(var(--color-border-def))",
                "border-med": "rgb(var(--color-border-med))",
                "border-str": "rgb(var(--color-border-str))",
                "border-xstr": "rgb(var(--color-border-xstr))",
                "divide-def": "rgb(var(--color-divide-def))",
                "code-bg": "rgb(var(--color-code-bg))",
                "frame-bg": "rgb(var(--color-frame-bg))",
                "button": "rgb(var(--color-button))",
                "button-hover": "rgb(var(--color-button-hover))",
                "icon": "rgb(var(--color-icon))",
                "highlight": "rgb(var(--color-highlight))",
                "grey-900": "rgb(var(--color-grey-900))",
                "grey-875": "rgb(var(--color-grey-875))",
                "grey-850": "rgb(var(--color-grey-850))",
                "grey-825": "rgb(var(--color-grey-825))",
                "grey-800": "rgb(var(--color-grey-800))",
                "grey-775": "rgb(var(--color-grey-775))",
                "grey-750": "rgb(var(--color-grey-750))",
                "grey-725": "rgb(var(--color-grey-725))",
                "grey-700": "rgb(var(--color-grey-700))",
                "grey-675": "rgb(var(--color-grey-675))",
                "grey-650": "rgb(var(--color-grey-650))",
                "grey-625": "rgb(var(--color-grey-625))",
                "grey-600": "rgb(var(--color-grey-600))",
                "grey-575": "rgb(var(--color-grey-575))",
                "grey-550": "rgb(var(--color-grey-550))",
                "grey-525": "rgb(var(--color-grey-525))",
                "grey-500": "rgb(var(--color-grey-500))",
                "grey-475": "rgb(var(--color-grey-475))",
                "grey-450": "rgb(var(--color-grey-450))",
                "grey-425": "rgb(var(--color-grey-425))",
                "grey-400": "rgb(var(--color-grey-400))",
                "grey-375": "rgb(var(--color-grey-375))",
                "grey-350": "rgb(var(--color-grey-350))",
                "grey-325": "rgb(var(--color-grey-325))",
                "grey-300": "rgb(var(--color-grey-300))",
                "grey-275": "rgb(var(--color-grey-275))",
                "grey-250": "rgb(var(--color-grey-250))",
                "grey-225": "rgb(var(--color-grey-225))",
                "grey-200": "rgb(var(--color-grey-200))",
                "grey-175": "rgb(var(--color-grey-175))",
                "grey-150": "rgb(var(--color-grey-150))",
                "grey-125": "rgb(var(--color-grey-125))",
                "grey-100": "rgb(var(--color-grey-100))",
            },
        },
    },
    plugins: [],
};
