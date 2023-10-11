import { useState } from "react";
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedFrame, { SVGTitleComponent } from "../../components/container/RoundedFrame";
import TextBox from "../../components/input/TextBox";
import Button from "../../components/basic/Button";
import { change_password } from "../../util/auth";
import NarrowPage from "../wrappers/NarrowPage";

function ForgottenPasswordFrame() {
    const { t } = useTranslation()
    const [email, setEmail] = useState("")
    const titleComponent = (
        <SVGTitleComponent
            icon={<FontAwesomeIcon icon="fa-lock" className="w-5 h-5 mr-3" />}
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
    return (
        <NarrowPage>
            <ForgottenPasswordFrame />
        </NarrowPage>
    )
}

export default ForgottenPassword
