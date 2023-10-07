import { useContext, useEffect } from "react"
import ThemeContext from "../../contexts/theme/ThemeContext"
import Editor, { useMonaco } from "@monaco-editor/react"

function CodeEditor(props) {
    const { theme } = useContext(ThemeContext)
    const monaco = useMonaco()

    useEffect(() => {
        monaco?.editor.defineTheme("dark-theme", {
            base: "vs-dark",
            inherit: true,
            rules: [],
            colors: {
                "editor.background": "#0c080f",
            },
        })
        monaco?.editor.defineTheme("light-theme", {
            base: "vs",
            inherit: true,
            rules: [],
            colors: {
                "editor.background": "#faf7ff",
            },
        })
    }, [monaco])

    useEffect(() => {
        monaco?.editor.setTheme(theme === "dark" ? "dark-theme" : "light-theme")
    }, [monaco, theme])

    return <Editor className="editor" {...props} />
}

export default CodeEditor
