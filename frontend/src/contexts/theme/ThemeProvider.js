import { useState } from "react";

import ThemeContext from "./ThemeContext";

function ThemeProvider({ children }) {
    const [theme, setTheme] = useState("light");

    const changeTheme = (newTheme) => {
        if (!["light", "dark"].includes(theme)) {
            return;
        }
        setTheme(newTheme);

        const doc = document.documentElement;
        doc.setAttribute("color-scheme", newTheme);

        if (newTheme === "light" && doc.classList.contains("dark")) {
            doc.classList.remove("dark");
        }
        if (newTheme === "dark" && !doc.classList.contains("dark")) {
            doc.classList.add("dark");
        }
    };
    return (
        <ThemeContext.Provider value={{ theme, changeTheme }}>
            {children}
        </ThemeContext.Provider>
    );
}

export default ThemeProvider;
