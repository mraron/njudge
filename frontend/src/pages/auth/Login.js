import { useContext, useState } from "react"
import { Link, Navigate, useNavigate } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import RoundedFrame, { SVGTitleComponent } from "../../components/container/RoundedFrame"
import TextBox from "../../components/input/TextBox"
import Button from "../../components/basic/Button"
import UserContext from "../../contexts/user/UserContext"
import { routeMap } from "../../config/RouteConfig"
import { login } from "../../util/auth"
import NarrowPage from "../wrappers/NarrowPage";

function LoginFrame() {
    const { t } = useTranslation()
    const navigate = useNavigate()
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-user-check" className="w-5 h-5 mr-3" />}
            title={t("login.login")}
        />
    )
    const handleLogin = (event) => {
        event.preventDefault()
        login(username, password).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_login", "success")
                navigate(routeMap.home)
            } else {
                window.flash("flash.unsuccessful_login", "failure")
            }
        })
    }
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 py-8">
                <form method="POST">
                    <div className="mb-4">
                        <TextBox id="username" label={t("login.username")} initText={username} onChange={setUsername} />
                    </div>
                    <div className="mb-2">
                        <TextBox
                            id="password"
                            label={t("login.password")}
                            initText={password}
                            type="password"
                            onChange={setPassword}
                        />
                    </div>
                    <div className="mb-6">
                        <Link to={routeMap.forgotten_password} className="link text-sm">
                            {t("login.forgotten_password")}
                        </Link>
                    </div>
                    <div className="mb-2 flex justify-center space-x-2">
                        <Button color="indigo" cls="py-2.5" onClick={handleLogin} fullWidth={true}>
                            {t("login.login")}
                        </Button>
                        <Button color="gray" cls="py-2.5" fullWidth={true}>
                            Google
                        </Button>
                    </div>
                </form>
            </div>
        </RoundedFrame>
    )
}

function Login() {
    const { userData, isLoggedIn } = useContext(UserContext)
    if (isLoggedIn) {
        return <Navigate to={routeMap.profile.replace(":user", encodeURIComponent(userData.username))} />
    }
    return (
        <NarrowPage>
            <LoginFrame />
        </NarrowPage>
    )
}

export default Login
