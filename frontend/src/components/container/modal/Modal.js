import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faClose } from "@fortawesome/free-solid-svg-icons";

function Modal({ children, isOpen, onClose }) {
    return (
        <>
            {isOpen && (
                <div
                    className="z-40 fixed inset-0 bg-black opacity-50"
                    onClick={onClose}>
                    <button
                        onClick={onClose}
                        className="flex items-center justify-center absolute top-3 right-3 p-2 rounded-full bg-transparent hover:bg-grey-100 dark:hover:bg-grey-700">
                        <FontAwesomeIcon
                            icon={faClose}
                            className="w-6 h-6 text-white"
                        />
                    </button>
                </div>
            )}
            {isOpen && (
                <div
                    className="z-50 fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"
                    onClick={(event) => event.stopPropagation()}>
                    {children}
                </div>
            )}
        </>
    );
}

export default Modal;
