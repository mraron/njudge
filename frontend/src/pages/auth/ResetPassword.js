import { useContext, useState } from "react";
import { Navigate, useParams } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedFrame from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import UserContext from "../../contexts/user/UserContext";
import { routeMap } from "../../config/RouteConfig";
import { reset_password } from "../../util/auth";
import Button from "../../components/util/Button";

function ResetPasswordFrame() {
    const { t } = useTranslation();
    const { user, token } = useParams();
    const [password, setPassword] = useState("");
    const [passwordConfirm, setPasswordConfirm] = useState("");

    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon="fa-unlock" className="w-4 h-4 mr-3" />}
            title={t("reset_password.change_password")}
        />
    );
    const handleResetPassword = (event) => {
        event.preventDefault();
        reset_password(user, token, password, passwordConfirm).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_password_change", "success");
            } else {
                window.flash(resp.message, "failure");
            }
        });
    };
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <form method="POST">
                <div className="px-10 py-8">
                    <div className="mb-4">
                        <TextBox
                            id="password"
                            label={t("reset_password.password")}
                            type="password"
                            initText={password}
                            onChange={setPassword}
                        />
                    </div>
                    <div className="mb-6">
                        <TextBox
                            id="passwordConfirm"
                            label={t("reset_password.confirm_password")}
                            type="password"
                            initText={passwordConfirm}
                            onChange={setPasswordConfirm}
                        />
                    </div>
                    <div className="flex justify-center">
                        <Button
                            type="submit"
                            theme="indigo"
                            onClick={handleResetPassword}
                            minWidth="12rem">
                            {t("reset_password.change_password")}
                        </Button>
                    </div>
                </div>
            </form>
        </RoundedFrame>
    );
}

function ResetPassword() {
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
                    <ResetPasswordFrame />
                </div>
            </div>
        </div>
    );
}

export default ResetPassword;
