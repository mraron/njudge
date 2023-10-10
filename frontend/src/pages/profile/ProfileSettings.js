import { useContext, useEffect, useState } from "react"
import { Link, useNavigate, useParams } from "react-router-dom";
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import Checkbox from "../../components/input/Checkbox"
import TextBox from "../../components/input/TextBox"
import RoundedFrame, { SVGTitleComponent } from "../../components/container/RoundedFrame"
import Button from "../../components/basic/Button"
import UserContext from "../../contexts/user/UserContext"
import { routeMap } from "../../config/RouteConfig"
import { changePassword, saveSettings } from "../../util/settings"

function PasswordChangeFrame() {
    const { t } = useTranslation()
    const { user } = useParams()
    const [oldPw, setOldPw] = useState("")
    const [newPw, setNewPw] = useState("")
    const [newPwConfirm, setNewPwConfirm] = useState("")
    const handleChangeOldPw = (newText) => setOldPw(newText)
    const handleChangeNewPw = (newText) => setNewPw(newText)
    const handleChangeNewPwConfirm = (newText) => setNewPwConfirm(newText)
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-lock" className="w-4 h-4 mr-3" />}
            title={t("profile_settings.password_change")}
        />
    )
    const handleChangePassword = async () => {
        changePassword(user, oldPw, newPw, newPwConfirm).then((resp) => {
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
                    <TextBox
                        id="oldPassword"
                        label={t("profile_settings.old_password")}
                        type="password"
                        initText={oldPw}
                        onChange={handleChangeOldPw}
                    />
                </div>
                <div className="mb-4 w-full">
                    <TextBox
                        id="newPassword"
                        label={t("profile_settings.new_password")}
                        type="password"
                        initText={newPw}
                        onChange={handleChangeNewPw}
                    />
                </div>
                <div className="mb-2 w-full">
                    <TextBox
                        id="newPasswordConfirm"
                        label={t("profile_settings.confirm_password")}
                        type="password"
                        initText={newPwConfirm}
                        onChange={handleChangeNewPwConfirm}
                    />
                </div>
                <div className="mb-6">
                    <Link to={routeMap.forgotten_password} className="link text-sm">
                        {t("profile_settings.forgotten_password")}
                    </Link>
                </div>
                <div className="flex justify-center">
                    <Button color="indigo" onClick={handleChangePassword} minWidth="8rem">
                        {t("profile_settings.save")}
                    </Button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function OtherSettingsFrame({ data }) {
    const { t } = useTranslation()
    const { user } = useParams()
    const [showUnsolved, setShowUnsolved] = useState(data.showUnsolved)
    const [hideSolved, setHideSolved] = useState(data.hideSolved)
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-cog" className="w-4 h-4 mr-3" />}
            title={t("profile_settings.other_settings")}
        />
    )
    const handleSaveSettings = async () => {
        saveSettings(user, showUnsolved, hideSolved).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_settings_save", "success")
            } else {
                window.flash("flash.unsuccessful_settings_save", "failure")
            }
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full">
                <div className="mb-3">
                    <Checkbox
                        id={"showUnsolved"}
                        label={t("profile_settings.show_unsolved")}
                        initChecked={showUnsolved}
                        onChange={setShowUnsolved}
                    />
                </div>
                <div className="mb-6">
                    <Checkbox
                        id={"hideSolved"}
                        label={t("profile_settings.hide_solved")}
                        initChecked={hideSolved}
                        onChange={setHideSolved}
                    />
                </div>
                <div className="flex justify-center">
                    <Button color="indigo" onClick={handleSaveSettings} minWidth="8rem">
                        {t("profile_settings.save")}
                    </Button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function ProfileSettings({ data }) {
    const navigate = useNavigate()
    const [isVisible, setVisible] = useState(false)
    const { user } = useParams()
    const { userData, isLoggedIn } = useContext(UserContext)

    useEffect(() => {
        if (!isLoggedIn || userData.username !== user) {
            navigate(routeMap.home)
            window.flash("flash.no_permission", "failure")
        } else {
            setVisible(true)
        }
    }, [])
    return (
        isVisible && (
            <div className="flex flex-col lg:flex-row w-full items-start space-y-3 lg:space-y-0 lg:space-x-3">
                <div className="w-full lg:w-96 min-w-0 shrink-0">
                    <PasswordChangeFrame />
                </div>
                <div className="w-full min-w-0">
                    <OtherSettingsFrame data={data} />
                </div>
            </div>
        )
    )
}

export default ProfileSettings
