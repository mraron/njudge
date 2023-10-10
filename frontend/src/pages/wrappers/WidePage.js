function WidePage({children}) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl px-3">
                {children}
            </div>
        </div>
    )
}

export default WidePage