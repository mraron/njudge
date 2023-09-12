import RoundedFrame from '../../components/RoundedFrame';
import TextBox, { SVGInputComponent } from '../../components/TextBox';
import { SVGAvatar, SVGConfirm, SVGLock, SVGMail } from '../../svg/SVGs';
import SVGTitleComponent from '../../svg/SVGTitleComponent';

function RegisterFrame() {
    const titleComponent = <SVGTitleComponent svg={<SVGAvatar cls="w-[1.1rem] h-[1.1rem] mr-2" />} title="Regisztráció" />
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 pt-8 pb-6 border-b border-default">
                <div className="mb-4 relative">
                    <TextBox id="username" label="Felhasználónév" inputComponent={SVGInputComponent(<SVGAvatar cls="w-4 h-4" />)} />
                </div>
                <TextBox id="email" label="E-mail cím" inputComponent={SVGInputComponent(<SVGMail cls="w-4 h-4" />)} />
            </div>
            <div className="px-10 pt-4 pb-8">
                <div className="mb-4">
                    <TextBox id="password" label="Jelszó" inputComponent={SVGInputComponent(<SVGLock cls="w-4 h-4" />)} type="password" />
                </div>
                <div className="mb-6">
                    <TextBox id="passwordConfirm" label="Jelszó megerősítése" inputComponent={SVGInputComponent(<SVGConfirm cls="w-5 h-5" />)} type="password" />
                </div>
                <div className="flex justify-center">
                    <button className="btn-indigo w-40">Regisztráció</button>
                </div>
            </div>
        </RoundedFrame>
    )
}

function Register() {
	return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full sm:max-w-md">
                    <div className="w-full px-4">
                        <RegisterFrame />
                    </div>
                </div>
            </div>
        </div>
	);
}

export default Register;