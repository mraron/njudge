import { useContext } from "react"
import ThemeContext from "../../contexts/theme/ThemeContext"
import Editor from "@monaco-editor/react"

function CodeEditor(props) {
    const { theme } = useContext(ThemeContext)

    return <Editor theme={`${theme}-theme`} className="h-10 editor" {...props} />
}

export default CodeEditor
