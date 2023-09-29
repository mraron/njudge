import { useContext, useState } from "react";
import { Navigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedFrame from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import { routeMap } from "../../config/RouteConfig";
import UserContext from "../../contexts/user/UserContext";
import { change_password } from "../../util/auth";

function ForgottenPasswordFrame() {
    const { t } = useTranslation();
    const [email, setEmail] = useState("");
    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon="fa-unlock" className="w-5 h-5 mr-3" />}
            title={t("forgotten_password.change_password")}
        />
    );
    const handleChangePassword = (event) => {
        event.preventDefault();
        change_password(email).then((ok) => {
            if (ok) {
                window.flash("flash.successful_email_pw_change", "success");
            } else {
                window.flash("flash.unsuccessful_email_pw_change", "failure");
            }
        });
    };
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <form method="POST">
                <div className="px-10 py-8">
                    <div className="mb-6">
                        <TextBox
                            id="email"
                            label={t("forgotten_password.email")}
                            initText={email}
                            onChange={setEmail}
                        />
                    </div>
                    <div className="flex justify-center">
                        <button
                            type="submit"
                            className="btn-indigo padding-btn-default min-w-[12rem]"
                            onClick={handleChangePassword}>
                            {t("forgotten_password.change_password")}
                        </button>
                    </div>
                </div>
            </form>
        </RoundedFrame>
    );
}

function ForgottenPassword() {
    const { userData, isLoggedIn } = useContext(UserContext);
    if (isLoggedIn) {
        return (
            <Navigate
                to={routeMap.profile.replace(
                    ":user",
                    encodeURIComponent(userData.username),
                )}
            />
        );
    }
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full sm:max-w-md">
                <div className="w-full px-4">
                    <ForgottenPasswordFrame />
                </div>
            </div>
        </div>
    );
}

export default ForgottenPassword;
