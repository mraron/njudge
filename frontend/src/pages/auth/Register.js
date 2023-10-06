import { useContext, useState } from "react"
import { useTranslation } from "react-i18next"
import { Link, Navigate, useNavigate } from "react-router-dom"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import RoundedFrame, {
    SVGTitleComponent,
} from "../../components/container/RoundedFrame"
import TextBox from "../../components/input/TextBox"
import Button from "../../components/basic/Button"
import UserContext from "../../contexts/user/UserContext"
import { routeMap } from "../../config/RouteConfig"
import { register } from "../../util/auth"

function RegisterFrame() {
    const { t } = useTranslation()
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [passwordConfirm, setPasswordConfirm] = useState("")
    const [email, setEmail] = useState("")
    const navigate = useNavigate()
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon icon="fa-user-plus" className="w-5 h-5 mr-3" />
            }
            title={t("register.register")}
        />
    )
    const handleRegister = (event) => {
        event.preventDefault()
        register(username, email, password, passwordConfirm).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_registration", "success")
                navigate(routeMap.home)
            } else {
                window.flash(resp.message, "failure")
            }
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 py-8">
                <form method="POST">
                    <div className="mb-4">
                        <TextBox
                            id="username"
                            label={t("register.username")}
                            initText={username}
                            onChange={setUsername}
                        />
                    </div>
                    <div className="mb-4">
                        <TextBox
                            id="email"
                            label={t("register.email")}
                            initText={email}
                            onChange={setEmail}
                        />
                    </div>
                    <div className="mb-4">
                        <TextBox
                            id="password"
                            label={t("register.password")}
                            type="password"
                            initText={password}
                            onChange={setPassword}
                        />
                    </div>
                    <div className="mb-2">
                        <TextBox
                            id="passwordConfirm"
                            label={t("register.confirm_password")}
                            type="password"
                            initText={passwordConfirm}
                            onChange={setPasswordConfirm}
                        />
                    </div>
                    <div className="mb-6">
                        <Link to={routeMap.login} className="link text-sm">
                            {t("register.already_registered")}
                        </Link>
                    </div>
                    <div className="mb-2 flex justify-center">
                        <Button
                            type="submit"
                            color="indigo"
                            cls="py-2.5"
                            onClick={handleRegister}
                            minWidth="12rem">
                            {t("register.register")}
                        </Button>
                    </div>
                </form>
            </div>
        </RoundedFrame>
    )
}

function Register() {
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
                    <RegisterFrame />
                </div>
            </div>
        </div>
    )
}

export default Register
