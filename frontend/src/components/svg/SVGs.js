export function SVGAttachmentFile({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M19 9V17.8C19 18.9201 19 19.4802 18.782 19.908C18.5903 20.2843 18.2843 20.5903 17.908 20.782C17.4802 21 16.9201 21 15.8 21H8.2C7.07989 21 6.51984 21 6.09202 20.782C5.71569 20.5903 5.40973 20.2843 5.21799 19.908C5 19.4802 5 18.9201 5 17.8V6.2C5 5.07989 5 4.51984 5.21799 4.09202C5.40973 3.71569 5.71569 3.40973 6.09202 3.21799C6.51984 3 7.0799 3 8.2 3H13M19 9L13 3M19 9H14C13.4477 9 13 8.55228 13 8V3"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGAttachmentDescription({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M9 17H15M9 13H15M9 9H10M13 3H8.2C7.0799 3 6.51984 3 6.09202 3.21799C5.71569 3.40973 5.40973 3.71569 5.21799 4.09202C5 4.51984 5 5.0799 5 6.2V17.8C5 18.9201 5 19.4802 5.21799 19.908C5.40973 20.2843 5.71569 20.5903 6.09202 20.782C6.51984 21 7.0799 21 8.2 21H15.8C16.9201 21 17.4802 21 17.908 20.782C18.2843 20.5903 18.5903 20.2843 18.782 19.908C19 19.4802 19 18.9201 19 17.8V9M13 3L19 9M13 3V7.4C13 7.96005 13 8.24008 13.109 8.45399C13.2049 8.64215 13.3578 8.79513 13.546 8.89101C13.7599 9 14.0399 9 14.6 9H19"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGInformation({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M12 11V16M12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21ZM12.0498 8V8.1L11.9502 8.1002V8H12.0498Z"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGSubmit() {
    return (
        <svg
            className="w-6 h-6 mr-2"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M10 14L13 21L20 4L3 11L6.5 12.5"
                stroke="#000000"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGAttachment() {
    return (
        <svg
            className="w-6 h-6 mr-2"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M20 10.9696L11.9628 18.5497C10.9782 19.4783 9.64274 20 8.25028 20C6.85782 20 5.52239 19.4783 4.53777 18.5497C3.55315 17.6211 3 16.3616 3 15.0483C3 13.7351 3.55315 12.4756 4.53777 11.547M14.429 6.88674L7.00403 13.8812C6.67583 14.1907 6.49144 14.6106 6.49144 15.0483C6.49144 15.4861 6.67583 15.9059 7.00403 16.2154C7.33224 16.525 7.77738 16.6989 8.24154 16.6989C8.70569 16.6989 9.15083 16.525 9.47904 16.2154L13.502 12.4254M8.55638 7.75692L12.575 3.96687C13.2314 3.34779 14.1217 3 15.05 3C15.9783 3 16.8686 3.34779 17.525 3.96687C18.1814 4.58595 18.5502 5.4256 18.5502 6.30111C18.5502 7.17662 18.1814 8.01628 17.525 8.63535L16.5 9.601"
                stroke="#000000"
                strokeWidth="1.5"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGDropdownListArrow({ isOpen }) {
    return (
        <svg
            className={`w-4 h-4 mr-1 fill-current ${
                isOpen ? "rotate-90" : ""
            } transition-all duration-100 shrink-0`}
            viewBox="0 0 48 48"
            xmlns="http://www.w3.org/2000/svg">
            <rect width="48" height="48" fill="none" />
            <path d="M19.5,37.4l11.9-12a1.9,1.9,0,0,0,0-2.8l-11.9-12A2,2,0,0,0,16,12h0V36h0a2,2,0,0,0,3.5,1.4Z" />
        </svg>
    );
}

export function SVGDropdownMenuArrow({ cls = null, isOpen }) {
    return (
        <svg
            className={`${cls} w-6 h-6 ml-4 ${
                isOpen ? "rotate-180" : ""
            } transition-all duration-150`}
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
    );
}

export function SVGCopy({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                fillRule="evenodd"
                clipRule="evenodd"
                d="M21 8C21 6.34315 19.6569 5 18 5H10C8.34315 5 7 6.34315 7 8V20C7 21.6569 8.34315 23 10 23H18C19.6569 23 21 21.6569 21 20V8ZM19 8C19 7.44772 18.5523 7 18 7H10C9.44772 7 9 7.44772 9 8V20C9 20.5523 9.44772 21 10 21H18C18.5523 21 19 20.5523 19 20V8Z"
                fill="#0F0F0F"
            />
            <path
                className="fill-current"
                d="M6 3H16C16.5523 3 17 2.55228 17 2C17 1.44772 16.5523 1 16 1H6C4.34315 1 3 2.34315 3 4V18C3 18.5523 3.44772 19 4 19C4.55228 19 5 18.5523 5 18V4C5 3.44772 5.44772 3 6 3Z"
                fill="#0F0F0F"
            />
        </svg>
    );
}

export function SVGAvatar({ cls = null }) {
    return (
        <svg
            className={`${cls} fill-current`}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 52 52"
            enableBackground="new 0 0 52 52"
            xmlSpace="preserve">
            <path
                d="M50,43v2.2c0,2.6-2.2,4.8-4.8,4.8H6.8C4.2,50,2,47.8,2,45.2V43c0-5.8,6.8-9.4,13.2-12.2
                c0.2-0.1,0.4-0.2,0.6-0.3c0.5-0.2,1-0.2,1.5,0.1c2.6,1.7,5.5,2.6,8.6,2.6s6.1-1,8.6-2.6c0.5-0.3,1-0.3,1.5-0.1
                c0.2,0.1,0.4,0.2,0.6,0.3C43.2,33.6,50,37.1,50,43z M26,2c6.6,0,11.9,5.9,11.9,13.2S32.6,28.4,26,28.4s-11.9-5.9-11.9-13.2
                S19.4,2,26,2z"
            />
        </svg>
    );
}

export function SVGHamburger() {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            className="w-6 h-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16M4 18h16"
            />
        </svg>
    );
}

export function SVGClose({ cls = null }) {
    return (
        <svg
            style={{ pointerEvents: "none" }}
            fill="#000000"
            className={`${cls} fill-current`}
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 460.775 460.775"
            xmlSpace="preserve">
            <path
                d="M285.08,230.397L456.218,59.27c6.076-6.077,6.076-15.911,0-21.986L423.511,4.565c-2.913-2.911-6.866-4.55-10.992-4.55
                c-4.127-4.127-8.08-4.55-10.993-4.55c-4.127,0-8.08,1.639-10.993,4.55l-171.138,171.14L59.25,4.565c-2.913-2.911-6.866-4.55-10.993-4.55
                c-4.126,0-8.08,1.639-10.992,4.55L4.558,37.284c-6.077,6.075-6.077,15.909,0,21.986l171.138,171.128L4.575,401.505
                c-6.074,6.077-6.074,15.911,0,21.986l32.709,32.719c2.911,2.911,6.865,4.55,10.992,4.55c4.127,0,8.08-1.639,10.994-4.55
                l171.117-171.12l171.118,171.12c2.913,2.911,6.866,4.55,10.993,4.55c4.128,0,8.081-1.639,10.992-4.55l32.709-32.719
                c6.074-6.075,6.074-15.909,0-21.986L285.08,230.397z"
            />
        </svg>
    );
}

export function SVGDropdownFilterArrow({ isOpen }) {
    return (
        <svg
            className={`${
                isOpen ? "rotate-180" : "rotate-0"
            } h-[1.85rem] text-sm font-medium transition duration-150`}
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
    );
}

export function SVGNotFound() {
    return (
        <svg
            className="w-80"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                d="M9 17C9.85038 16.3697 10.8846 16 12 16C13.1154 16 14.1496 16.3697 15 17"
                stroke="#1C274C"
                strokeWidth="1.5"
                strokeLinecap="round"
                className="stroke-grey-750"
            />
            <ellipse
                cx="15"
                cy="10.5"
                rx="1"
                ry="1.5"
                fill="#1C274C"
                className="fill-grey-750"
            />
            <ellipse
                cx="9"
                cy="10.5"
                rx="1"
                ry="1.5"
                fill="#1C274C"
                className="fill-grey-750"
            />
            <path
                d="M22 12C22 16.714 22 19.0711 20.5355 20.5355C19.0711 22 16.714 22 12 22C7.28595 22 4.92893 22 3.46447 20.5355C2 19.0711 2 16.714 2 12C2 7.28595 2 4.92893 3.46447 3.46447C4.92893 2 7.28595 2 12 2C16.714 2 19.0711 2 20.5355 3.46447C21.5093 4.43821 21.8356 5.80655 21.9449 8"
                stroke="#1C274C"
                strokeWidth="1.5"
                strokeLinecap="round"
                className="stroke-grey-750"
            />
        </svg>
    );
}

export function SVGStatistics({ cls = null }) {
    return (
        <svg
            className={`${cls} fill-current`}
            fill="#000000"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 100 100"
            enableBackground="new 0 0 100 100"
            xmlSpace="preserve">
            <path
                d="M46.05,60.163H31.923c-0.836,0-1.513,0.677-1.513,1.513v21.934c0,0.836,0.677,1.513,1.513,1.513H46.05
                c0.836,0,1.512-0.677,1.512-1.513V61.675C47.562,60.839,46.885,60.163,46.05,60.163z"
            />
            <path
                d="M68.077,14.878H53.95c-0.836,0-1.513,0.677-1.513,1.513v67.218c0,0.836,0.677,1.513,1.513,1.513h14.127
                c0.836,0,1.513-0.677,1.513-1.513V16.391C69.59,15.555,68.913,14.878,68.077,14.878z"
            />
            <path
                d="M90.217,35.299H76.09c-0.836,0-1.513,0.677-1.513,1.513v46.797c0,0.836,0.677,1.513,1.513,1.513h14.126
                c0.836,0,1.513-0.677,1.513-1.513V36.812C91.729,35.977,91.052,35.299,90.217,35.299z"
            />
            <path
                d="M23.91,35.299H9.783c-0.836,0-1.513,0.677-1.513,1.513v46.797c0,0.836,0.677,1.513,1.513,1.513H23.91
                c0.836,0,1.513-0.677,1.513-1.513V36.812C25.423,35.977,24.746,35.299,23.91,35.299z"
            />
        </svg>
    );
}

export function SVGRecent({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                fillRule="evenodd"
                clipRule="evenodd"
                d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12 6.477 2 12 2s10 4.477 10 10zm-4.581 3.324a1 1 0 0 0-.525-1.313L13 12.341V6.5a1 1 0 0 0-2 0v6.17c0 .6.357 1.143.909 1.379l4.197 1.8a1 1 0 0 0 1.313-.525z"
                fill="#000000"
            />
        </svg>
    );
}

export function SVGResults() {
    return (
        <svg
            className="w-6 h-6 mr-2 fill-current"
            fill="#000000"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 100 100"
            enableBackground="new 0 0 100 100"
            xmlSpace="preserve">
            <path
                d="M27.953,46.506c-1.385-2.83-2.117-6.008-2.117-9.192c0-1.743,0.252-3.534,0.768-5.468c0.231-0.87,0.521-1.702,0.847-2.509
                    c-1.251-0.683-2.626-1.103-4.101-1.103c-5.47,0-9.898,5.153-9.898,11.517c0,4.452,2.176,8.305,5.354,10.222L5.391,56.217
                    c-0.836,0.393-1.387,1.337-1.387,2.392v10.588c0,1.419,0.991,2.569,2.21,2.569h7.929V60.656c0-3.237,1.802-6.172,4.599-7.481
                    l10.262-4.779C28.624,47.792,28.273,47.161,27.953,46.506z"
            />
            <path
                d="M60.137,34.801h34.092v-0.001c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761c0,0,0-0.001,0-0.001
                    l0-6.43h0c0-0.973-0.789-1.761-1.761-1.761c-0.002,0-0.004,0.001-0.006,0.001v-0.005H56.133c1.614,2.114,2.844,4.627,3.526,7.435
                    C59.874,33.168,60.03,33.999,60.137,34.801z"
            />
            <path
                d="M95.996,66.436c0-0.973-0.789-1.761-1.761-1.761c-0.002,0-0.004,0.001-0.006,0.001v-0.005H72.007v7.095v1.994
                    c0,0.293-0.016,0.582-0.045,0.867h22.267v-0.001c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761l0-0.001
                    L95.996,66.436L95.996,66.436z"
            />
            <path
                d="M94.235,44.762c-0.002,0-0.004,0.001-0.006,0.001v-0.005H58.944c-0.159,0.419-0.327,0.836-0.514,1.249
                    c-0.364,0.802-0.773,1.569-1.224,2.297l10.288,4.908c0.781,0.378,1.473,0.897,2.078,1.503h24.657v-0.001
                    c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761c0,0,0-0.001,0-0.001l0-6.43h0
                    C95.996,45.55,95.207,44.762,94.235,44.762z"
            />
            <path
                d="M65.323,57.702l-11.551-5.51l-4.885-2.33c2.134-1.344,3.866-3.418,5-5.917c0.899-1.984,1.435-4.231,1.435-6.631
                    c0-1.348-0.213-2.627-0.512-3.863c-1.453-5.983-6.126-10.392-11.736-10.392c-5.504,0-10.106,4.251-11.648,10.065
                    c-0.356,1.333-0.602,2.72-0.602,4.189c0,2.552,0.596,4.93,1.609,7c1.171,2.4,2.906,4.379,5.018,5.651l-4.678,2.178l-11.926,5.554
                    c-1.037,0.485-1.717,1.654-1.717,2.959v11.111v1.994c0,1.756,1.224,3.181,2.735,3.181h42.417c1.511,0,2.735-1.424,2.735-3.181
                    v-1.994V60.656C67.019,59.355,66.349,58.198,65.323,57.702z"
            />
        </svg>
    );
}

export function SVGGoogle() {
    return (
        <svg
            className="w-6 h-6"
            viewBox="0 0 32 32"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                d="M30.0014 16.3109C30.0014 15.1598 29.9061 14.3198 29.6998 13.4487H16.2871V18.6442H24.1601C24.0014 19.9354 23.1442 21.8798 21.2394 23.1864L21.2127 23.3604L25.4536 26.58L25.7474 26.6087C28.4458 24.1665 30.0014 20.5731 30.0014 16.3109Z"
                fill="#4285F4"
            />
            <path
                d="M16.2863 29.9998C20.1434 29.9998 23.3814 28.7553 25.7466 26.6086L21.2386 23.1863C20.0323 24.0108 18.4132 24.5863 16.2863 24.5863C12.5086 24.5863 9.30225 22.1441 8.15929 18.7686L7.99176 18.7825L3.58208 22.127L3.52441 22.2841C5.87359 26.8574 10.699 29.9998 16.2863 29.9998Z"
                fill="#34A853"
            />
            <path
                d="M8.15964 18.769C7.85806 17.8979 7.68352 16.9645 7.68352 16.0001C7.68352 15.0356 7.85806 14.1023 8.14377 13.2312L8.13578 13.0456L3.67083 9.64746L3.52475 9.71556C2.55654 11.6134 2.00098 13.7445 2.00098 16.0001C2.00098 18.2556 2.55654 20.3867 3.52475 22.2845L8.15964 18.769Z"
                fill="#FBBC05"
            />
            <path
                d="M16.2864 7.4133C18.9689 7.4133 20.7784 8.54885 21.8102 9.4978L25.8419 5.64C23.3658 3.38445 20.1435 2 16.2864 2C10.699 2 5.8736 5.1422 3.52441 9.71549L8.14345 13.2311C9.30229 9.85555 12.5086 7.4133 16.2864 7.4133Z"
                fill="#EB4335"
            />
        </svg>
    );
}

export function SVGLogin() {
    return (
        <svg
            className="w-6 h-6 mr-2"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                fillRule="evenodd"
                clipRule="evenodd"
                d="M8 6C8 3.79086 9.79086 2 12 2H17.5C19.9853 2 22 4.01472 22 6.5V17.5C22 19.9853 19.9853 22 17.5 22H12C9.79086 22 8 20.2091 8 18V17C8 16.4477 8.44772 16 9 16C9.55228 16 10 16.4477 10 17V18C10 19.1046 10.8954 20 12 20H17.5C18.8807 20 20 18.8807 20 17.5V6.5C20 5.11929 18.8807 4 17.5 4H12C10.8954 4 10 4.89543 10 6V7C10 7.55228 9.55228 8 9 8C8.44772 8 8 7.55228 8 7V6ZM12.2929 8.29289C12.6834 7.90237 13.3166 7.90237 13.7071 8.29289L16.7071 11.2929C17.0976 11.6834 17.0976 12.3166 16.7071 12.7071L13.7071 15.7071C13.3166 16.0976 12.6834 16.0976 12.2929 15.7071C11.9024 15.3166 11.9024 14.6834 12.2929 14.2929L13.5858 13L5 13C4.44772 13 4 12.5523 4 12C4 11.4477 4.44772 11 5 11L13.5858 11L12.2929 9.70711C11.9024 9.31658 11.9024 8.68342 12.2929 8.29289Z"
                fill="#0F1729"
            />
        </svg>
    );
}

export function SVGLock({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M9.52513 4.52513C10.1815 3.86875 11.0717 3.5 12 3.5C12.9283 3.5 13.8185 3.86875 14.4749 4.52513C14.7873 4.83751 15.0344 5.20276 15.2078 5.59999L15.4078 6.05825C15.6287 6.56443 16.2181 6.79568 16.7243 6.57477L17.6408 6.17478C18.147 5.95387 18.3783 5.36445 18.1574 4.85827L17.9574 4.40001C17.6355 3.66243 17.1763 2.98389 16.5962 2.40381C15.3772 1.18482 13.7239 0.5 12 0.5C10.2761 0.5 8.62279 1.18482 7.40381 2.40381C6.18482 3.62279 5.5 5.27609 5.5 7V10H5C3.34315 10 2 11.3431 2 13V20C2 21.6569 3.34315 23 5 23H19C20.6569 23 22 21.6569 22 20V13C22 11.3431 20.6569 10 19 10H8.5V7C8.5 6.07174 8.86875 5.1815 9.52513 4.52513Z"
                fill="#000000"
            />
        </svg>
    );
}

export function SVGMail({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 -2.5 20 20"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink">
            <g
                className="fill-current"
                stroke="none"
                strokeWidth="1"
                fill="none"
                fillRule="evenodd">
                <g
                    className="fill-currentt"
                    transform="translate(-300.000000, -922.000000)"
                    fill="#000000">
                    <g
                        className="fill-current"
                        transform="translate(56.000000, 160.000000)">
                        <path
                            className="fill-current"
                            d="M262,764.291 L254,771.318 L246,764.281 L246,764 L262,764 L262,764.291 Z M246,775 L246,766.945 L254,773.98 L262,766.953 L262,775 L246,775 Z M244,777 L264,777 L264,762 L244,762 L244,777 Z"
                        />
                    </g>
                </g>
            </g>
        </svg>
    );
}

export function SVGConfirm({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink">
            <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                <rect fillRule="nonzero" x="0" y="0" width="24" height="24" />
                <circle
                    className="stroke-current"
                    stroke="#0C0310"
                    strokeWidth="2"
                    strokeLinecap="round"
                    cx="12"
                    cy="12"
                    r="9"
                />
                <path
                    className="stroke-current"
                    d="M8.5,12.5 L10.151,14.5638 C10.3372,14.7965 10.6843,14.8157 10.895,14.605 L15.5,10"
                    stroke="#0C0310"
                    strokeWidth="2"
                    strokeLinecap="round"
                />
            </g>
        </svg>
    );
}

export function SVGSpinner({ cls = null }) {
    return (
        <svg
            className={`${cls} text-grey-275 dark:text-grey-625 animate-spin-slow fill-indigo-600`}
            viewBox="-6 -6 112 113"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                fill="currentColor"
            />
            <path
                d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                fill="currentFill"
            />
        </svg>
    );
}

export function SVGLeftArrow({ cls = null }) {
    return (
        <svg
            className={`${cls} fill-current`}
            viewBox="0 0 1024 1024"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M768 903.232l-50.432 56.768L256 512l461.568-448 50.432 56.768L364.928 512z"
                fill="#000000"
            />
        </svg>
    );
}

export function SVGDoubleRightArrow({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                fillRule="evenodd"
                clipRule="evenodd"
                d="M12.293 7.293a1 1 0 0 1 1.414 0l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L15.586 12l-3.293-3.293a1 1 0 0 1 0-1.414Z"
                fill="#000000"
            />
            <path
                className="fill-current"
                fillRule="evenodd"
                clipRule="evenodd"
                d="M6.293 7.293a1 1 0 0 1 1.414 0l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L9.586 12 6.293 8.707a1 1 0 0 1 0-1.414Z"
                fill="#000000"
            />
        </svg>
    );
}

export function SVGDots({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="stroke-current"
                d="M12 13C12.5523 13 13 12.5523 13 12C13 11.4477 12.5523 11 12 11C11.4477 11 11 11.4477 11 12C11 12.5523 11.4477 13 12 13Z"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
            <path
                className="stroke-current"
                d="M19 13C19.5523 13 20 12.5523 20 12C20 11.4477 19.5523 11 19 11C18.4477 11 18 11.4477 18 12C18 12.5523 18.4477 13 19 13Z"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
            <path
                className="stroke-current"
                d="M5 13C5.55228 13 6 12.5523 6 12C6 11.4477 5.55228 11 5 11C4.44772 11 4 11.4477 4 12C4 12.5523 4.44772 13 5 13Z"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGChange({ cls = null }) {
    return (
        <svg
            fill="#000000"
            className={`${cls} fill-current`}
            viewBox="0 0 56 56"
            xmlns="http://www.w3.org/2000/svg">
            <path d="M .3321 40.0118 C .3321 41.5117 1.4337 42.5430 3.0977 42.5430 L 8.4415 42.5430 C 12.4727 42.5430 14.9103 41.3711 17.8165 37.9727 L 23.0899 31.8320 L 28.3399 37.9727 C 31.2462 41.3711 33.6603 42.5664 37.7618 42.5664 L 42.0508 42.5664 L 42.0508 47.7695 C 42.0508 49.0352 42.8476 49.8320 44.1604 49.8320 C 44.7229 49.8320 45.3085 49.6211 45.7304 49.2461 L 54.5898 41.9336 C 55.6679 41.0664 55.6448 39.6602 54.5898 38.7930 L 45.7304 31.4336 C 45.3085 31.0586 44.7229 30.8477 44.1604 30.8477 C 42.8476 30.8477 42.0508 31.6445 42.0508 32.9102 L 42.0508 37.4571 L 37.8790 37.4571 C 35.4649 37.4571 33.9649 36.6836 32.0430 34.4571 L 26.4415 27.9180 L 32.0430 21.4024 C 33.9649 19.1524 35.4649 18.3789 37.8790 18.3789 L 42.0508 18.3789 L 42.0508 23.0898 C 42.0508 24.3555 42.8476 25.1524 44.1604 25.1524 C 44.7229 25.1524 45.3085 24.9414 45.7304 24.5664 L 54.5898 17.2539 C 55.6679 16.3867 55.6448 15.0039 54.5898 14.1133 L 45.7304 6.7539 C 45.3085 6.3789 44.7229 6.1680 44.1604 6.1680 C 42.8476 6.1680 42.0508 6.9649 42.0508 8.2305 L 42.0508 13.2930 L 37.7618 13.2930 C 33.6603 13.2930 31.2462 14.4883 28.3399 17.8867 L 23.0899 24.0274 L 17.8165 17.8867 C 14.9103 14.4883 12.4727 13.2930 8.4415 13.2930 L 3.0977 13.2930 C 1.4337 13.2930 .3321 14.3242 .3321 15.8477 C .3321 17.3477 1.4571 18.4024 3.0977 18.4024 L 8.5352 18.4024 C 10.8087 18.4024 12.2384 19.1758 14.1368 21.4024 L 19.7384 27.9180 L 14.1368 34.4571 C 12.2149 36.6836 10.7852 37.4571 8.5352 37.4571 L 3.0977 37.4571 C 1.4571 37.4571 .3321 38.5118 .3321 40.0118 Z" />
        </svg>
    );
}

export function SVGSettings({ cls = null }) {
    return (
        <svg
            className={`${cls} fill-current`}
            fill="#000000"
            viewBox="0 0 1920 1920"
            xmlns="http://www.w3.org/2000/svg">
            <path
                d="M1703.534 960c0-41.788-3.84-84.48-11.633-127.172l210.184-182.174-199.454-340.856-265.186 88.433c-66.974-55.567-143.323-99.389-223.85-128.415L1158.932 0h-397.78L706.49 269.704c-81.43 29.138-156.423 72.282-223.962 128.414l-265.073-88.32L18 650.654l210.184 182.174C220.39 875.52 216.55 918.212 216.55 960s3.84 84.48 11.633 127.172L18 1269.346l199.454 340.856 265.186-88.433c66.974 55.567 143.322 99.389 223.85 128.415L761.152 1920h397.779l54.663-269.704c81.318-29.138 156.424-72.282 223.963-128.414l265.073 88.433 199.454-340.856-210.184-182.174c7.793-42.805 11.633-85.497 11.633-127.285m-743.492 395.294c-217.976 0-395.294-177.318-395.294-395.294 0-217.976 177.318-395.294 395.294-395.294 217.977 0 395.294 177.318 395.294 395.294 0 217.976-177.317 395.294-395.294 395.294"
                fillRule="evenodd"
            />
        </svg>
    );
}

export function SVGWrong({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink">
            <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                <rect fillRule="nonzero" x="0" y="0" width="24" height="24" />
                <circle
                    className="stroke-current"
                    stroke="#0C0310"
                    strokeWidth="2"
                    strokeLinecap="round"
                    cx="12"
                    cy="12"
                    r="9"
                />
                <line
                    className="stroke-current"
                    x1="14.1213"
                    y1="9.87866"
                    x2="9.8787"
                    y2="14.1213"
                    stroke="#0C0310"
                    strokeWidth="2"
                    strokeLinecap="round"
                />
                <line
                    className="stroke-current"
                    x1="9.87866"
                    y1="9.87866"
                    x2="14.1213"
                    y2="14.1213"
                    stroke="#0C0310"
                    strokeWidth="2"
                    strokeLinecap="round"
                />
            </g>
        </svg>
    );
}

export function SVGCorrect({ cls = null }) {
    return <SVGConfirm cls={cls} />;
}

export function SVGCode({ cls = null }) {
    return (
        <svg
            className={`${cls} fill-current`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                fillRule="evenodd"
                clipRule="evenodd"
                d="M14.9523 6.2635L10.4523 18.2635L9.04784 17.7368L13.5478 5.73682L14.9523 6.2635ZM19.1894 12.0001L15.9698 8.78042L17.0304 7.71976L21.3108 12.0001L17.0304 16.2804L15.9698 15.2198L19.1894 12.0001ZM8.03032 15.2198L4.81065 12.0002L8.03032 8.78049L6.96966 7.71983L2.68933 12.0002L6.96966 16.2805L8.03032 15.2198Z"
                fill="#080341"
            />
        </svg>
    );
}

export function SVGWrongSimple({ cls = null, title }) {
    return (
        <svg
            className={`${cls} fill-current`}
            fill="#000000"
            viewBox="0 0 200 200"
            xmlns="http://www.w3.org/2000/svg">
            <title>{title}</title>
            <path d="M114,100l49-49a9.9,9.9,0,0,0-14-14L100,86,51,37A9.9,9.9,0,0,0,37,51l49,49L37,149a9.9,9.9,0,0,0,14,14l49-49,49,49a9.9,9.9,0,0,0,14-14Z" />
        </svg>
    );
}

export function SVGCorrectSimple({ cls = null, title }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <title>{title}</title>
            <path
                className="stroke-current"
                d="M4 12.6111L8.92308 17.5L20 6.5"
                stroke="#000000"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
            />
        </svg>
    );
}

export function SVGHDD({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 512 512"
            xmlSpace="preserve">
            <path
                className="fill-current"
                d="M481.798,203.986l-85.257-69.984c-15.802-12.967-35.629-20.067-56.089-20.067H256h-84.457
                c-20.452,0-40.28,7.1-56.085,20.067l-85.258,69.984C11.938,217.201,0.012,238.638,0,262.916v62.38
                c0.016,40.199,32.579,72.762,72.77,72.77H256h183.23c40.191-0.008,72.762-32.571,72.77-72.77v-62.38
                C511.992,238.638,500.066,217.201,481.798,203.986z M41.584,262.916c0.008-8.677,3.458-16.345,9.137-22.044
                c5.703-5.676,13.372-9.134,22.049-9.141H256h183.23c8.677,0.008,16.345,3.466,22.053,9.141c5.675,5.699,9.125,13.367,9.134,22.044
                v62.38c-0.008,8.677-3.458,16.345-9.134,22.052c-5.708,5.676-13.376,9.126-22.053,9.134H256H72.77
                c-8.677-0.008-16.346-3.458-22.049-9.134c-5.679-5.707-9.129-13.375-9.137-22.052V262.916z"
            />
            <path
                className="fill-current"
                d="M326.168,319.444c12.924,0,23.393-10.478,23.393-23.39c0-12.912-10.47-23.389-23.393-23.389
                c-12.919,0-23.394,10.478-23.394,23.389C302.775,308.966,313.249,319.444,326.168,319.444z"
            />
            <path
                className="fill-current"
                d="M404.578,319.444c12.912,0,23.39-10.478,23.39-23.39c0-12.912-10.478-23.389-23.39-23.389
                c-12.919,0-23.397,10.478-23.397,23.389C381.181,308.966,391.659,319.444,404.578,319.444z"
            />
        </svg>
    );
}

export function SVGClock({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M23 12C23 18.0751 18.0751 23 12 23C5.92487 23 1 18.0751 1 12C1 5.92487 5.92487 1 12 1C18.0751 1 23 5.92487 23 12ZM3.00683 12C3.00683 16.9668 7.03321 20.9932 12 20.9932C16.9668 20.9932 20.9932 16.9668 20.9932 12C20.9932 7.03321 16.9668 3.00683 12 3.00683C7.03321 3.00683 3.00683 7.03321 3.00683 12Z"
                fill="#0F0F0F"
            />
            <path
                className="fill-current"
                d="M12 5C11.4477 5 11 5.44771 11 6V12.4667C11 12.4667 11 12.7274 11.1267 12.9235C11.2115 13.0898 11.3437 13.2343 11.5174 13.3346L16.1372 16.0019C16.6155 16.278 17.2271 16.1141 17.5032 15.6358C17.7793 15.1575 17.6155 14.5459 17.1372 14.2698L13 11.8812V6C13 5.44772 12.5523 5 12 5Z"
                fill="#0F0F0F"
            />
        </svg>
    );
}

export function SVGCheckmark({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            version="1.1"
            baseProfile="tiny"
            xmlns="http://www.w3.org/2000/svg"
            xmlnsXlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 42 42"
            xmlSpace="preserve">
            <path
                className="fill-current"
                d="M39.04,7.604l-2.398-1.93c-1.182-0.95-1.869-0.939-2.881,0.311L16.332,27.494l-8.111-6.739
                c-1.119-0.94-1.819-0.89-2.739,0.26l-1.851,2.41c-0.939,1.182-0.819,1.853,0.291,2.78l11.56,9.562c1.19,1,1.86,0.897,2.78-0.222
                l21.079-25.061C40.331,9.294,40.271,8.583,39.04,7.604z"
            />
        </svg>
    );
}

export function SVGPartiallyCorrect({ cls = null, title }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <title>{title}</title>
            <g clipPath="url(#clip0_949_23339)">
                <path
                    className="fill-current"
                    d="M17.5821 6.95711C17.9726 6.56658 17.9726 5.93342 17.5821 5.54289C17.1916 5.15237 16.5584 5.15237 16.1679 5.54289L5.54545 16.1653L1.70711 12.327C1.31658 11.9365 0.683417 11.9365 0.292893 12.327C-0.0976311 12.7175 -0.097631 13.3507 0.292893 13.7412L4.83835 18.2866C5.22887 18.6772 5.86204 18.6772 6.25256 18.2866L17.5821 6.95711Z"
                    fill="#000000"
                />
                <path
                    className="fill-current"
                    d="M23.5821 6.95711C23.9726 6.56658 23.9726 5.93342 23.5821 5.54289C23.1915 5.15237 22.5584 5.15237 22.1678 5.54289L10.8383 16.8724C10.4478 17.263 10.4478 17.8961 10.8383 18.2866C11.2288 18.6772 11.862 18.6772 12.2525 18.2866L23.5821 6.95711Z"
                    fill="#000000"
                />
            </g>
        </svg>
    );
}

export function SVGDotsSmall({ cls = null, title }) {
    return (
        <svg
            className={`${cls} stroke-current`}
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            stroke="#000000"
            strokeWidth="1"
            strokeLinecap="round"
            strokeLinejoin="miter">
            <title>{title}</title>
            <line
                x1="5.99"
                y1="12"
                x2="6"
                y2="12"
                strokeLinecap="round"
                strokeWidth="2"
            />
            <line
                x1="11.99"
                y1="12"
                x2="12"
                y2="12"
                strokeLinecap="round"
                strokeWidth="2"
            />
            <line
                x1="17.99"
                y1="12"
                x2="18"
                y2="12"
                strokeLinecap="round"
                strokeWidth="2"
            />
        </svg>
    );
}

export function SVGDash({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 20 20"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M2 9.5C2 9.22386 2.22386 9 2.5 9H17.5C17.7761 9 18 9.22386 18 9.5C18 9.77614 17.7761 10 17.5 10H2.5C2.22386 10 2 9.77614 2 9.5Z"
                fill="#212121">
                <title>Not tried yet</title>
            </path>
        </svg>
    );
}

export function SVGDownload({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M16.25 19.75C15.6977 19.75 15.25 20.1977 15.25 20.75V21.75C15.25 22.3023 15.6977 22.75 16.25 22.75H19.75C21.1307 22.75 22.25 21.6307 22.25 20.25V4.25C22.25 2.86929 21.1307 1.75 19.75 1.75H16.25C15.6977 1.75 15.25 2.19772 15.25 2.75V3.75C15.25 4.30228 15.6977 4.75 16.25 4.75H19.25V19.75H16.25Z"
                fill="#000000"
            />
            <path
                className="fill-current"
                d="M10.75 16.25C10.75 16.6544 10.9936 17.0191 11.3673 17.1739C11.741 17.3286 12.1711 17.2431 12.4571 16.9571L16.4571 12.9571C16.8476 12.5666 16.8476 11.9334 16.4571 11.5429L12.4571 7.54286C12.1711 7.25687 11.741 7.17131 11.3673 7.32609C10.9936 7.48087 10.75 7.84551 10.75 8.24997V10.75H2.75C2.19771 10.75 1.75 11.1977 1.75 11.75V12.75C1.75 13.3023 2.19772 13.75 2.75 13.75H10.75V16.25Z"
                fill="#000000"
            />
        </svg>
    );
}

export function SVGView({ cls = null }) {
    return (
        <svg
            className={`${cls}`}
            viewBox="-3.5 0 32 32"
            version="1.1"
            xmlns="http://www.w3.org/2000/svg">
            <path
                className="fill-current"
                d="M12.406 13.844c1.188 0 2.156 0.969 2.156 2.156s-0.969 2.125-2.156 2.125-2.125-0.938-2.125-2.125 0.938-2.156 2.125-2.156zM12.406 8.531c7.063 0 12.156 6.625 12.156 6.625 0.344 0.438 0.344 1.219 0 1.656 0 0-5.094 6.625-12.156 6.625s-12.156-6.625-12.156-6.625c-0.344-0.438-0.344-1.219 0-1.656 0 0 5.094-6.625 12.156-6.625zM12.406 21.344c2.938 0 5.344-2.406 5.344-5.344s-2.406-5.344-5.344-5.344-5.344 2.406-5.344 5.344 2.406 5.344 5.344 5.344z"
            />
        </svg>
    );
}
