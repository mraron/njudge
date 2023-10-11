import { useState } from "react";
import { useParams } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedFrame, { SVGTitleComponent } from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import Button from "../../components/basic/Button";
import { reset_password } from "../../util/auth";
import NarrowPage from "../wrappers/NarrowPage";

function ResetPasswordFrame() {
    const { t } = useTranslation()
    const { user, token } = useParams()
    const [password, setPassword] = useState("")
    const [passwordConfirm, setPasswordConfirm] = useState("")

    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-lock" className="w-4 h-4 mr-3" />}
            title={t("reset_password.change_password")}
        />
    )
    const handleResetPassword = (event) => {
        event.preventDefault()
        reset_password(user, token, password, passwordConfirm).then((resp) => {
            if (resp.success) {
                window.flash("flash.successful_password_change", "success")
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
                    <div className="mb-2 flex justify-center">
                        <Button
                            type="submit"
                            color="indigo"
                            cls="py-2.5"
                            onClick={handleResetPassword}
                            minWidth="12rem">
                            {t("reset_password.change_password")}
                        </Button>
                    </div>
                </form>
            </div>
        </RoundedFrame>
    )
}

function ResetPassword() {
    return (
        <NarrowPage>
            <ResetPasswordFrame />
        </NarrowPage>
    )
}

export default ResetPassword
