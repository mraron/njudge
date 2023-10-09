function Modal({ children, isOpen, onClose }) {
    return (
        <div>
            {isOpen && <div className="z-40 fixed inset-0 bg-white dark:bg-black opacity-50" onClick={onClose} />}
            {isOpen && (
                <div
                    className="z-50 fixed left-3 right-3 top-1/2 sm:left-1/2 sm:right-auto transform sm:-translate-x-1/2 -translate-y-1/2"
                    onClick={(event) => event.stopPropagation()}>
                    {children}
                </div>
            )}
        </div>
    )
}

export default Modal
