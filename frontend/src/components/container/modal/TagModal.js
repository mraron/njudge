import { useTranslation } from "react-i18next";
import { useContext } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Modal from "./Modal";
import RoundedFrame from "../RoundedFrame";
import TagDropdown from "../../input/TagDropdown";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import JudgeDataContext from "../../../contexts/judgeData/JudgeDataContext";
import Button from "../../util/Button";

function TagModal({ isOpen, onClose }) {
    const { t } = useTranslation();
    const { judgeData } = useContext(JudgeDataContext);
    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon="fa-tags" className="w-4 h-4 mr-2" />}
            title={t("tag_modal.title")}
        />
    );
    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <RoundedFrame titleComponent={titleComponent} cls="shadow-md">
                <div className="w-96 px-6 py-5">
                    <div className="mb-5">
                        <TagDropdown
                            itemNames={judgeData.tags.map(t)}
                            initTags={[0, 1]}
                        />
                    </div>
                    <div className="flex justify-center">
                        <div className="mr-2">
                            <Button theme="indigo">
                                {t("tag_modal.save")}
                            </Button>
                        </div>
                        <Button theme="gray" onClick={onClose}>
                            {t("tag_modal.save")}
                        </Button>
                    </div>
                </div>
            </RoundedFrame>
        </Modal>
    );
}

export default TagModal;
