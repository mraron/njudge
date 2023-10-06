import { useContext, useState } from "react"
import { Navigate } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import RoundedFrame, {
    SVGTitleComponent,
} from "../../components/container/RoundedFrame"
import TextBox from "../../components/input/TextBox"
import Button from "../../components/basic/Button"
import UserContext from "../../contexts/user/UserContext"
import { routeMap } from "../../config/RouteConfig"
import { change_password } from "../../util/auth"

function ForgottenPasswordFrame() {
    const { t } = useTranslation()
    const [email, setEmail] = useState("")
    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon="fa-lock" className="w-5 h-5 mr-3" />}
            title={t("forgotten_password.change_password")}
        />
    )
    const handleChangePassword = (event) => {
        event.preventDefault()
        change_password(email).then((ok) => {
            if (ok) {
                window.flash("flash.successful_email_pw_change", "success")
            } else {
                window.flash("flash.unsuccessful_email_pw_change", "failure")
            }
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 py-8">
                <form method="POST">
                    <div className="mb-6">
                        <TextBox
                            id="email"
                            label={t("forgotten_password.email")}
                            initText={email}
                            onChange={setEmail}
                        />
                    </div>
                    <div className="mb-2 flex justify-center">
                        <Button
                            type="submit"
                            color="indigo"
                            cls="py-2.5"
                            onClick={handleChangePassword}
                            minWidth="12rem">
                            {t("forgotten_password.change_password")}
                        </Button>
                    </div>
                </form>
            </div>
        </RoundedFrame>
    )
}

function ForgottenPassword() {
    const { userData, isLoggedIn } = useContext(UserContext)
    if (isLoggedIn) {
        return (
            <Navigate
                to={routeMap.profile.replace(
                    ":user",
                    encodeURIComponent(userData.username),
                )}
            />
        )
    }
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full sm:max-w-md">
                <div className="w-full px-4">
                    <ForgottenPasswordFrame />
                </div>
            </div>
        </div>
    )
}

export default ForgottenPassword
