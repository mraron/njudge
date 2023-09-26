import { useTranslation } from "react-i18next";
import { useContext } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTags } from "@fortawesome/free-solid-svg-icons";
import Modal from "./Modal";
import RoundedFrame from "../RoundedFrame";
import TagDropdown from "../../input/TagDropdown";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import JudgeDataContext from "../../../contexts/judgeData/JudgeDataContext";

function TagModal({isOpen, onClose}) {
    const { t } = useTranslation()
    const { judgeData } = useContext(JudgeDataContext)
    const titleComponent = (
        <SVGTitleComponent svg={<FontAwesomeIcon icon={faTags} className="w-5 h-5 mr-2" />} title={t("tag_modal.title")} />
    )
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <RoundedFrame titleComponent={titleComponent} cls="shadow-md">
                <div className="w-96 px-6 py-5">
                    <div className="mb-5">
                        <TagDropdown itemNames={judgeData.tags.map(t)} initTags={[0, 1]} />
                    </div>
                    <div className="flex justify-center">
                        <button className="w-full btn-indigo padding-btn-default mr-2">{t("tag_modal.save")}</button>
                        <button className="w-full btn-gray padding-btn-default" onClick={onClose}>{t("tag_modal.cancel")}</button>
                    </div>
                </div>
            </RoundedFrame>
        </Modal>
    )
}

export default TagModal