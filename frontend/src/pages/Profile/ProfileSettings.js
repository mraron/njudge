import { useState } from "react";

import Checkbox from "../../components/Checkbox";
import RoundedFrame from "../../components/RoundedFrame";
import TextBox from "../../components/TextBox"
import {SVGChange, SVGConfirm, SVGCorrectSimple, SVGLock, SVGSettings} from "../../svg/SVGs";
import SVGTitleComponent from "../../svg/SVGTitleComponent";

function PasswordChangeFrame() {
    const [oldPw, setOldPw] = useState("");
    const [newPw, setNewPw] = useState("");
    const [newPwConfirm, setNewPwConfirm] = useState("");

    const handleChangeOldPw = (newText) => setOldPw(newText);
    const handleChangeNewPw = (newText) => setNewPw(newText);
    const handleChangeNewPwConfirm = (newText) => setNewPwConfirm(newText);
    const titleComponent = <SVGTitleComponent svg={<SVGChange cls="w-6 h-6 mr-2" />} title="Jelszó megváltoztatása" />

    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-4 w-full">
                    <TextBox id="oldPassword" label="Régi jelszó" type="password" initText={oldPw} onChange={handleChangeOldPw} />
                </div>
                <div className="mb-4 w-full">
                    <TextBox id="newPassword" label="Új jelszó" type="password" initText={newPw} onChange={handleChangeNewPw} />
                </div>
                <div className="mb-6 w-full">
                    <TextBox id="newPasswordConfirm" label="Új jelszó megerősítése" type="password" initText={newPwConfirm} onChange={handleChangeNewPwConfirm} />
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo w-32">Mentés</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function OtherSettingsFrame() {
    const titleComponent = <SVGTitleComponent svg={<SVGSettings cls="w-5 h-5 mr-2" />} title="Egyéb beállítások" />
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-2">
                    <Checkbox id={"showTagsUnsolved"} label="Megoldatlan feladatok címkéinek mutatása"></Checkbox>
                </div>
                <div className="mb-6">
                    <Checkbox id={"hideSolved"} label="Megoldott feladatok elrejtése"></Checkbox>
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo w-32">Mentés</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function ProfileSettings() {
    return (
        <div className="flex flex-col md:flex-row w-full items-start">
            <div className="w-full md:w-96 mb-3 shrink-0">
                <PasswordChangeFrame />
            </div>
            <div className="w-full mb-3 md:ml-3">
                <OtherSettingsFrame />
            </div>
        </div>
    );
}

export default ProfileSettings;