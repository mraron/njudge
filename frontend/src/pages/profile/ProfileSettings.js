import React, {useContext, useEffect, useState} from "react";
import Checkbox from "../../components/input/Checkbox";
import RoundedFrame from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox"
import {SVGLock, SVGSettings} from "../../svg/SVGs";
import SVGTitleComponent from "../../svg/SVGTitleComponent";
import {useNavigate, useParams} from "react-router-dom";
import UserContext from "../../contexts/user/UserContext";
import {routeMap} from "../../config/RouteConfig";
import {useTranslation} from "react-i18next";

function PasswordChangeFrame() {
    const {t} = useTranslation()
    const [oldPw, setOldPw] = useState("");
    const [newPw, setNewPw] = useState("");
    const [newPwConfirm, setNewPwConfirm] = useState("")
    const handleChangeOldPw = (newText) => setOldPw(newText);
    const handleChangeNewPw = (newText) => setNewPw(newText);
    const handleChangeNewPwConfirm = (newText) => setNewPwConfirm(newText);
    const titleComponent = <SVGTitleComponent svg={<SVGLock cls="w-5 h-5 mr-2"/>}
                                              title={t("profile_settings.password_change")}/>

    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-4 w-full">
                    <TextBox id="oldPassword" label={t("profile_settings.old_password")} type="password" initText={oldPw}
                             onChange={handleChangeOldPw}/>
                </div>
                <div className="mb-4 w-full">
                    <TextBox id="newPassword" label={t("profile_settings.new_password")} type="password" initText={newPw}
                             onChange={handleChangeNewPw}/>
                </div>
                <div className="mb-6 w-full">
                    <TextBox id="newPasswordConfirm" label={t("profile_settings.confirm_password")} type="password"
                             initText={newPwConfirm} onChange={handleChangeNewPwConfirm}/>
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo w-32">{t("profile_settings.save")}</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function OtherSettingsFrame() {
    const {t} = useTranslation()
    const titleComponent = <SVGTitleComponent svg={<SVGSettings cls="w-5 h-5 mr-2"/>}
                                              title={t("profile_settings.other_settings")}/>
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-3">
                    <Checkbox id={"showTagsUnsolved"} label={t("profile_settings.unsolved_tags")}></Checkbox>
                </div>
                <div className="mb-6">
                    <Checkbox id={"hideSolved"} label={t("profile_settings.hide_solved")}></Checkbox>
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo w-32">{t("profile_settings.save")}</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function ProfileSettings({ data }) {
    const {t} = useTranslation()
    const navigate = useNavigate()
    const [isVisible, setVisible] = useState(false)
    const {user} = useParams()
    const {userData, isLoggedIn} = useContext(UserContext)

    useEffect(() => {
        if (!isLoggedIn || userData.username !== user) {
            navigate(routeMap.home)
            window.flash(t("flash.no_permission"), "failure")
        } else {
            setVisible(true)
        }
    }, [])
    return (
        isVisible &&
        <div className="flex flex-col lg:flex-row w-full items-start">
            <div className="w-full lg:w-96 mb-3 shrink-0">
                <PasswordChangeFrame/>
            </div>
            <div className="w-full mb-3 lg:ml-3">
                <OtherSettingsFrame/>
            </div>
        </div>
    );
}

export default ProfileSettings;