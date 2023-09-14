import {SVGSpinner} from "../svg/SVGs";

function PageLoadingAnimation() {
    return (
        <div className="flex justify-center">
            <SVGSpinner cls="w-12 h-12 mx-12 my-16" />
        </div>
    )
}

export default PageLoadingAnimation