import { Link } from "react-router-dom";
import { SVGNotFound } from "../../components/svg/SVGs";
import RoundedFrame from "../../components/container/RoundedFrame";

function NotFoundFrame() {
    return (
        <RoundedFrame title="Az oldal nem elérhető">
            <div className="px-10 py-8 flex flex-col relative justify-between">
                <p className="z-10">
                    A keresett oldal nem található. Győződj meg róla, hogy a
                    megadott link helyes.
                </p>
                <div className="flex justify-center absolute inset-0">
                    <SVGNotFound />
                </div>
                <div className="flex justify-center mt-8">
                    <Link
                        className="z-10 w-60 btn-indigo padding-btn-default text-center"
                        to="/">
                        Vissza a főoldalra
                    </Link>
                </div>
            </div>
        </RoundedFrame>
    );
}

function NotFound() {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-md px-4">
                <NotFoundFrame />
            </div>
        </div>
    );
}

export default NotFound;
