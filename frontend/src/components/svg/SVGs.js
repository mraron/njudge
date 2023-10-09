export function SVGDropdownListArrow({ isOpen }) {
    return (
        <svg
            className={`transform w-4 h-4 mr-1 fill-current ${
                isOpen ? "rotate-90" : ""
            } transition-transform duration-100 shrink-0`}
            viewBox="0 0 48 48"
            xmlns="http://www.w3.org/2000/svg">
            <rect width="48" height="48" fill="none" />
            <path d="M19.5,37.4l11.9-12a1.9,1.9,0,0,0,0-2.8l-11.9-12A2,2,0,0,0,16,12h0V36h0a2,2,0,0,0,3.5,1.4Z" />
        </svg>
    )
}

export function SVGDropdownMenuArrow({ cls = null, isOpen }) {
    return (
        <svg
            className={`transform ${cls} w-6 h-6 ml-4 ${
                isOpen ? "rotate-180" : ""
            } transition-transform duration-100 shrink-0`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                d="M7 10L12 15L17 10"
                className="stroke-current"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    )
}

export function SVGDropdownFrameArrow({ isOpen }) {
    return (
        <svg
            className={`w-7 h-7 transform ${
                isOpen ? "rotate-180" : "rotate-0"
            } h-[1.85rem] text-sm font-medium transition-transform duration-100 shrink-0`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M7 10L12 15L17 10"
                stroke="#ffffff"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    )
}

export function SVGNotFound() {
    return (
        <svg className="w-80 text-grey-750" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path
                d="M9 17C9.85038 16.3697 10.8846 16 12 16C13.1154 16 14.1496 16.3697 15 17"
                stroke="#1C274C"
                strokeWidth="1.5"
                strokeLinecap="round"
                className="stroke-current"
            />
            <ellipse cx="15" cy="10.5" rx="1" ry="1.5" fill="#1C274C" className="fill-current" />
            <ellipse cx="9" cy="10.5" rx="1" ry="1.5" fill="#1C274C" className="fill-current" />
            <path
                d="M22 12C22 16.714 22 19.0711 20.5355 20.5355C19.0711 22 16.714 22 12 22C7.28595 22 4.92893 22 3.46447 20.5355C2 19.0711 2 16.714 2 12C2 7.28595 2 4.92893 3.46447 3.46447C4.92893 2 7.28595 2 12 2C16.714 2 19.0711 2 20.5355 3.46447C21.5093 4.43821 21.8356 5.80655 21.9449 8"
                stroke="#1C274C"
                strokeWidth="1.5"
                strokeLinecap="round"
                className="stroke-current"
            />
        </svg>
    )
}

export function SVGSpinner({ cls = null }) {
    return (
        <svg
            className={`${cls} animate-spin-slow`}
            viewBox="0 0 14.00 14.00"
            xmlns="http://www.w3.org/2000/svg"
            fill="#000000">
            <g fill="none" fillRule="evenodd">
                <circle cx="7" cy="7" r="6" className="stroke-grey-625" strokeWidth="2"></circle>
                <path className="fill-indigo-600" fillRule="nonzero" d="M7 0a7 7 0 0 1 7 7h-2a5 5 0 0 0-5-5V0z"></path>
            </g>
        </svg>
    )
}

export function SVGEllipsis({ cls = null, title }) {
    return (
        <svg
            className={`inline ${cls} stroke-current shrink-0`}
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            stroke="#000000"
            strokeWidth="1"
            strokeLinecap="round"
            strokeLinejoin="miter">
            <title>{title}</title>
            <line x1="5.99" y1="12" x2="6" y2="12" strokeLinecap="round" strokeWidth="2" />
            <line x1="11.99" y1="12" x2="12" y2="12" strokeLinecap="round" strokeWidth="2" />
            <line x1="17.99" y1="12" x2="18" y2="12" strokeLinecap="round" strokeWidth="2" />
        </svg>
    )
}

export function SVGView({ cls = null }) {
    return (
        <svg className={`${cls} shrink-0`} viewBox="-3.5 0 32 32" version="1.1" xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M12.406 13.844c1.188 0 2.156 0.969 2.156 2.156s-0.969 2.125-2.156 2.125-2.125-0.938-2.125-2.125 0.938-2.156 2.125-2.156zM12.406 8.531c7.063 0 12.156 6.625 12.156 6.625 0.344 0.438 0.344 1.219 0 1.656 0 0-5.094 6.625-12.156 6.625s-12.156-6.625-12.156-6.625c-0.344-0.438-0.344-1.219 0-1.656 0 0 5.094-6.625 12.156-6.625zM12.406 21.344c2.938 0 5.344-2.406 5.344-5.344s-2.406-5.344-5.344-5.344-5.344 2.406-5.344 5.344 2.406 5.344 5.344 5.344z"
            />
        </svg>
    )
}
