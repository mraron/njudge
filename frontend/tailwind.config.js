module.exports = {
    content: ["./src/**/*.{js, jsx, ts, tsx}", "./src/*.html"],
    darkMode: "class",
    theme: {
        extend: {
            transitionProperty: {
                "opacity": "opacity",
                "transform": "transform",
                "width-opacity": "width, opacity",
                "height-opacity": "height, opacity"
            },
            borderRadius: {
                "container": "0.4rem"
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
                "bordercol": "rgb(var(--color-bordercol))",
                "dividecol": "rgb(var(--color-dividecol))",
                "codebgcol": "rgb(var(--color-codebgcol))",
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
