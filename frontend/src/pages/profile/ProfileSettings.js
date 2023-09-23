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
import {changePassword, saveSettings} from "../../util/settings";

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
    const handleChangePassword = async () => {
        changePassword(oldPw, newPw, newPwConfirm).then(resp => {
            if (resp.success) {
                window.flash("flash.successful_password_change", "success")
            } else {
                window.flash(resp.message, "failure")
            }
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-4 w-full">
                    <TextBox id="oldPassword" label={t("profile_settings.old_password")} type="password"
                             initText={oldPw}
                             onChange={handleChangeOldPw}/>
                </div>
                <div className="mb-4 w-full">
                    <TextBox id="newPassword" label={t("profile_settings.new_password")} type="password"
                             initText={newPw}
                             onChange={handleChangeNewPw}/>
                </div>
                <div className="mb-6 w-full">
                    <TextBox id="newPasswordConfirm" label={t("profile_settings.confirm_password")} type="password"
                             initText={newPwConfirm} onChange={handleChangeNewPwConfirm}/>
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo padding-btn-default w-32"
                            onClick={handleChangePassword}>{t("profile_settings.save")}</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function OtherSettingsFrame({data}) {
    const {t} = useTranslation()
    const [showUnsolved, setShowUnsolved] = useState(data.showUnsolved)
    const [hideSolved, setHideSolved] = useState(data.hideSolved)
    const titleComponent = <SVGTitleComponent svg={<SVGSettings cls="w-5 h-5 mr-2"/>}
                                              title={t("profile_settings.other_settings")}/>
    const handleSaveSettings = async () => {
        saveSettings(showUnsolved, hideSolved).then(resp => {
            if (resp.success) {
                window.flash(t("flash.successful_settings_save"), "success")
            } else {
                window.flash(t("flash.unsuccessful_settings_save"), "failure")
            }
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-3">
                    <Checkbox id={"showUnsolved"} label={t("profile_settings.show_unsolved")} initChecked={showUnsolved}
                              onChange={setShowUnsolved}/>
                </div>
                <div className="mb-6">
                    <Checkbox id={"hideSolved"} label={t("profile_settings.hide_solved")} initChecked={hideSolved}
                              onChange={setHideSolved}/>
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo padding-btn-default w-32"
                            onClick={handleSaveSettings}>{t("profile_settings.save")}</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function ProfileSettings({data}) {
    const navigate = useNavigate()
    const [isVisible, setVisible] = useState(false)
    const {user} = useParams()
    const {userData, isLoggedIn} = useContext(UserContext)

    useEffect(() => {
        if (!isLoggedIn || userData.username !== user) {
            navigate(routeMap.home)
            window.flash("flash.no_permission", "failure")
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
                <OtherSettingsFrame data={data}/>
            </div>
        </div>
    );
}

export default ProfileSettings;