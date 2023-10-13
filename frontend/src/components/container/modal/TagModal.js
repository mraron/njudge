import { useTranslation } from "react-i18next"
import { useContext } from "react"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import AnimatedModal from "./AnimatedModal"
import RoundedFrame, { SVGTitleComponent } from "../RoundedFrame"
import TagDropdown from "../../input/TagDropdown"
import JudgeDataContext from "../../../contexts/judgeData/JudgeDataContext"
import Button from "../../basic/Button"

function TagModal({ isOpen, onClose }) {
    const { t } = useTranslation()
    const { judgeData } = useContext(JudgeDataContext)
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-tags" className="w-5 h-5 mr-3" />}
            title={t("tag_modal.title")}
        />
    )
    return (
        <AnimatedModal isOpen={isOpen} onClose={onClose}>
            <RoundedFrame titleComponent={titleComponent} cls="shadow-md w-full md:w-96">
                <div className="px-6 py-5">
                    <div className="mb-5">
                        <TagDropdown items={judgeData.tags.map(t)} initTags={[0, 1]} />
                    </div>
                    <div className="flex justify-center space-x-2">
                        <Button color="indigo" fullWidth={true}>
                            {t("tag_modal.save")}
                        </Button>
                        <Button color="gray" fullWidth={true} onClick={onClose}>
                            {t("tag_modal.cancel")}
                        </Button>
                    </div>
                </div>
            </RoundedFrame>
        </AnimatedModal>
    )
}

export default TagModal
