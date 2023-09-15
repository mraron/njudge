
export function SVGAttachmentFile() {
    return (
        <svg className="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="stroke-current" d="M19 9V17.8C19 18.9201 19 19.4802 18.782 19.908C18.5903 20.2843 18.2843 20.5903 17.908 20.782C17.4802 21 16.9201 21 15.8 21H8.2C7.07989 21 6.51984 21 6.09202 20.782C5.71569 20.5903 5.40973 20.2843 5.21799 19.908C5 19.4802 5 18.9201 5 17.8V6.2C5 5.07989 5 4.51984 5.21799 4.09202C5.40973 3.71569 5.71569 3.40973 6.09202 3.21799C6.51984 3 7.0799 3 8.2 3H13M19 9L13 3M19 9H14C13.4477 9 13 8.55228 13 8V3" stroke="#000000" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGAttachmentDescription() {
    return (
        <svg  className="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="stroke-current" d="M9 17H15M9 13H15M9 9H10M13 3H8.2C7.0799 3 6.51984 3 6.09202 3.21799C5.71569 3.40973 5.40973 3.71569 5.21799 4.09202C5 4.51984 5 5.0799 5 6.2V17.8C5 18.9201 5 19.4802 5.21799 19.908C5.40973 20.2843 5.71569 20.5903 6.09202 20.782C6.51984 21 7.0799 21 8.2 21H15.8C16.9201 21 17.4802 21 17.908 20.782C18.2843 20.5903 18.5903 20.2843 18.782 19.908C19 19.4802 19 18.9201 19 17.8V9M13 3L19 9M13 3V7.4C13 7.96005 13 8.24008 13.109 8.45399C13.2049 8.64215 13.3578 8.79513 13.546 8.89101C13.7599 9 14.0399 9 14.6 9H19" stroke="#000000" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGInformation() {
    return (
        <svg className="w-6 h-6 mr-2" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="fill-white" clipRule="evenodd" d="m12 3.75c-4.55635 0-8.25 3.69365-8.25 8.25 0 4.5563 3.69365 8.25 8.25 8.25 4.5563 0 8.25-3.6937 8.25-8.25 0-4.55635-3.6937-8.25-8.25-8.25zm-9.75 8.25c0-5.38478 4.36522-9.75 9.75-9.75 5.3848 0 9.75 4.36522 9.75 9.75 0 5.3848-4.3652 9.75-9.75 9.75-5.38478 0-9.75-4.3652-9.75-9.75zm9.75-.75c.4142 0 .75.3358.75.75v3.5c0 .4142-.3358.75-.75.75s-.75-.3358-.75-.75v-3.5c0-.4142.3358-.75.75-.75zm0-3.25c-.5523 0-1 .44772-1 1s.4477 1 1 1h.01c.5523 0 1-.44772 1-1s-.4477-1-1-1z" fill="#000000" fillRule="evenodd"/>
        </svg>
    )
}

export function SVGSubmit() {
    return (
        <svg className="w-6 h-6 mr-2" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="stroke-white" d="M10 14L13 21L20 4L3 11L6.5 12.5" stroke="#000000" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGAttachment() {
    return (
        <svg className="w-6 h-6 mr-2" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="stroke-white" d="M20 10.9696L11.9628 18.5497C10.9782 19.4783 9.64274 20 8.25028 20C6.85782 20 5.52239 19.4783 4.53777 18.5497C3.55315 17.6211 3 16.3616 3 15.0483C3 13.7351 3.55315 12.4756 4.53777 11.547M14.429 6.88674L7.00403 13.8812C6.67583 14.1907 6.49144 14.6106 6.49144 15.0483C6.49144 15.4861 6.67583 15.9059 7.00403 16.2154C7.33224 16.525 7.77738 16.6989 8.24154 16.6989C8.70569 16.6989 9.15083 16.525 9.47904 16.2154L13.502 12.4254M8.55638 7.75692L12.575 3.96687C13.2314 3.34779 14.1217 3 15.05 3C15.9783 3 16.8686 3.34779 17.525 3.96687C18.1814 4.58595 18.5502 5.4256 18.5502 6.30111C18.5502 7.17662 18.1814 8.01628 17.525 8.63535L16.5 9.601" stroke="#000000" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function DropdownListArrow({ isOpen }) {
    return (
        <svg className={`w-4 h-4 mr-2 fill-current ${isOpen? "rotate-90": ""} transition-all duration-100 shrink-0`} viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
            <rect width="48" height="48" fill="none"/>
            <path d="M19.5,37.4l11.9-12a1.9,1.9,0,0,0,0-2.8l-11.9-12A2,2,0,0,0,16,12h0V36h0a2,2,0,0,0,3.5,1.4Z"/>
        </svg>
    )
}

export function SVGDropdownMenuArrow({ isOpen }) {
    return (
        <svg className={`w-6 h-6 ml-4 ${isOpen? "rotate-180": ""} text-white transition-all duration-150`} viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M7 10L12 15L17 10" stroke="#ffffff" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGCopy() {
    return (
        <svg className="w-4 h-4 fill-white" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="fill-white" fillRule="evenodd" clipRule="evenodd" d="M21 8C21 6.34315 19.6569 5 18 5H10C8.34315 5 7 6.34315 7 8V20C7 21.6569 8.34315 23 10 23H18C19.6569 23 21 21.6569 21 20V8ZM19 8C19 7.44772 18.5523 7 18 7H10C9.44772 7 9 7.44772 9 8V20C9 20.5523 9.44772 21 10 21H18C18.5523 21 19 20.5523 19 20V8Z" fill="#0F0F0F"/>
            <path className="fill-white" d="M6 3H16C16.5523 3 17 2.55228 17 2C17 1.44772 16.5523 1 16 1H6C4.34315 1 3 2.34315 3 4V18C3 18.5523 3.44772 19 4 19C4.55228 19 5 18.5523 5 18V4C5 3.44772 5.44772 3 6 3Z" fill="#0F0F0F"/>
        </svg>
    )
}

export function SVGAvatar({ cls }) {
    return (
        <svg className={`${cls}`} viewBox="0 0 20 20" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
            <g stroke="none" strokeWidth="1" fillRule="evenodd">
                <g transform="translate(-140.000000, -2159.000000)" className="fill-current">
                    <g transform="translate(56.000000, 160.000000)">
                        <path d="M100.562548,2016.99998 L87.4381713,2016.99998 C86.7317804,2016.99998 86.2101535,2016.30298 86.4765813,2015.66198 C87.7127655,2012.69798 90.6169306,2010.99998 93.9998492,2010.99998 C97.3837885,2010.99998 100.287954,2012.69798 101.524138,2015.66198 C101.790566,2016.30298 101.268939,2016.99998 100.562548,2016.99998 M89.9166645,2004.99998 C89.9166645,2002.79398 91.7489936,2000.99998 93.9998492,2000.99998 C96.2517256,2000.99998 98.0830339,2002.79398 98.0830339,2004.99998 C98.0830339,2007.20598 96.2517256,2008.99998 93.9998492,2008.99998 C91.7489936,2008.99998 89.9166645,2007.20598 89.9166645,2004.99998 M103.955674,2016.63598 C103.213556,2013.27698 100.892265,2010.79798 97.837022,2009.67298 C99.4560048,2008.39598 100.400241,2006.33098 100.053171,2004.06998 C99.6509769,2001.44698 97.4235996,1999.34798 94.7348224,1999.04198 C91.0232075,1998.61898 87.8750721,2001.44898 87.8750721,2004.99998 C87.8750721,2006.88998 88.7692896,2008.57398 90.1636971,2009.67298 C87.1074334,2010.79798 84.7871636,2013.27698 84.044024,2016.63598 C83.7745338,2017.85698 84.7789973,2018.99998 86.0539717,2018.99998 L101.945727,2018.99998 C103.221722,2018.99998 104.226185,2017.85698 103.955674,2016.63598">
                        </path>
                    </g>
                </g>
            </g>
        </svg>
    )
}

export function SVGHamburger() {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
    )
}

export function SVGClose({ size }) {
    return (
        <svg style={{ pointerEvents: "none" }} fill="#000000" className={`${size} fill-white`} version="1.1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" viewBox="0 0 460.775 460.775" xmlSpace="preserve">
            <path d="M285.08,230.397L456.218,59.27c6.076-6.077,6.076-15.911,0-21.986L423.511,4.565c-2.913-2.911-6.866-4.55-10.992-4.55
                c-4.127-4.127-8.08-4.55-10.993-4.55c-4.127,0-8.08,1.639-10.993,4.55l-171.138,171.14L59.25,4.565c-2.913-2.911-6.866-4.55-10.993-4.55
                c-4.126,0-8.08,1.639-10.992,4.55L4.558,37.284c-6.077,6.075-6.077,15.909,0,21.986l171.138,171.128L4.575,401.505
                c-6.074,6.077-6.074,15.911,0,21.986l32.709,32.719c2.911,2.911,6.865,4.55,10.992,4.55c4.127,0,8.08-1.639,10.994-4.55
                l171.117-171.12l171.118,171.12c2.913,2.911,6.866,4.55,10.993,4.55c4.128,0,8.081-1.639,10.992-4.55l32.709-32.719
                c6.074-6.075,6.074-15.909,0-21.986L285.08,230.397z"/>
        </svg>
    )
}

export function SVGDropdownFilterArrow({ isOpen }) {
    return (
        <svg className={`${isOpen? "rotate-180": "rotate-0"} text-white h-10 w-16 text-sm font-medium transition duration-150`} viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M7 10L12 15L17 10" stroke="#ffffff" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGNotFound() {
    return (
        <svg className="w-80" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M9 17C9.85038 16.3697 10.8846 16 12 16C13.1154 16 14.1496 16.3697 15 17" stroke="#1C274C" strokeWidth="1.5" strokeLinecap="round"  className="stroke-grey-750" />
            <ellipse cx="15" cy="10.5" rx="1" ry="1.5" fill="#1C274C" className="fill-grey-750" />
            <ellipse cx="9" cy="10.5" rx="1" ry="1.5" fill="#1C274C" className="fill-grey-750" />
            <path d="M22 12C22 16.714 22 19.0711 20.5355 20.5355C19.0711 22 16.714 22 12 22C7.28595 22 4.92893 22 3.46447 20.5355C2 19.0711 2 16.714 2 12C2 7.28595 2 4.92893 3.46447 3.46447C4.92893 2 7.28595 2 12 2C16.714 2 19.0711 2 20.5355 3.46447C21.5093 4.43821 21.8356 5.80655 21.9449 8" stroke="#1C274C" strokeWidth="1.5" strokeLinecap="round"  className="stroke-grey-750" />
        </svg>
    )
}

export function SVGStatistics({ cls }) {
    return (
        <svg className={`${cls} fill-current`} fill="#000000" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 100 100" enableBackground="new 0 0 100 100" xmlSpace="preserve">
            <path d="M46.05,60.163H31.923c-0.836,0-1.513,0.677-1.513,1.513v21.934c0,0.836,0.677,1.513,1.513,1.513H46.05
                c0.836,0,1.512-0.677,1.512-1.513V61.675C47.562,60.839,46.885,60.163,46.05,60.163z"/>
            <path d="M68.077,14.878H53.95c-0.836,0-1.513,0.677-1.513,1.513v67.218c0,0.836,0.677,1.513,1.513,1.513h14.127
                c0.836,0,1.513-0.677,1.513-1.513V16.391C69.59,15.555,68.913,14.878,68.077,14.878z"/>
            <path d="M90.217,35.299H76.09c-0.836,0-1.513,0.677-1.513,1.513v46.797c0,0.836,0.677,1.513,1.513,1.513h14.126
                c0.836,0,1.513-0.677,1.513-1.513V36.812C91.729,35.977,91.052,35.299,90.217,35.299z"/>
            <path d="M23.91,35.299H9.783c-0.836,0-1.513,0.677-1.513,1.513v46.797c0,0.836,0.677,1.513,1.513,1.513H23.91
                c0.836,0,1.513-0.677,1.513-1.513V36.812C25.423,35.977,24.746,35.299,23.91,35.299z"/>
        </svg>
    )
}

export function SVGRecent() {
    return (
        <svg className="w-6 h-6 mr-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 30" version="1.1">
            <g transform="translate(0,-289.0625)">
                <path className="fill-current" d="M 15 3 C 8.3844276 3 3 8.38443 3 15 C 3 21.61557 8.3844276 27 15 27 C 21.615572 27 27 21.61557 27 15 C 27 8.38443 21.615572 3 15 3 z M 15 5 C 20.534692 5 25 9.46531 25 15 C 25 20.53469 20.534692 25 15 25 C 9.4653079 25 5 20.53469 5 15 C 5 9.46531 9.4653079 5 15 5 z M 15 7 C 14.446 7 14 7.446 14 8 L 14 15 C 14 15.554 14.446 16 15 16 L 22 16 C 22.554 16 23 15.554 23 15 C 23 14.446 22.554 14 22 14 L 16 14 L 16 8 C 16 7.446 15.554 7 15 7 z " transform="translate(0,289.0625)"/>
            </g>
        </svg>
    )
}

export function SVGResults() {
    return (
        <svg className="w-6 h-6 mr-2 fill-current" fill="#000000" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" 
                viewBox="0 0 100 100" enableBackground="new 0 0 100 100" xmlSpace="preserve">
                <path d="M27.953,46.506c-1.385-2.83-2.117-6.008-2.117-9.192c0-1.743,0.252-3.534,0.768-5.468c0.231-0.87,0.521-1.702,0.847-2.509
                    c-1.251-0.683-2.626-1.103-4.101-1.103c-5.47,0-9.898,5.153-9.898,11.517c0,4.452,2.176,8.305,5.354,10.222L5.391,56.217
                    c-0.836,0.393-1.387,1.337-1.387,2.392v10.588c0,1.419,0.991,2.569,2.21,2.569h7.929V60.656c0-3.237,1.802-6.172,4.599-7.481
                    l10.262-4.779C28.624,47.792,28.273,47.161,27.953,46.506z"/>
                <path d="M60.137,34.801h34.092v-0.001c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761c0,0,0-0.001,0-0.001
                    l0-6.43h0c0-0.973-0.789-1.761-1.761-1.761c-0.002,0-0.004,0.001-0.006,0.001v-0.005H56.133c1.614,2.114,2.844,4.627,3.526,7.435
                    C59.874,33.168,60.03,33.999,60.137,34.801z"/>
                <path d="M95.996,66.436c0-0.973-0.789-1.761-1.761-1.761c-0.002,0-0.004,0.001-0.006,0.001v-0.005H72.007v7.095v1.994
                    c0,0.293-0.016,0.582-0.045,0.867h22.267v-0.001c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761l0-0.001
                    L95.996,66.436L95.996,66.436z"/>
                <path d="M94.235,44.762c-0.002,0-0.004,0.001-0.006,0.001v-0.005H58.944c-0.159,0.419-0.327,0.836-0.514,1.249
                    c-0.364,0.802-0.773,1.569-1.224,2.297l10.288,4.908c0.781,0.378,1.473,0.897,2.078,1.503h24.657v-0.001
                    c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761c0,0,0-0.001,0-0.001l0-6.43h0
                    C95.996,45.55,95.207,44.762,94.235,44.762z"/>
                <path d="M65.323,57.702l-11.551-5.51l-4.885-2.33c2.134-1.344,3.866-3.418,5-5.917c0.899-1.984,1.435-4.231,1.435-6.631
                    c0-1.348-0.213-2.627-0.512-3.863c-1.453-5.983-6.126-10.392-11.736-10.392c-5.504,0-10.106,4.251-11.648,10.065
                    c-0.356,1.333-0.602,2.72-0.602,4.189c0,2.552,0.596,4.93,1.609,7c1.171,2.4,2.906,4.379,5.018,5.651l-4.678,2.178l-11.926,5.554
                    c-1.037,0.485-1.717,1.654-1.717,2.959v11.111v1.994c0,1.756,1.224,3.181,2.735,3.181h42.417c1.511,0,2.735-1.424,2.735-3.181
                    v-1.994V60.656C67.019,59.355,66.349,58.198,65.323,57.702z"/>
        </svg>
    )
}

export function SVGGoogle() {
    return (
        <svg className="w-6 h-6" viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M30.0014 16.3109C30.0014 15.1598 29.9061 14.3198 29.6998 13.4487H16.2871V18.6442H24.1601C24.0014 19.9354 23.1442 21.8798 21.2394 23.1864L21.2127 23.3604L25.4536 26.58L25.7474 26.6087C28.4458 24.1665 30.0014 20.5731 30.0014 16.3109Z" fill="#4285F4"/>
            <path d="M16.2863 29.9998C20.1434 29.9998 23.3814 28.7553 25.7466 26.6086L21.2386 23.1863C20.0323 24.0108 18.4132 24.5863 16.2863 24.5863C12.5086 24.5863 9.30225 22.1441 8.15929 18.7686L7.99176 18.7825L3.58208 22.127L3.52441 22.2841C5.87359 26.8574 10.699 29.9998 16.2863 29.9998Z" fill="#34A853"/>
            <path d="M8.15964 18.769C7.85806 17.8979 7.68352 16.9645 7.68352 16.0001C7.68352 15.0356 7.85806 14.1023 8.14377 13.2312L8.13578 13.0456L3.67083 9.64746L3.52475 9.71556C2.55654 11.6134 2.00098 13.7445 2.00098 16.0001C2.00098 18.2556 2.55654 20.3867 3.52475 22.2845L8.15964 18.769Z" fill="#FBBC05"/>
            <path d="M16.2864 7.4133C18.9689 7.4133 20.7784 8.54885 21.8102 9.4978L25.8419 5.64C23.3658 3.38445 20.1435 2 16.2864 2C10.699 2 5.8736 5.1422 3.52441 9.71549L8.14345 13.2311C9.30229 9.85555 12.5086 7.4133 16.2864 7.4133Z" fill="#EB4335"/>
        </svg>
    )
}

export function SVGLogin() {
    return (
        <svg className="w-6 h-6 mr-2" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="fill-current" fillRule="evenodd" clipRule="evenodd" d="M8 6C8 3.79086 9.79086 2 12 2H17.5C19.9853 2 22 4.01472 22 6.5V17.5C22 19.9853 19.9853 22 17.5 22H12C9.79086 22 8 20.2091 8 18V17C8 16.4477 8.44772 16 9 16C9.55228 16 10 16.4477 10 17V18C10 19.1046 10.8954 20 12 20H17.5C18.8807 20 20 18.8807 20 17.5V6.5C20 5.11929 18.8807 4 17.5 4H12C10.8954 4 10 4.89543 10 6V7C10 7.55228 9.55228 8 9 8C8.44772 8 8 7.55228 8 7V6ZM12.2929 8.29289C12.6834 7.90237 13.3166 7.90237 13.7071 8.29289L16.7071 11.2929C17.0976 11.6834 17.0976 12.3166 16.7071 12.7071L13.7071 15.7071C13.3166 16.0976 12.6834 16.0976 12.2929 15.7071C11.9024 15.3166 11.9024 14.6834 12.2929 14.2929L13.5858 13L5 13C4.44772 13 4 12.5523 4 12C4 11.4477 4.44772 11 5 11L13.5858 11L12.2929 9.70711C11.9024 9.31658 11.9024 8.68342 12.2929 8.29289Z" fill="#0F1729"/>
        </svg>
    )
}

export function SVGLock({ cls }) {
    return (
        <svg className={`${cls} fill-current`} height="800px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path fillRule="evenodd" d="M12,2 C14.6887547,2 16.8818181,4.12230671 16.9953805,6.78311038 L17,7 L17,10 C18.6568542,10 20,11.3431458 20,13 L20,19 C20,20.6568542 18.6568542,22 17,22 L7,22 C5.34314575,22 4,20.6568542 4,19 L4,13 C4,11.3431458 5.34314575,10 7,10 L7,7 C7,4.23857625 9.23857625,2 12,2 Z M17,12 L7,12 C6.44771525,12 6,12.4477153 6,13 L6,19 C6,19.5522847 6.44771525,20 7,20 L17,20 C17.5522847,20 18,19.5522847 18,19 L18,13 C18,12.4477153 17.5522847,12 17,12 Z M12.1762728,4.00509269 L12,4 C10.4023191,4 9.09633912,5.24891996 9.00509269,6.82372721 L9,7 L9,10 L15,10 L15,7 C15,5.40231912 13.75108,4.09633912 12.1762728,4.00509269 L12,4 L12.1762728,4.00509269 Z"/>
        </svg>
    )
}

export function SVGMail({ cls }) {
    return (
        <svg className={`${cls}`} viewBox="0 -2.5 20 20" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
            <g className="fill-current" stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                <g className="fill-currentt" transform="translate(-300.000000, -922.000000)" fill="#000000">
                    <g className="fill-current" transform="translate(56.000000, 160.000000)">
                        <path className="fill-current" d="M262,764.291 L254,771.318 L246,764.281 L246,764 L262,764 L262,764.291 Z M246,775 L246,766.945 L254,773.98 L262,766.953 L262,775 L246,775 Z M244,777 L264,777 L264,762 L244,762 L244,777 Z" />
                    </g>
                </g>
            </g>
        </svg>
    )
}

export function SVGConfirm({ cls }) {
    return (
        <svg className={`${cls}`} viewBox="0 0 24 24" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
            <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                <rect fillRule="nonzero" x="0" y="0" width="24" height="24" />
                <circle className="stroke-current" stroke="#0C0310" strokeWidth="2" strokeLinecap="round" cx="12" cy="12" r="9" />
                <path className="stroke-current" d="M8.5,12.5 L10.151,14.5638 C10.3372,14.7965 10.6843,14.8157 10.895,14.605 L15.5,10" stroke="#0C0310" strokeWidth="2" strokeLinecap="round" />
            </g>
        </svg>
    )
}

export function SVGSpinner({ cls }) {
    return (
        <svg className={`${cls} text-grey-700 animate-spin-slow fill-indigo-600`} viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
            <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
        </svg>
    )
}

export function SVGLeftArrow({ cls }) {
    return (
        <svg className={`${cls} fill-current`} viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg">
            <path className="fill-current" d="M768 903.232l-50.432 56.768L256 512l461.568-448 50.432 56.768L364.928 512z" fill="#000000" />
        </svg>
    )
}

export function SVGDoubleLeftArrow({ cls }) {
    return (
        <svg className={`${cls} fill-current`} viewBox="0 0 20 20" version="1.1" xmlns="http://www.w3.org/2000/svg">
            <path className="fill-current" d="M 11 3.2910156 L 10.646484 3.6464844 L 4.2910156 10 L 10.646484 16.353516 L 11 16.708984 L 11.708984 16 L 11.353516 15.646484 L 5.7089844 10 L 11.353516 4.3535156 L 11.708984 4 L 11 3.2910156 z M 15 3.2910156 L 14.646484 3.6464844 L 8.2910156 10 L 14.646484 16.353516 L 15 16.708984 L 15.708984 16 L 15.353516 15.646484 L 9.7089844 10 L 15.353516 4.3535156 L 15.708984 4 L 15 3.2910156 z " />
        </svg>
    )
}

export function SVGDots({ cls }) {
    return (
        <svg className={`${cls}`} viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="stroke-current" d="M12 13C12.5523 13 13 12.5523 13 12C13 11.4477 12.5523 11 12 11C11.4477 11 11 11.4477 11 12C11 12.5523 11.4477 13 12 13Z" stroke="#000000" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
            <path className="stroke-current" d="M19 13C19.5523 13 20 12.5523 20 12C20 11.4477 19.5523 11 19 11C18.4477 11 18 11.4477 18 12C18 12.5523 18.4477 13 19 13Z" stroke="#000000" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
            <path className="stroke-current" d="M5 13C5.55228 13 6 12.5523 6 12C6 11.4477 5.55228 11 5 11C4.44772 11 4 11.4477 4 12C4 12.5523 4.44772 13 5 13Z" stroke="#000000" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGChange({ cls }) {
    return (
        <svg fill="#000000" className={`${cls} fill-current`} viewBox="0 0 56 56" xmlns="http://www.w3.org/2000/svg">
            <path d="M .3321 40.0118 C .3321 41.5117 1.4337 42.5430 3.0977 42.5430 L 8.4415 42.5430 C 12.4727 42.5430 14.9103 41.3711 17.8165 37.9727 L 23.0899 31.8320 L 28.3399 37.9727 C 31.2462 41.3711 33.6603 42.5664 37.7618 42.5664 L 42.0508 42.5664 L 42.0508 47.7695 C 42.0508 49.0352 42.8476 49.8320 44.1604 49.8320 C 44.7229 49.8320 45.3085 49.6211 45.7304 49.2461 L 54.5898 41.9336 C 55.6679 41.0664 55.6448 39.6602 54.5898 38.7930 L 45.7304 31.4336 C 45.3085 31.0586 44.7229 30.8477 44.1604 30.8477 C 42.8476 30.8477 42.0508 31.6445 42.0508 32.9102 L 42.0508 37.4571 L 37.8790 37.4571 C 35.4649 37.4571 33.9649 36.6836 32.0430 34.4571 L 26.4415 27.9180 L 32.0430 21.4024 C 33.9649 19.1524 35.4649 18.3789 37.8790 18.3789 L 42.0508 18.3789 L 42.0508 23.0898 C 42.0508 24.3555 42.8476 25.1524 44.1604 25.1524 C 44.7229 25.1524 45.3085 24.9414 45.7304 24.5664 L 54.5898 17.2539 C 55.6679 16.3867 55.6448 15.0039 54.5898 14.1133 L 45.7304 6.7539 C 45.3085 6.3789 44.7229 6.1680 44.1604 6.1680 C 42.8476 6.1680 42.0508 6.9649 42.0508 8.2305 L 42.0508 13.2930 L 37.7618 13.2930 C 33.6603 13.2930 31.2462 14.4883 28.3399 17.8867 L 23.0899 24.0274 L 17.8165 17.8867 C 14.9103 14.4883 12.4727 13.2930 8.4415 13.2930 L 3.0977 13.2930 C 1.4337 13.2930 .3321 14.3242 .3321 15.8477 C .3321 17.3477 1.4571 18.4024 3.0977 18.4024 L 8.5352 18.4024 C 10.8087 18.4024 12.2384 19.1758 14.1368 21.4024 L 19.7384 27.9180 L 14.1368 34.4571 C 12.2149 36.6836 10.7852 37.4571 8.5352 37.4571 L 3.0977 37.4571 C 1.4571 37.4571 .3321 38.5118 .3321 40.0118 Z"/>
        </svg>
    )
}

export function SVGSettings({ cls }) {
    return (
        <svg className={`${cls} fill-current`} fill="#000000" viewBox="0 0 1920 1920" xmlns="http://www.w3.org/2000/svg">
            <path d="M1703.534 960c0-41.788-3.84-84.48-11.633-127.172l210.184-182.174-199.454-340.856-265.186 88.433c-66.974-55.567-143.323-99.389-223.85-128.415L1158.932 0h-397.78L706.49 269.704c-81.43 29.138-156.423 72.282-223.962 128.414l-265.073-88.32L18 650.654l210.184 182.174C220.39 875.52 216.55 918.212 216.55 960s3.84 84.48 11.633 127.172L18 1269.346l199.454 340.856 265.186-88.433c66.974 55.567 143.322 99.389 223.85 128.415L761.152 1920h397.779l54.663-269.704c81.318-29.138 156.424-72.282 223.963-128.414l265.073 88.433 199.454-340.856-210.184-182.174c7.793-42.805 11.633-85.497 11.633-127.285m-743.492 395.294c-217.976 0-395.294-177.318-395.294-395.294 0-217.976 177.318-395.294 395.294-395.294 217.977 0 395.294 177.318 395.294 395.294 0 217.976-177.317 395.294-395.294 395.294" fillRule="evenodd"/>
        </svg>
    )
}

export function SVGWrong({ cls }) {
    return (
        <svg fill="#000000" className={`${cls} fill-current`} viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
            <path d="M100,15a85,85,0,1,0,85,85A84.93,84.93,0,0,0,100,15Zm0,150a65,65,0,1,1,65-65A64.87,64.87,0,0,1,100,165Z"/>
            <path d="M128.5,74a9.67,9.67,0,0,0-14,0L100,88.5l-14-14a9.9,9.9,0,0,0-14,14l14,14-14,14a9.9,9.9,0,0,0,14,14l14-14,14,14a9.9,9.9,0,0,0,14-14l-14-14,14-14A10.77,10.77,0,0,0,128.5,74Z"/>
        </svg>
    )
}

export function SVGCorrect({ cls }) {
    return <SVGConfirm cls={cls} />
}

export function SVGCode({ cls }) {
    return (
        <svg className={`${cls} fill-current`} viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="fill-current" fillRule="evenodd" clipRule="evenodd" d="M14.9523 6.2635L10.4523 18.2635L9.04784 17.7368L13.5478 5.73682L14.9523 6.2635ZM19.1894 12.0001L15.9698 8.78042L17.0304 7.71976L21.3108 12.0001L17.0304 16.2804L15.9698 15.2198L19.1894 12.0001ZM8.03032 15.2198L4.81065 12.0002L8.03032 8.78049L6.96966 7.71983L2.68933 12.0002L6.96966 16.2805L8.03032 15.2198Z" fill="#080341"/>
        </svg>
    )
}

export function SVGWrongSimple({ cls }) {
    return (
        <svg className={`${cls} fill-current`} fill="#000000" viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
            <path d="M114,100l49-49a9.9,9.9,0,0,0-14-14L100,86,51,37A9.9,9.9,0,0,0,37,51l49,49L37,149a9.9,9.9,0,0,0,14,14l49-49,49,49a9.9,9.9,0,0,0,14-14Z"/>
        </svg>
    )
}

export function SVGCorrectSimple({ cls }) {
    return (
        <svg className={`${cls}`} viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path className="stroke-current" d="M4 12.6111L8.92308 17.5L20 6.5" stroke="#000000" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
    )
}

export function SVGKey({ cls }) {
    return (
        <svg className={`${cls}`} version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink"
             viewBox="796 796 200 200" enableBackground="new 796 796 200 200" xmlSpace="preserve">
            <path className="fill-current" d="M975.253,834.755l-26.093-26.092C940.994,800.497,930.138,796,918.591,796c-11.549,0-22.405,4.498-30.571,12.664
                c-16.855,16.855-16.856,44.282,0,61.14l1.594,1.594c1.717,1.718,1.717,4.502,0,6.22l-82.075,82.074
                c-4.608,4.609-4.608,12.078,0,16.688c2.304,2.303,5.323,3.455,8.343,3.455c3.02,0,6.04-1.152,8.343-3.455l3.724-3.725l20.121,20.119
                c4.301,4.303,11.272,4.303,15.573,0c4.301-4.299,4.301-11.271,0-15.572l-4.451-4.451c-1.528-1.527-2.387-3.6-2.387-5.762
                s0.858-4.234,2.387-5.762l0.446-0.445c3.182-3.184,8.341-3.184,11.522,0l4.452,4.449c4.301,4.303,11.273,4.303,15.574,0
                c4.301-4.299,4.3-11.273,0-15.572l-20.121-20.119l35.234-35.236c1.718-1.718,4.503-1.718,6.221,0l1.592,1.593
                c8.166,8.165,19.022,12.663,30.569,12.663s22.404-4.498,30.569-12.663c8.167-8.167,12.666-19.024,12.666-30.572
                C987.915,853.777,983.418,842.92,975.253,834.755z M958.567,879.209c-3.711,3.709-8.641,5.751-13.886,5.751
                s-10.175-2.042-13.884-5.751l-26.092-26.092c-7.656-7.656-7.656-20.112,0-27.768c3.708-3.709,8.64-5.752,13.885-5.752
                s10.175,2.042,13.884,5.751l26.093,26.092c3.708,3.709,5.752,8.64,5.752,13.884S962.275,875.5,958.567,879.209z"/>
        </svg>
    )
}

export function SVGCheckmark({ cls }) {
    return (
        <svg className={`${cls}`} version="1.1" baseProfile="tiny" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink"
             viewBox="0 0 42 42" xmlSpace="preserve">
            <path className="fill-current" d="M39.04,7.604l-2.398-1.93c-1.182-0.95-1.869-0.939-2.881,0.311L16.332,27.494l-8.111-6.739
                c-1.119-0.94-1.819-0.89-2.739,0.26l-1.851,2.41c-0.939,1.182-0.819,1.853,0.291,2.78l11.56,9.562c1.19,1,1.86,0.897,2.78-0.222
                l21.079-25.061C40.331,9.294,40.271,8.583,39.04,7.604z"/>
        </svg>
    )
}