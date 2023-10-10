function NarrowPage({ children }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full sm:max-w-md">
                <div className="w-full px-3">{children}</div>
            </div>
        </div>
    )
}

export default NarrowPage
