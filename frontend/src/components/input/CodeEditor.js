import Editor from "@monaco-editor/react"

function CodeEditor({ language, value, onChange, readOnly }) {
    return (
        <div className={`h-80 sm:h-[32rem]`}>
            <Editor
                theme="standard-theme"
                className="editor"
                options={{
                    fontFamily: "JetBrains Mono",
                    fontSize: 13,
                    lineHeight: 21,
                    readOnly: readOnly,
                }}
                value={value}
                language={language}
                onChange={onChange}
            />
        </div>
    )
}

export default CodeEditor
