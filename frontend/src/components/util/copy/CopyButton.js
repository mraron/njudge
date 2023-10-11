import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { TERipple } from "tw-elements-react";

function CopyButton({ text, isVisible }) {
    const { t } = useTranslation()

    const handleCopy = () => {
        navigator.clipboard.writeText(text)
        window.flash("info.successful_copy", "success")
    }
    return (
        isVisible && (
            <TERipple
                className="rounded-lg bg-grey-825 hover:bg-grey-800 border border-borxstrcol overflow-hidden"
                rippleColor="#808080">
                <button
                    className={`rounded-lg text-grey-200 flex justify-center items-center p-2`}
                    aria-label={t("aria_label.copy")}
                    onClick={handleCopy}>
                    <FontAwesomeIcon icon="fa-regular fa-copy" className="w-3.5 h-3.5" />
                </button>
            </TERipple>
        )
    )
}

export default CopyButton
