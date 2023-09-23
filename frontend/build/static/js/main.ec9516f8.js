/*! For license information please see main.ec9516f8.js.LICENSE.txt */
!(function () {
    "use strict";
    var e = {
            463: function (e, t, n) {
                var r = n(791),
                    a = n(296);
                function l(e) {
                    for (
                        var t =
                                "https://reactjs.org/docs/error-decoder.html?invariant=" +
                                e,
                            n = 1;
                        n < arguments.length;
                        n++
                    )
                        t += "&args[]=" + encodeURIComponent(arguments[n]);
                    return (
                        "Minified React error #" +
                        e +
                        "; visit " +
                        t +
                        " for the full message or use the non-minified dev environment for full errors and additional helpful warnings."
                    );
                }
                var i = new Set(),
                    o = {};
                function s(e, t) {
                    u(e, t), u(e + "Capture", t);
                }
                function u(e, t) {
                    for (o[e] = t, e = 0; e < t.length; e++) i.add(t[e]);
                }
                var c = !(
                        "undefined" === typeof window ||
                        "undefined" === typeof window.document ||
                        "undefined" === typeof window.document.createElement
                    ),
                    d = Object.prototype.hasOwnProperty,
                    f =
                        /^[:A-Z_a-z\u00C0-\u00D6\u00D8-\u00F6\u00F8-\u02FF\u0370-\u037D\u037F-\u1FFF\u200C-\u200D\u2070-\u218F\u2C00-\u2FEF\u3001-\uD7FF\uF900-\uFDCF\uFDF0-\uFFFD][:A-Z_a-z\u00C0-\u00D6\u00D8-\u00F6\u00F8-\u02FF\u0370-\u037D\u037F-\u1FFF\u200C-\u200D\u2070-\u218F\u2C00-\u2FEF\u3001-\uD7FF\uF900-\uFDCF\uFDF0-\uFFFD\-.0-9\u00B7\u0300-\u036F\u203F-\u2040]*$/,
                    p = {},
                    m = {};
                function h(e, t, n, r, a, l, i) {
                    (this.acceptsBooleans = 2 === t || 3 === t || 4 === t),
                        (this.attributeName = r),
                        (this.attributeNamespace = a),
                        (this.mustUseProperty = n),
                        (this.propertyName = e),
                        (this.type = t),
                        (this.sanitizeURL = l),
                        (this.removeEmptyString = i);
                }
                var v = {};
                "children dangerouslySetInnerHTML defaultValue defaultChecked innerHTML suppressContentEditableWarning suppressHydrationWarning style"
                    .split(" ")
                    .forEach(function (e) {
                        v[e] = new h(e, 0, !1, e, null, !1, !1);
                    }),
                    [
                        ["acceptCharset", "accept-charset"],
                        ["className", "class"],
                        ["htmlFor", "for"],
                        ["httpEquiv", "http-equiv"],
                    ].forEach(function (e) {
                        var t = e[0];
                        v[t] = new h(t, 1, !1, e[1], null, !1, !1);
                    }),
                    [
                        "contentEditable",
                        "draggable",
                        "spellCheck",
                        "value",
                    ].forEach(function (e) {
                        v[e] = new h(e, 2, !1, e.toLowerCase(), null, !1, !1);
                    }),
                    [
                        "autoReverse",
                        "externalResourcesRequired",
                        "focusable",
                        "preserveAlpha",
                    ].forEach(function (e) {
                        v[e] = new h(e, 2, !1, e, null, !1, !1);
                    }),
                    "allowFullScreen async autoFocus autoPlay controls default defer disabled disablePictureInPicture disableRemotePlayback formNoValidate hidden loop noModule noValidate open playsInline readOnly required reversed scoped seamless itemScope"
                        .split(" ")
                        .forEach(function (e) {
                            v[e] = new h(
                                e,
                                3,
                                !1,
                                e.toLowerCase(),
                                null,
                                !1,
                                !1,
                            );
                        }),
                    ["checked", "multiple", "muted", "selected"].forEach(
                        function (e) {
                            v[e] = new h(e, 3, !0, e, null, !1, !1);
                        },
                    ),
                    ["capture", "download"].forEach(function (e) {
                        v[e] = new h(e, 4, !1, e, null, !1, !1);
                    }),
                    ["cols", "rows", "size", "span"].forEach(function (e) {
                        v[e] = new h(e, 6, !1, e, null, !1, !1);
                    }),
                    ["rowSpan", "start"].forEach(function (e) {
                        v[e] = new h(e, 5, !1, e.toLowerCase(), null, !1, !1);
                    });
                var g = /[\-:]([a-z])/g;
                function y(e) {
                    return e[1].toUpperCase();
                }
                function b(e, t, n, r) {
                    var a = v.hasOwnProperty(t) ? v[t] : null;
                    (null !== a
                        ? 0 !== a.type
                        : r ||
                          !(2 < t.length) ||
                          ("o" !== t[0] && "O" !== t[0]) ||
                          ("n" !== t[1] && "N" !== t[1])) &&
                        ((function (e, t, n, r) {
                            if (
                                null === t ||
                                "undefined" === typeof t ||
                                (function (e, t, n, r) {
                                    if (null !== n && 0 === n.type) return !1;
                                    switch (typeof t) {
                                        case "function":
                                        case "symbol":
                                            return !0;
                                        case "boolean":
                                            return (
                                                !r &&
                                                (null !== n
                                                    ? !n.acceptsBooleans
                                                    : "data-" !==
                                                          (e = e
                                                              .toLowerCase()
                                                              .slice(0, 5)) &&
                                                      "aria-" !== e)
                                            );
                                        default:
                                            return !1;
                                    }
                                })(e, t, n, r)
                            )
                                return !0;
                            if (r) return !1;
                            if (null !== n)
                                switch (n.type) {
                                    case 3:
                                        return !t;
                                    case 4:
                                        return !1 === t;
                                    case 5:
                                        return isNaN(t);
                                    case 6:
                                        return isNaN(t) || 1 > t;
                                }
                            return !1;
                        })(t, n, a, r) && (n = null),
                        r || null === a
                            ? (function (e) {
                                  return (
                                      !!d.call(m, e) ||
                                      (!d.call(p, e) &&
                                          (f.test(e)
                                              ? (m[e] = !0)
                                              : ((p[e] = !0), !1)))
                                  );
                              })(t) &&
                              (null === n
                                  ? e.removeAttribute(t)
                                  : e.setAttribute(t, "" + n))
                            : a.mustUseProperty
                            ? (e[a.propertyName] =
                                  null === n ? 3 !== a.type && "" : n)
                            : ((t = a.attributeName),
                              (r = a.attributeNamespace),
                              null === n
                                  ? e.removeAttribute(t)
                                  : ((n =
                                        3 === (a = a.type) ||
                                        (4 === a && !0 === n)
                                            ? ""
                                            : "" + n),
                                    r
                                        ? e.setAttributeNS(r, t, n)
                                        : e.setAttribute(t, n))));
                }
                "accent-height alignment-baseline arabic-form baseline-shift cap-height clip-path clip-rule color-interpolation color-interpolation-filters color-profile color-rendering dominant-baseline enable-background fill-opacity fill-rule flood-color flood-opacity font-family font-size font-size-adjust font-stretch font-style font-variant font-weight glyph-name glyph-orientation-horizontal glyph-orientation-vertical horiz-adv-x horiz-origin-x image-rendering letter-spacing lighting-color marker-end marker-mid marker-start overline-position overline-thickness paint-order panose-1 pointer-events rendering-intent shape-rendering stop-color stop-opacity strikethrough-position strikethrough-thickness stroke-dasharray stroke-dashoffset stroke-linecap stroke-linejoin stroke-miterlimit stroke-opacity stroke-width text-anchor text-decoration text-rendering underline-position underline-thickness unicode-bidi unicode-range units-per-em v-alphabetic v-hanging v-ideographic v-mathematical vector-effect vert-adv-y vert-origin-x vert-origin-y word-spacing writing-mode xmlns:xlink x-height"
                    .split(" ")
                    .forEach(function (e) {
                        var t = e.replace(g, y);
                        v[t] = new h(t, 1, !1, e, null, !1, !1);
                    }),
                    "xlink:actuate xlink:arcrole xlink:role xlink:show xlink:title xlink:type"
                        .split(" ")
                        .forEach(function (e) {
                            var t = e.replace(g, y);
                            v[t] = new h(
                                t,
                                1,
                                !1,
                                e,
                                "http://www.w3.org/1999/xlink",
                                !1,
                                !1,
                            );
                        }),
                    ["xml:base", "xml:lang", "xml:space"].forEach(function (e) {
                        var t = e.replace(g, y);
                        v[t] = new h(
                            t,
                            1,
                            !1,
                            e,
                            "http://www.w3.org/XML/1998/namespace",
                            !1,
                            !1,
                        );
                    }),
                    ["tabIndex", "crossOrigin"].forEach(function (e) {
                        v[e] = new h(e, 1, !1, e.toLowerCase(), null, !1, !1);
                    }),
                    (v.xlinkHref = new h(
                        "xlinkHref",
                        1,
                        !1,
                        "xlink:href",
                        "http://www.w3.org/1999/xlink",
                        !0,
                        !1,
                    )),
                    ["src", "href", "action", "formAction"].forEach(
                        function (e) {
                            v[e] = new h(
                                e,
                                1,
                                !1,
                                e.toLowerCase(),
                                null,
                                !0,
                                !0,
                            );
                        },
                    );
                var x = r.__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED,
                    w = Symbol.for("react.element"),
                    j = Symbol.for("react.portal"),
                    k = Symbol.for("react.fragment"),
                    S = Symbol.for("react.strict_mode"),
                    N = Symbol.for("react.profiler"),
                    C = Symbol.for("react.provider"),
                    E = Symbol.for("react.context"),
                    L = Symbol.for("react.forward_ref"),
                    _ = Symbol.for("react.suspense"),
                    P = Symbol.for("react.suspense_list"),
                    O = Symbol.for("react.memo"),
                    z = Symbol.for("react.lazy");
                Symbol.for("react.scope"), Symbol.for("react.debug_trace_mode");
                var M = Symbol.for("react.offscreen");
                Symbol.for("react.legacy_hidden"),
                    Symbol.for("react.cache"),
                    Symbol.for("react.tracing_marker");
                var R = Symbol.iterator;
                function T(e) {
                    return null === e || "object" !== typeof e
                        ? null
                        : "function" ===
                          typeof (e = (R && e[R]) || e["@@iterator"])
                        ? e
                        : null;
                }
                var F,
                    I = Object.assign;
                function D(e) {
                    if (void 0 === F)
                        try {
                            throw Error();
                        } catch (n) {
                            var t = n.stack.trim().match(/\n( *(at )?)/);
                            F = (t && t[1]) || "";
                        }
                    return "\n" + F + e;
                }
                var U = !1;
                function B(e, t) {
                    if (!e || U) return "";
                    U = !0;
                    var n = Error.prepareStackTrace;
                    Error.prepareStackTrace = void 0;
                    try {
                        if (t)
                            if (
                                ((t = function () {
                                    throw Error();
                                }),
                                Object.defineProperty(t.prototype, "props", {
                                    set: function () {
                                        throw Error();
                                    },
                                }),
                                "object" === typeof Reflect &&
                                    Reflect.construct)
                            ) {
                                try {
                                    Reflect.construct(t, []);
                                } catch (u) {
                                    var r = u;
                                }
                                Reflect.construct(e, [], t);
                            } else {
                                try {
                                    t.call();
                                } catch (u) {
                                    r = u;
                                }
                                e.call(t.prototype);
                            }
                        else {
                            try {
                                throw Error();
                            } catch (u) {
                                r = u;
                            }
                            e();
                        }
                    } catch (u) {
                        if (u && r && "string" === typeof u.stack) {
                            for (
                                var a = u.stack.split("\n"),
                                    l = r.stack.split("\n"),
                                    i = a.length - 1,
                                    o = l.length - 1;
                                1 <= i && 0 <= o && a[i] !== l[o];

                            )
                                o--;
                            for (; 1 <= i && 0 <= o; i--, o--)
                                if (a[i] !== l[o]) {
                                    if (1 !== i || 1 !== o)
                                        do {
                                            if (
                                                (i--, 0 > --o || a[i] !== l[o])
                                            ) {
                                                var s =
                                                    "\n" +
                                                    a[i].replace(
                                                        " at new ",
                                                        " at ",
                                                    );
                                                return (
                                                    e.displayName &&
                                                        s.includes(
                                                            "<anonymous>",
                                                        ) &&
                                                        (s = s.replace(
                                                            "<anonymous>",
                                                            e.displayName,
                                                        )),
                                                    s
                                                );
                                            }
                                        } while (1 <= i && 0 <= o);
                                    break;
                                }
                        }
                    } finally {
                        (U = !1), (Error.prepareStackTrace = n);
                    }
                    return (e = e ? e.displayName || e.name : "") ? D(e) : "";
                }
                function A(e) {
                    switch (e.tag) {
                        case 5:
                            return D(e.type);
                        case 16:
                            return D("Lazy");
                        case 13:
                            return D("Suspense");
                        case 19:
                            return D("SuspenseList");
                        case 0:
                        case 2:
                        case 15:
                            return (e = B(e.type, !1));
                        case 11:
                            return (e = B(e.type.render, !1));
                        case 1:
                            return (e = B(e.type, !0));
                        default:
                            return "";
                    }
                }
                function V(e) {
                    if (null == e) return null;
                    if ("function" === typeof e)
                        return e.displayName || e.name || null;
                    if ("string" === typeof e) return e;
                    switch (e) {
                        case k:
                            return "Fragment";
                        case j:
                            return "Portal";
                        case N:
                            return "Profiler";
                        case S:
                            return "StrictMode";
                        case _:
                            return "Suspense";
                        case P:
                            return "SuspenseList";
                    }
                    if ("object" === typeof e)
                        switch (e.$$typeof) {
                            case E:
                                return (
                                    (e.displayName || "Context") + ".Consumer"
                                );
                            case C:
                                return (
                                    (e._context.displayName || "Context") +
                                    ".Provider"
                                );
                            case L:
                                var t = e.render;
                                return (
                                    (e = e.displayName) ||
                                        (e =
                                            "" !==
                                            (e = t.displayName || t.name || "")
                                                ? "ForwardRef(" + e + ")"
                                                : "ForwardRef"),
                                    e
                                );
                            case O:
                                return null !== (t = e.displayName || null)
                                    ? t
                                    : V(e.type) || "Memo";
                            case z:
                                (t = e._payload), (e = e._init);
                                try {
                                    return V(e(t));
                                } catch (n) {}
                        }
                    return null;
                }
                function $(e) {
                    var t = e.type;
                    switch (e.tag) {
                        case 24:
                            return "Cache";
                        case 9:
                            return (t.displayName || "Context") + ".Consumer";
                        case 10:
                            return (
                                (t._context.displayName || "Context") +
                                ".Provider"
                            );
                        case 18:
                            return "DehydratedFragment";
                        case 11:
                            return (
                                (e =
                                    (e = t.render).displayName || e.name || ""),
                                t.displayName ||
                                    ("" !== e
                                        ? "ForwardRef(" + e + ")"
                                        : "ForwardRef")
                            );
                        case 7:
                            return "Fragment";
                        case 5:
                            return t;
                        case 4:
                            return "Portal";
                        case 3:
                            return "Root";
                        case 6:
                            return "Text";
                        case 16:
                            return V(t);
                        case 8:
                            return t === S ? "StrictMode" : "Mode";
                        case 22:
                            return "Offscreen";
                        case 12:
                            return "Profiler";
                        case 21:
                            return "Scope";
                        case 13:
                            return "Suspense";
                        case 19:
                            return "SuspenseList";
                        case 25:
                            return "TracingMarker";
                        case 1:
                        case 0:
                        case 17:
                        case 2:
                        case 14:
                        case 15:
                            if ("function" === typeof t)
                                return t.displayName || t.name || null;
                            if ("string" === typeof t) return t;
                    }
                    return null;
                }
                function H(e) {
                    switch (typeof e) {
                        case "boolean":
                        case "number":
                        case "string":
                        case "undefined":
                        case "object":
                            return e;
                        default:
                            return "";
                    }
                }
                function K(e) {
                    var t = e.type;
                    return (
                        (e = e.nodeName) &&
                        "input" === e.toLowerCase() &&
                        ("checkbox" === t || "radio" === t)
                    );
                }
                function W(e) {
                    e._valueTracker ||
                        (e._valueTracker = (function (e) {
                            var t = K(e) ? "checked" : "value",
                                n = Object.getOwnPropertyDescriptor(
                                    e.constructor.prototype,
                                    t,
                                ),
                                r = "" + e[t];
                            if (
                                !e.hasOwnProperty(t) &&
                                "undefined" !== typeof n &&
                                "function" === typeof n.get &&
                                "function" === typeof n.set
                            ) {
                                var a = n.get,
                                    l = n.set;
                                return (
                                    Object.defineProperty(e, t, {
                                        configurable: !0,
                                        get: function () {
                                            return a.call(this);
                                        },
                                        set: function (e) {
                                            (r = "" + e), l.call(this, e);
                                        },
                                    }),
                                    Object.defineProperty(e, t, {
                                        enumerable: n.enumerable,
                                    }),
                                    {
                                        getValue: function () {
                                            return r;
                                        },
                                        setValue: function (e) {
                                            r = "" + e;
                                        },
                                        stopTracking: function () {
                                            (e._valueTracker = null),
                                                delete e[t];
                                        },
                                    }
                                );
                            }
                        })(e));
                }
                function Q(e) {
                    if (!e) return !1;
                    var t = e._valueTracker;
                    if (!t) return !0;
                    var n = t.getValue(),
                        r = "";
                    return (
                        e &&
                            (r = K(e)
                                ? e.checked
                                    ? "true"
                                    : "false"
                                : e.value),
                        (e = r) !== n && (t.setValue(e), !0)
                    );
                }
                function q(e) {
                    if (
                        "undefined" ===
                        typeof (e =
                            e ||
                            ("undefined" !== typeof document
                                ? document
                                : void 0))
                    )
                        return null;
                    try {
                        return e.activeElement || e.body;
                    } catch (t) {
                        return e.body;
                    }
                }
                function G(e, t) {
                    var n = t.checked;
                    return I({}, t, {
                        defaultChecked: void 0,
                        defaultValue: void 0,
                        value: void 0,
                        checked: null != n ? n : e._wrapperState.initialChecked,
                    });
                }
                function Y(e, t) {
                    var n = null == t.defaultValue ? "" : t.defaultValue,
                        r = null != t.checked ? t.checked : t.defaultChecked;
                    (n = H(null != t.value ? t.value : n)),
                        (e._wrapperState = {
                            initialChecked: r,
                            initialValue: n,
                            controlled:
                                "checkbox" === t.type || "radio" === t.type
                                    ? null != t.checked
                                    : null != t.value,
                        });
                }
                function X(e, t) {
                    null != (t = t.checked) && b(e, "checked", t, !1);
                }
                function Z(e, t) {
                    X(e, t);
                    var n = H(t.value),
                        r = t.type;
                    if (null != n)
                        "number" === r
                            ? ((0 === n && "" === e.value) || e.value != n) &&
                              (e.value = "" + n)
                            : e.value !== "" + n && (e.value = "" + n);
                    else if ("submit" === r || "reset" === r)
                        return void e.removeAttribute("value");
                    t.hasOwnProperty("value")
                        ? ee(e, t.type, n)
                        : t.hasOwnProperty("defaultValue") &&
                          ee(e, t.type, H(t.defaultValue)),
                        null == t.checked &&
                            null != t.defaultChecked &&
                            (e.defaultChecked = !!t.defaultChecked);
                }
                function J(e, t, n) {
                    if (
                        t.hasOwnProperty("value") ||
                        t.hasOwnProperty("defaultValue")
                    ) {
                        var r = t.type;
                        if (
                            !(
                                ("submit" !== r && "reset" !== r) ||
                                (void 0 !== t.value && null !== t.value)
                            )
                        )
                            return;
                        (t = "" + e._wrapperState.initialValue),
                            n || t === e.value || (e.value = t),
                            (e.defaultValue = t);
                    }
                    "" !== (n = e.name) && (e.name = ""),
                        (e.defaultChecked = !!e._wrapperState.initialChecked),
                        "" !== n && (e.name = n);
                }
                function ee(e, t, n) {
                    ("number" === t && q(e.ownerDocument) === e) ||
                        (null == n
                            ? (e.defaultValue =
                                  "" + e._wrapperState.initialValue)
                            : e.defaultValue !== "" + n &&
                              (e.defaultValue = "" + n));
                }
                var te = Array.isArray;
                function ne(e, t, n, r) {
                    if (((e = e.options), t)) {
                        t = {};
                        for (var a = 0; a < n.length; a++) t["$" + n[a]] = !0;
                        for (n = 0; n < e.length; n++)
                            (a = t.hasOwnProperty("$" + e[n].value)),
                                e[n].selected !== a && (e[n].selected = a),
                                a && r && (e[n].defaultSelected = !0);
                    } else {
                        for (
                            n = "" + H(n), t = null, a = 0;
                            a < e.length;
                            a++
                        ) {
                            if (e[a].value === n)
                                return (
                                    (e[a].selected = !0),
                                    void (r && (e[a].defaultSelected = !0))
                                );
                            null !== t || e[a].disabled || (t = e[a]);
                        }
                        null !== t && (t.selected = !0);
                    }
                }
                function re(e, t) {
                    if (null != t.dangerouslySetInnerHTML) throw Error(l(91));
                    return I({}, t, {
                        value: void 0,
                        defaultValue: void 0,
                        children: "" + e._wrapperState.initialValue,
                    });
                }
                function ae(e, t) {
                    var n = t.value;
                    if (null == n) {
                        if (
                            ((n = t.children), (t = t.defaultValue), null != n)
                        ) {
                            if (null != t) throw Error(l(92));
                            if (te(n)) {
                                if (1 < n.length) throw Error(l(93));
                                n = n[0];
                            }
                            t = n;
                        }
                        null == t && (t = ""), (n = t);
                    }
                    e._wrapperState = { initialValue: H(n) };
                }
                function le(e, t) {
                    var n = H(t.value),
                        r = H(t.defaultValue);
                    null != n &&
                        ((n = "" + n) !== e.value && (e.value = n),
                        null == t.defaultValue &&
                            e.defaultValue !== n &&
                            (e.defaultValue = n)),
                        null != r && (e.defaultValue = "" + r);
                }
                function ie(e) {
                    var t = e.textContent;
                    t === e._wrapperState.initialValue &&
                        "" !== t &&
                        null !== t &&
                        (e.value = t);
                }
                function oe(e) {
                    switch (e) {
                        case "svg":
                            return "http://www.w3.org/2000/svg";
                        case "math":
                            return "http://www.w3.org/1998/Math/MathML";
                        default:
                            return "http://www.w3.org/1999/xhtml";
                    }
                }
                function se(e, t) {
                    return null == e || "http://www.w3.org/1999/xhtml" === e
                        ? oe(t)
                        : "http://www.w3.org/2000/svg" === e &&
                          "foreignObject" === t
                        ? "http://www.w3.org/1999/xhtml"
                        : e;
                }
                var ue,
                    ce,
                    de =
                        ((ce = function (e, t) {
                            if (
                                "http://www.w3.org/2000/svg" !==
                                    e.namespaceURI ||
                                "innerHTML" in e
                            )
                                e.innerHTML = t;
                            else {
                                for (
                                    (ue =
                                        ue ||
                                        document.createElement(
                                            "div",
                                        )).innerHTML =
                                        "<svg>" +
                                        t.valueOf().toString() +
                                        "</svg>",
                                        t = ue.firstChild;
                                    e.firstChild;

                                )
                                    e.removeChild(e.firstChild);
                                for (; t.firstChild; )
                                    e.appendChild(t.firstChild);
                            }
                        }),
                        "undefined" !== typeof MSApp &&
                        MSApp.execUnsafeLocalFunction
                            ? function (e, t, n, r) {
                                  MSApp.execUnsafeLocalFunction(function () {
                                      return ce(e, t);
                                  });
                              }
                            : ce);
                function fe(e, t) {
                    if (t) {
                        var n = e.firstChild;
                        if (n && n === e.lastChild && 3 === n.nodeType)
                            return void (n.nodeValue = t);
                    }
                    e.textContent = t;
                }
                var pe = {
                        animationIterationCount: !0,
                        aspectRatio: !0,
                        borderImageOutset: !0,
                        borderImageSlice: !0,
                        borderImageWidth: !0,
                        boxFlex: !0,
                        boxFlexGroup: !0,
                        boxOrdinalGroup: !0,
                        columnCount: !0,
                        columns: !0,
                        flex: !0,
                        flexGrow: !0,
                        flexPositive: !0,
                        flexShrink: !0,
                        flexNegative: !0,
                        flexOrder: !0,
                        gridArea: !0,
                        gridRow: !0,
                        gridRowEnd: !0,
                        gridRowSpan: !0,
                        gridRowStart: !0,
                        gridColumn: !0,
                        gridColumnEnd: !0,
                        gridColumnSpan: !0,
                        gridColumnStart: !0,
                        fontWeight: !0,
                        lineClamp: !0,
                        lineHeight: !0,
                        opacity: !0,
                        order: !0,
                        orphans: !0,
                        tabSize: !0,
                        widows: !0,
                        zIndex: !0,
                        zoom: !0,
                        fillOpacity: !0,
                        floodOpacity: !0,
                        stopOpacity: !0,
                        strokeDasharray: !0,
                        strokeDashoffset: !0,
                        strokeMiterlimit: !0,
                        strokeOpacity: !0,
                        strokeWidth: !0,
                    },
                    me = ["Webkit", "ms", "Moz", "O"];
                function he(e, t, n) {
                    return null == t || "boolean" === typeof t || "" === t
                        ? ""
                        : n ||
                          "number" !== typeof t ||
                          0 === t ||
                          (pe.hasOwnProperty(e) && pe[e])
                        ? ("" + t).trim()
                        : t + "px";
                }
                function ve(e, t) {
                    for (var n in ((e = e.style), t))
                        if (t.hasOwnProperty(n)) {
                            var r = 0 === n.indexOf("--"),
                                a = he(n, t[n], r);
                            "float" === n && (n = "cssFloat"),
                                r ? e.setProperty(n, a) : (e[n] = a);
                        }
                }
                Object.keys(pe).forEach(function (e) {
                    me.forEach(function (t) {
                        (t = t + e.charAt(0).toUpperCase() + e.substring(1)),
                            (pe[t] = pe[e]);
                    });
                });
                var ge = I(
                    { menuitem: !0 },
                    {
                        area: !0,
                        base: !0,
                        br: !0,
                        col: !0,
                        embed: !0,
                        hr: !0,
                        img: !0,
                        input: !0,
                        keygen: !0,
                        link: !0,
                        meta: !0,
                        param: !0,
                        source: !0,
                        track: !0,
                        wbr: !0,
                    },
                );
                function ye(e, t) {
                    if (t) {
                        if (
                            ge[e] &&
                            (null != t.children ||
                                null != t.dangerouslySetInnerHTML)
                        )
                            throw Error(l(137, e));
                        if (null != t.dangerouslySetInnerHTML) {
                            if (null != t.children) throw Error(l(60));
                            if (
                                "object" !== typeof t.dangerouslySetInnerHTML ||
                                !("__html" in t.dangerouslySetInnerHTML)
                            )
                                throw Error(l(61));
                        }
                        if (null != t.style && "object" !== typeof t.style)
                            throw Error(l(62));
                    }
                }
                function be(e, t) {
                    if (-1 === e.indexOf("-")) return "string" === typeof t.is;
                    switch (e) {
                        case "annotation-xml":
                        case "color-profile":
                        case "font-face":
                        case "font-face-src":
                        case "font-face-uri":
                        case "font-face-format":
                        case "font-face-name":
                        case "missing-glyph":
                            return !1;
                        default:
                            return !0;
                    }
                }
                var xe = null;
                function we(e) {
                    return (
                        (e = e.target || e.srcElement || window)
                            .correspondingUseElement &&
                            (e = e.correspondingUseElement),
                        3 === e.nodeType ? e.parentNode : e
                    );
                }
                var je = null,
                    ke = null,
                    Se = null;
                function Ne(e) {
                    if ((e = ba(e))) {
                        if ("function" !== typeof je) throw Error(l(280));
                        var t = e.stateNode;
                        t && ((t = wa(t)), je(e.stateNode, e.type, t));
                    }
                }
                function Ce(e) {
                    ke ? (Se ? Se.push(e) : (Se = [e])) : (ke = e);
                }
                function Ee() {
                    if (ke) {
                        var e = ke,
                            t = Se;
                        if (((Se = ke = null), Ne(e), t))
                            for (e = 0; e < t.length; e++) Ne(t[e]);
                    }
                }
                function Le(e, t) {
                    return e(t);
                }
                function _e() {}
                var Pe = !1;
                function Oe(e, t, n) {
                    if (Pe) return e(t, n);
                    Pe = !0;
                    try {
                        return Le(e, t, n);
                    } finally {
                        (Pe = !1), (null !== ke || null !== Se) && (_e(), Ee());
                    }
                }
                function ze(e, t) {
                    var n = e.stateNode;
                    if (null === n) return null;
                    var r = wa(n);
                    if (null === r) return null;
                    n = r[t];
                    e: switch (t) {
                        case "onClick":
                        case "onClickCapture":
                        case "onDoubleClick":
                        case "onDoubleClickCapture":
                        case "onMouseDown":
                        case "onMouseDownCapture":
                        case "onMouseMove":
                        case "onMouseMoveCapture":
                        case "onMouseUp":
                        case "onMouseUpCapture":
                        case "onMouseEnter":
                            (r = !r.disabled) ||
                                (r = !(
                                    "button" === (e = e.type) ||
                                    "input" === e ||
                                    "select" === e ||
                                    "textarea" === e
                                )),
                                (e = !r);
                            break e;
                        default:
                            e = !1;
                    }
                    if (e) return null;
                    if (n && "function" !== typeof n)
                        throw Error(l(231, t, typeof n));
                    return n;
                }
                var Me = !1;
                if (c)
                    try {
                        var Re = {};
                        Object.defineProperty(Re, "passive", {
                            get: function () {
                                Me = !0;
                            },
                        }),
                            window.addEventListener("test", Re, Re),
                            window.removeEventListener("test", Re, Re);
                    } catch (ce) {
                        Me = !1;
                    }
                function Te(e, t, n, r, a, l, i, o, s) {
                    var u = Array.prototype.slice.call(arguments, 3);
                    try {
                        t.apply(n, u);
                    } catch (c) {
                        this.onError(c);
                    }
                }
                var Fe = !1,
                    Ie = null,
                    De = !1,
                    Ue = null,
                    Be = {
                        onError: function (e) {
                            (Fe = !0), (Ie = e);
                        },
                    };
                function Ae(e, t, n, r, a, l, i, o, s) {
                    (Fe = !1), (Ie = null), Te.apply(Be, arguments);
                }
                function Ve(e) {
                    var t = e,
                        n = e;
                    if (e.alternate) for (; t.return; ) t = t.return;
                    else {
                        e = t;
                        do {
                            0 !== (4098 & (t = e).flags) && (n = t.return),
                                (e = t.return);
                        } while (e);
                    }
                    return 3 === t.tag ? n : null;
                }
                function $e(e) {
                    if (13 === e.tag) {
                        var t = e.memoizedState;
                        if (
                            (null === t &&
                                null !== (e = e.alternate) &&
                                (t = e.memoizedState),
                            null !== t)
                        )
                            return t.dehydrated;
                    }
                    return null;
                }
                function He(e) {
                    if (Ve(e) !== e) throw Error(l(188));
                }
                function Ke(e) {
                    return null !==
                        (e = (function (e) {
                            var t = e.alternate;
                            if (!t) {
                                if (null === (t = Ve(e))) throw Error(l(188));
                                return t !== e ? null : e;
                            }
                            for (var n = e, r = t; ; ) {
                                var a = n.return;
                                if (null === a) break;
                                var i = a.alternate;
                                if (null === i) {
                                    if (null !== (r = a.return)) {
                                        n = r;
                                        continue;
                                    }
                                    break;
                                }
                                if (a.child === i.child) {
                                    for (i = a.child; i; ) {
                                        if (i === n) return He(a), e;
                                        if (i === r) return He(a), t;
                                        i = i.sibling;
                                    }
                                    throw Error(l(188));
                                }
                                if (n.return !== r.return) (n = a), (r = i);
                                else {
                                    for (var o = !1, s = a.child; s; ) {
                                        if (s === n) {
                                            (o = !0), (n = a), (r = i);
                                            break;
                                        }
                                        if (s === r) {
                                            (o = !0), (r = a), (n = i);
                                            break;
                                        }
                                        s = s.sibling;
                                    }
                                    if (!o) {
                                        for (s = i.child; s; ) {
                                            if (s === n) {
                                                (o = !0), (n = i), (r = a);
                                                break;
                                            }
                                            if (s === r) {
                                                (o = !0), (r = i), (n = a);
                                                break;
                                            }
                                            s = s.sibling;
                                        }
                                        if (!o) throw Error(l(189));
                                    }
                                }
                                if (n.alternate !== r) throw Error(l(190));
                            }
                            if (3 !== n.tag) throw Error(l(188));
                            return n.stateNode.current === n ? e : t;
                        })(e))
                        ? We(e)
                        : null;
                }
                function We(e) {
                    if (5 === e.tag || 6 === e.tag) return e;
                    for (e = e.child; null !== e; ) {
                        var t = We(e);
                        if (null !== t) return t;
                        e = e.sibling;
                    }
                    return null;
                }
                var Qe = a.unstable_scheduleCallback,
                    qe = a.unstable_cancelCallback,
                    Ge = a.unstable_shouldYield,
                    Ye = a.unstable_requestPaint,
                    Xe = a.unstable_now,
                    Ze = a.unstable_getCurrentPriorityLevel,
                    Je = a.unstable_ImmediatePriority,
                    et = a.unstable_UserBlockingPriority,
                    tt = a.unstable_NormalPriority,
                    nt = a.unstable_LowPriority,
                    rt = a.unstable_IdlePriority,
                    at = null,
                    lt = null;
                var it = Math.clz32
                        ? Math.clz32
                        : function (e) {
                              return (
                                  (e >>>= 0),
                                  0 === e ? 32 : (31 - ((ot(e) / st) | 0)) | 0
                              );
                          },
                    ot = Math.log,
                    st = Math.LN2;
                var ut = 64,
                    ct = 4194304;
                function dt(e) {
                    switch (e & -e) {
                        case 1:
                            return 1;
                        case 2:
                            return 2;
                        case 4:
                            return 4;
                        case 8:
                            return 8;
                        case 16:
                            return 16;
                        case 32:
                            return 32;
                        case 64:
                        case 128:
                        case 256:
                        case 512:
                        case 1024:
                        case 2048:
                        case 4096:
                        case 8192:
                        case 16384:
                        case 32768:
                        case 65536:
                        case 131072:
                        case 262144:
                        case 524288:
                        case 1048576:
                        case 2097152:
                            return 4194240 & e;
                        case 4194304:
                        case 8388608:
                        case 16777216:
                        case 33554432:
                        case 67108864:
                            return 130023424 & e;
                        case 134217728:
                            return 134217728;
                        case 268435456:
                            return 268435456;
                        case 536870912:
                            return 536870912;
                        case 1073741824:
                            return 1073741824;
                        default:
                            return e;
                    }
                }
                function ft(e, t) {
                    var n = e.pendingLanes;
                    if (0 === n) return 0;
                    var r = 0,
                        a = e.suspendedLanes,
                        l = e.pingedLanes,
                        i = 268435455 & n;
                    if (0 !== i) {
                        var o = i & ~a;
                        0 !== o ? (r = dt(o)) : 0 !== (l &= i) && (r = dt(l));
                    } else
                        0 !== (i = n & ~a)
                            ? (r = dt(i))
                            : 0 !== l && (r = dt(l));
                    if (0 === r) return 0;
                    if (
                        0 !== t &&
                        t !== r &&
                        0 === (t & a) &&
                        ((a = r & -r) >= (l = t & -t) ||
                            (16 === a && 0 !== (4194240 & l)))
                    )
                        return t;
                    if (
                        (0 !== (4 & r) && (r |= 16 & n),
                        0 !== (t = e.entangledLanes))
                    )
                        for (e = e.entanglements, t &= r; 0 < t; )
                            (a = 1 << (n = 31 - it(t))), (r |= e[n]), (t &= ~a);
                    return r;
                }
                function pt(e, t) {
                    switch (e) {
                        case 1:
                        case 2:
                        case 4:
                            return t + 250;
                        case 8:
                        case 16:
                        case 32:
                        case 64:
                        case 128:
                        case 256:
                        case 512:
                        case 1024:
                        case 2048:
                        case 4096:
                        case 8192:
                        case 16384:
                        case 32768:
                        case 65536:
                        case 131072:
                        case 262144:
                        case 524288:
                        case 1048576:
                        case 2097152:
                            return t + 5e3;
                        default:
                            return -1;
                    }
                }
                function mt(e) {
                    return 0 !== (e = -1073741825 & e.pendingLanes)
                        ? e
                        : 1073741824 & e
                        ? 1073741824
                        : 0;
                }
                function ht() {
                    var e = ut;
                    return 0 === (4194240 & (ut <<= 1)) && (ut = 64), e;
                }
                function vt(e) {
                    for (var t = [], n = 0; 31 > n; n++) t.push(e);
                    return t;
                }
                function gt(e, t, n) {
                    (e.pendingLanes |= t),
                        536870912 !== t &&
                            ((e.suspendedLanes = 0), (e.pingedLanes = 0)),
                        ((e = e.eventTimes)[(t = 31 - it(t))] = n);
                }
                function yt(e, t) {
                    var n = (e.entangledLanes |= t);
                    for (e = e.entanglements; n; ) {
                        var r = 31 - it(n),
                            a = 1 << r;
                        (a & t) | (e[r] & t) && (e[r] |= t), (n &= ~a);
                    }
                }
                var bt = 0;
                function xt(e) {
                    return 1 < (e &= -e)
                        ? 4 < e
                            ? 0 !== (268435455 & e)
                                ? 16
                                : 536870912
                            : 4
                        : 1;
                }
                var wt,
                    jt,
                    kt,
                    St,
                    Nt,
                    Ct = !1,
                    Et = [],
                    Lt = null,
                    _t = null,
                    Pt = null,
                    Ot = new Map(),
                    zt = new Map(),
                    Mt = [],
                    Rt =
                        "mousedown mouseup touchcancel touchend touchstart auxclick dblclick pointercancel pointerdown pointerup dragend dragstart drop compositionend compositionstart keydown keypress keyup input textInput copy cut paste click change contextmenu reset submit".split(
                            " ",
                        );
                function Tt(e, t) {
                    switch (e) {
                        case "focusin":
                        case "focusout":
                            Lt = null;
                            break;
                        case "dragenter":
                        case "dragleave":
                            _t = null;
                            break;
                        case "mouseover":
                        case "mouseout":
                            Pt = null;
                            break;
                        case "pointerover":
                        case "pointerout":
                            Ot.delete(t.pointerId);
                            break;
                        case "gotpointercapture":
                        case "lostpointercapture":
                            zt.delete(t.pointerId);
                    }
                }
                function Ft(e, t, n, r, a, l) {
                    return null === e || e.nativeEvent !== l
                        ? ((e = {
                              blockedOn: t,
                              domEventName: n,
                              eventSystemFlags: r,
                              nativeEvent: l,
                              targetContainers: [a],
                          }),
                          null !== t && null !== (t = ba(t)) && jt(t),
                          e)
                        : ((e.eventSystemFlags |= r),
                          (t = e.targetContainers),
                          null !== a && -1 === t.indexOf(a) && t.push(a),
                          e);
                }
                function It(e) {
                    var t = ya(e.target);
                    if (null !== t) {
                        var n = Ve(t);
                        if (null !== n)
                            if (13 === (t = n.tag)) {
                                if (null !== (t = $e(n)))
                                    return (
                                        (e.blockedOn = t),
                                        void Nt(e.priority, function () {
                                            kt(n);
                                        })
                                    );
                            } else if (
                                3 === t &&
                                n.stateNode.current.memoizedState.isDehydrated
                            )
                                return void (e.blockedOn =
                                    3 === n.tag
                                        ? n.stateNode.containerInfo
                                        : null);
                    }
                    e.blockedOn = null;
                }
                function Dt(e) {
                    if (null !== e.blockedOn) return !1;
                    for (var t = e.targetContainers; 0 < t.length; ) {
                        var n = Gt(
                            e.domEventName,
                            e.eventSystemFlags,
                            t[0],
                            e.nativeEvent,
                        );
                        if (null !== n)
                            return (
                                null !== (t = ba(n)) && jt(t),
                                (e.blockedOn = n),
                                !1
                            );
                        var r = new (n = e.nativeEvent).constructor(n.type, n);
                        (xe = r),
                            n.target.dispatchEvent(r),
                            (xe = null),
                            t.shift();
                    }
                    return !0;
                }
                function Ut(e, t, n) {
                    Dt(e) && n.delete(t);
                }
                function Bt() {
                    (Ct = !1),
                        null !== Lt && Dt(Lt) && (Lt = null),
                        null !== _t && Dt(_t) && (_t = null),
                        null !== Pt && Dt(Pt) && (Pt = null),
                        Ot.forEach(Ut),
                        zt.forEach(Ut);
                }
                function At(e, t) {
                    e.blockedOn === t &&
                        ((e.blockedOn = null),
                        Ct ||
                            ((Ct = !0),
                            a.unstable_scheduleCallback(
                                a.unstable_NormalPriority,
                                Bt,
                            )));
                }
                function Vt(e) {
                    function t(t) {
                        return At(t, e);
                    }
                    if (0 < Et.length) {
                        At(Et[0], e);
                        for (var n = 1; n < Et.length; n++) {
                            var r = Et[n];
                            r.blockedOn === e && (r.blockedOn = null);
                        }
                    }
                    for (
                        null !== Lt && At(Lt, e),
                            null !== _t && At(_t, e),
                            null !== Pt && At(Pt, e),
                            Ot.forEach(t),
                            zt.forEach(t),
                            n = 0;
                        n < Mt.length;
                        n++
                    )
                        (r = Mt[n]).blockedOn === e && (r.blockedOn = null);
                    for (; 0 < Mt.length && null === (n = Mt[0]).blockedOn; )
                        It(n), null === n.blockedOn && Mt.shift();
                }
                var $t = x.ReactCurrentBatchConfig,
                    Ht = !0;
                function Kt(e, t, n, r) {
                    var a = bt,
                        l = $t.transition;
                    $t.transition = null;
                    try {
                        (bt = 1), Qt(e, t, n, r);
                    } finally {
                        (bt = a), ($t.transition = l);
                    }
                }
                function Wt(e, t, n, r) {
                    var a = bt,
                        l = $t.transition;
                    $t.transition = null;
                    try {
                        (bt = 4), Qt(e, t, n, r);
                    } finally {
                        (bt = a), ($t.transition = l);
                    }
                }
                function Qt(e, t, n, r) {
                    if (Ht) {
                        var a = Gt(e, t, n, r);
                        if (null === a) Hr(e, t, r, qt, n), Tt(e, r);
                        else if (
                            (function (e, t, n, r, a) {
                                switch (t) {
                                    case "focusin":
                                        return (Lt = Ft(Lt, e, t, n, r, a)), !0;
                                    case "dragenter":
                                        return (_t = Ft(_t, e, t, n, r, a)), !0;
                                    case "mouseover":
                                        return (Pt = Ft(Pt, e, t, n, r, a)), !0;
                                    case "pointerover":
                                        var l = a.pointerId;
                                        return (
                                            Ot.set(
                                                l,
                                                Ft(
                                                    Ot.get(l) || null,
                                                    e,
                                                    t,
                                                    n,
                                                    r,
                                                    a,
                                                ),
                                            ),
                                            !0
                                        );
                                    case "gotpointercapture":
                                        return (
                                            (l = a.pointerId),
                                            zt.set(
                                                l,
                                                Ft(
                                                    zt.get(l) || null,
                                                    e,
                                                    t,
                                                    n,
                                                    r,
                                                    a,
                                                ),
                                            ),
                                            !0
                                        );
                                }
                                return !1;
                            })(a, e, t, n, r)
                        )
                            r.stopPropagation();
                        else if ((Tt(e, r), 4 & t && -1 < Rt.indexOf(e))) {
                            for (; null !== a; ) {
                                var l = ba(a);
                                if (
                                    (null !== l && wt(l),
                                    null === (l = Gt(e, t, n, r)) &&
                                        Hr(e, t, r, qt, n),
                                    l === a)
                                )
                                    break;
                                a = l;
                            }
                            null !== a && r.stopPropagation();
                        } else Hr(e, t, r, null, n);
                    }
                }
                var qt = null;
                function Gt(e, t, n, r) {
                    if (((qt = null), null !== (e = ya((e = we(r))))))
                        if (null === (t = Ve(e))) e = null;
                        else if (13 === (n = t.tag)) {
                            if (null !== (e = $e(t))) return e;
                            e = null;
                        } else if (3 === n) {
                            if (t.stateNode.current.memoizedState.isDehydrated)
                                return 3 === t.tag
                                    ? t.stateNode.containerInfo
                                    : null;
                            e = null;
                        } else t !== e && (e = null);
                    return (qt = e), null;
                }
                function Yt(e) {
                    switch (e) {
                        case "cancel":
                        case "click":
                        case "close":
                        case "contextmenu":
                        case "copy":
                        case "cut":
                        case "auxclick":
                        case "dblclick":
                        case "dragend":
                        case "dragstart":
                        case "drop":
                        case "focusin":
                        case "focusout":
                        case "input":
                        case "invalid":
                        case "keydown":
                        case "keypress":
                        case "keyup":
                        case "mousedown":
                        case "mouseup":
                        case "paste":
                        case "pause":
                        case "play":
                        case "pointercancel":
                        case "pointerdown":
                        case "pointerup":
                        case "ratechange":
                        case "reset":
                        case "resize":
                        case "seeked":
                        case "submit":
                        case "touchcancel":
                        case "touchend":
                        case "touchstart":
                        case "volumechange":
                        case "change":
                        case "selectionchange":
                        case "textInput":
                        case "compositionstart":
                        case "compositionend":
                        case "compositionupdate":
                        case "beforeblur":
                        case "afterblur":
                        case "beforeinput":
                        case "blur":
                        case "fullscreenchange":
                        case "focus":
                        case "hashchange":
                        case "popstate":
                        case "select":
                        case "selectstart":
                            return 1;
                        case "drag":
                        case "dragenter":
                        case "dragexit":
                        case "dragleave":
                        case "dragover":
                        case "mousemove":
                        case "mouseout":
                        case "mouseover":
                        case "pointermove":
                        case "pointerout":
                        case "pointerover":
                        case "scroll":
                        case "toggle":
                        case "touchmove":
                        case "wheel":
                        case "mouseenter":
                        case "mouseleave":
                        case "pointerenter":
                        case "pointerleave":
                            return 4;
                        case "message":
                            switch (Ze()) {
                                case Je:
                                    return 1;
                                case et:
                                    return 4;
                                case tt:
                                case nt:
                                    return 16;
                                case rt:
                                    return 536870912;
                                default:
                                    return 16;
                            }
                        default:
                            return 16;
                    }
                }
                var Xt = null,
                    Zt = null,
                    Jt = null;
                function en() {
                    if (Jt) return Jt;
                    var e,
                        t,
                        n = Zt,
                        r = n.length,
                        a = "value" in Xt ? Xt.value : Xt.textContent,
                        l = a.length;
                    for (e = 0; e < r && n[e] === a[e]; e++);
                    var i = r - e;
                    for (t = 1; t <= i && n[r - t] === a[l - t]; t++);
                    return (Jt = a.slice(e, 1 < t ? 1 - t : void 0));
                }
                function tn(e) {
                    var t = e.keyCode;
                    return (
                        "charCode" in e
                            ? 0 === (e = e.charCode) && 13 === t && (e = 13)
                            : (e = t),
                        10 === e && (e = 13),
                        32 <= e || 13 === e ? e : 0
                    );
                }
                function nn() {
                    return !0;
                }
                function rn() {
                    return !1;
                }
                function an(e) {
                    function t(t, n, r, a, l) {
                        for (var i in ((this._reactName = t),
                        (this._targetInst = r),
                        (this.type = n),
                        (this.nativeEvent = a),
                        (this.target = l),
                        (this.currentTarget = null),
                        e))
                            e.hasOwnProperty(i) &&
                                ((t = e[i]), (this[i] = t ? t(a) : a[i]));
                        return (
                            (this.isDefaultPrevented = (
                                null != a.defaultPrevented
                                    ? a.defaultPrevented
                                    : !1 === a.returnValue
                            )
                                ? nn
                                : rn),
                            (this.isPropagationStopped = rn),
                            this
                        );
                    }
                    return (
                        I(t.prototype, {
                            preventDefault: function () {
                                this.defaultPrevented = !0;
                                var e = this.nativeEvent;
                                e &&
                                    (e.preventDefault
                                        ? e.preventDefault()
                                        : "unknown" !== typeof e.returnValue &&
                                          (e.returnValue = !1),
                                    (this.isDefaultPrevented = nn));
                            },
                            stopPropagation: function () {
                                var e = this.nativeEvent;
                                e &&
                                    (e.stopPropagation
                                        ? e.stopPropagation()
                                        : "unknown" !== typeof e.cancelBubble &&
                                          (e.cancelBubble = !0),
                                    (this.isPropagationStopped = nn));
                            },
                            persist: function () {},
                            isPersistent: nn,
                        }),
                        t
                    );
                }
                var ln,
                    on,
                    sn,
                    un = {
                        eventPhase: 0,
                        bubbles: 0,
                        cancelable: 0,
                        timeStamp: function (e) {
                            return e.timeStamp || Date.now();
                        },
                        defaultPrevented: 0,
                        isTrusted: 0,
                    },
                    cn = an(un),
                    dn = I({}, un, { view: 0, detail: 0 }),
                    fn = an(dn),
                    pn = I({}, dn, {
                        screenX: 0,
                        screenY: 0,
                        clientX: 0,
                        clientY: 0,
                        pageX: 0,
                        pageY: 0,
                        ctrlKey: 0,
                        shiftKey: 0,
                        altKey: 0,
                        metaKey: 0,
                        getModifierState: Nn,
                        button: 0,
                        buttons: 0,
                        relatedTarget: function (e) {
                            return void 0 === e.relatedTarget
                                ? e.fromElement === e.srcElement
                                    ? e.toElement
                                    : e.fromElement
                                : e.relatedTarget;
                        },
                        movementX: function (e) {
                            return "movementX" in e
                                ? e.movementX
                                : (e !== sn &&
                                      (sn && "mousemove" === e.type
                                          ? ((ln = e.screenX - sn.screenX),
                                            (on = e.screenY - sn.screenY))
                                          : (on = ln = 0),
                                      (sn = e)),
                                  ln);
                        },
                        movementY: function (e) {
                            return "movementY" in e ? e.movementY : on;
                        },
                    }),
                    mn = an(pn),
                    hn = an(I({}, pn, { dataTransfer: 0 })),
                    vn = an(I({}, dn, { relatedTarget: 0 })),
                    gn = an(
                        I({}, un, {
                            animationName: 0,
                            elapsedTime: 0,
                            pseudoElement: 0,
                        }),
                    ),
                    yn = I({}, un, {
                        clipboardData: function (e) {
                            return "clipboardData" in e
                                ? e.clipboardData
                                : window.clipboardData;
                        },
                    }),
                    bn = an(yn),
                    xn = an(I({}, un, { data: 0 })),
                    wn = {
                        Esc: "Escape",
                        Spacebar: " ",
                        Left: "ArrowLeft",
                        Up: "ArrowUp",
                        Right: "ArrowRight",
                        Down: "ArrowDown",
                        Del: "Delete",
                        Win: "OS",
                        Menu: "ContextMenu",
                        Apps: "ContextMenu",
                        Scroll: "ScrollLock",
                        MozPrintableKey: "Unidentified",
                    },
                    jn = {
                        8: "Backspace",
                        9: "Tab",
                        12: "Clear",
                        13: "Enter",
                        16: "Shift",
                        17: "Control",
                        18: "Alt",
                        19: "Pause",
                        20: "CapsLock",
                        27: "Escape",
                        32: " ",
                        33: "PageUp",
                        34: "PageDown",
                        35: "End",
                        36: "Home",
                        37: "ArrowLeft",
                        38: "ArrowUp",
                        39: "ArrowRight",
                        40: "ArrowDown",
                        45: "Insert",
                        46: "Delete",
                        112: "F1",
                        113: "F2",
                        114: "F3",
                        115: "F4",
                        116: "F5",
                        117: "F6",
                        118: "F7",
                        119: "F8",
                        120: "F9",
                        121: "F10",
                        122: "F11",
                        123: "F12",
                        144: "NumLock",
                        145: "ScrollLock",
                        224: "Meta",
                    },
                    kn = {
                        Alt: "altKey",
                        Control: "ctrlKey",
                        Meta: "metaKey",
                        Shift: "shiftKey",
                    };
                function Sn(e) {
                    var t = this.nativeEvent;
                    return t.getModifierState
                        ? t.getModifierState(e)
                        : !!(e = kn[e]) && !!t[e];
                }
                function Nn() {
                    return Sn;
                }
                var Cn = I({}, dn, {
                        key: function (e) {
                            if (e.key) {
                                var t = wn[e.key] || e.key;
                                if ("Unidentified" !== t) return t;
                            }
                            return "keypress" === e.type
                                ? 13 === (e = tn(e))
                                    ? "Enter"
                                    : String.fromCharCode(e)
                                : "keydown" === e.type || "keyup" === e.type
                                ? jn[e.keyCode] || "Unidentified"
                                : "";
                        },
                        code: 0,
                        location: 0,
                        ctrlKey: 0,
                        shiftKey: 0,
                        altKey: 0,
                        metaKey: 0,
                        repeat: 0,
                        locale: 0,
                        getModifierState: Nn,
                        charCode: function (e) {
                            return "keypress" === e.type ? tn(e) : 0;
                        },
                        keyCode: function (e) {
                            return "keydown" === e.type || "keyup" === e.type
                                ? e.keyCode
                                : 0;
                        },
                        which: function (e) {
                            return "keypress" === e.type
                                ? tn(e)
                                : "keydown" === e.type || "keyup" === e.type
                                ? e.keyCode
                                : 0;
                        },
                    }),
                    En = an(Cn),
                    Ln = an(
                        I({}, pn, {
                            pointerId: 0,
                            width: 0,
                            height: 0,
                            pressure: 0,
                            tangentialPressure: 0,
                            tiltX: 0,
                            tiltY: 0,
                            twist: 0,
                            pointerType: 0,
                            isPrimary: 0,
                        }),
                    ),
                    _n = an(
                        I({}, dn, {
                            touches: 0,
                            targetTouches: 0,
                            changedTouches: 0,
                            altKey: 0,
                            metaKey: 0,
                            ctrlKey: 0,
                            shiftKey: 0,
                            getModifierState: Nn,
                        }),
                    ),
                    Pn = an(
                        I({}, un, {
                            propertyName: 0,
                            elapsedTime: 0,
                            pseudoElement: 0,
                        }),
                    ),
                    On = I({}, pn, {
                        deltaX: function (e) {
                            return "deltaX" in e
                                ? e.deltaX
                                : "wheelDeltaX" in e
                                ? -e.wheelDeltaX
                                : 0;
                        },
                        deltaY: function (e) {
                            return "deltaY" in e
                                ? e.deltaY
                                : "wheelDeltaY" in e
                                ? -e.wheelDeltaY
                                : "wheelDelta" in e
                                ? -e.wheelDelta
                                : 0;
                        },
                        deltaZ: 0,
                        deltaMode: 0,
                    }),
                    zn = an(On),
                    Mn = [9, 13, 27, 32],
                    Rn = c && "CompositionEvent" in window,
                    Tn = null;
                c && "documentMode" in document && (Tn = document.documentMode);
                var Fn = c && "TextEvent" in window && !Tn,
                    In = c && (!Rn || (Tn && 8 < Tn && 11 >= Tn)),
                    Dn = String.fromCharCode(32),
                    Un = !1;
                function Bn(e, t) {
                    switch (e) {
                        case "keyup":
                            return -1 !== Mn.indexOf(t.keyCode);
                        case "keydown":
                            return 229 !== t.keyCode;
                        case "keypress":
                        case "mousedown":
                        case "focusout":
                            return !0;
                        default:
                            return !1;
                    }
                }
                function An(e) {
                    return "object" === typeof (e = e.detail) && "data" in e
                        ? e.data
                        : null;
                }
                var Vn = !1;
                var $n = {
                    color: !0,
                    date: !0,
                    datetime: !0,
                    "datetime-local": !0,
                    email: !0,
                    month: !0,
                    number: !0,
                    password: !0,
                    range: !0,
                    search: !0,
                    tel: !0,
                    text: !0,
                    time: !0,
                    url: !0,
                    week: !0,
                };
                function Hn(e) {
                    var t = e && e.nodeName && e.nodeName.toLowerCase();
                    return "input" === t ? !!$n[e.type] : "textarea" === t;
                }
                function Kn(e, t, n, r) {
                    Ce(r),
                        0 < (t = Wr(t, "onChange")).length &&
                            ((n = new cn("onChange", "change", null, n, r)),
                            e.push({ event: n, listeners: t }));
                }
                var Wn = null,
                    Qn = null;
                function qn(e) {
                    Dr(e, 0);
                }
                function Gn(e) {
                    if (Q(xa(e))) return e;
                }
                function Yn(e, t) {
                    if ("change" === e) return t;
                }
                var Xn = !1;
                if (c) {
                    var Zn;
                    if (c) {
                        var Jn = "oninput" in document;
                        if (!Jn) {
                            var er = document.createElement("div");
                            er.setAttribute("oninput", "return;"),
                                (Jn = "function" === typeof er.oninput);
                        }
                        Zn = Jn;
                    } else Zn = !1;
                    Xn =
                        Zn &&
                        (!document.documentMode || 9 < document.documentMode);
                }
                function tr() {
                    Wn &&
                        (Wn.detachEvent("onpropertychange", nr),
                        (Qn = Wn = null));
                }
                function nr(e) {
                    if ("value" === e.propertyName && Gn(Qn)) {
                        var t = [];
                        Kn(t, Qn, e, we(e)), Oe(qn, t);
                    }
                }
                function rr(e, t, n) {
                    "focusin" === e
                        ? (tr(),
                          (Qn = n),
                          (Wn = t).attachEvent("onpropertychange", nr))
                        : "focusout" === e && tr();
                }
                function ar(e) {
                    if (
                        "selectionchange" === e ||
                        "keyup" === e ||
                        "keydown" === e
                    )
                        return Gn(Qn);
                }
                function lr(e, t) {
                    if ("click" === e) return Gn(t);
                }
                function ir(e, t) {
                    if ("input" === e || "change" === e) return Gn(t);
                }
                var or =
                    "function" === typeof Object.is
                        ? Object.is
                        : function (e, t) {
                              return (
                                  (e === t && (0 !== e || 1 / e === 1 / t)) ||
                                  (e !== e && t !== t)
                              );
                          };
                function sr(e, t) {
                    if (or(e, t)) return !0;
                    if (
                        "object" !== typeof e ||
                        null === e ||
                        "object" !== typeof t ||
                        null === t
                    )
                        return !1;
                    var n = Object.keys(e),
                        r = Object.keys(t);
                    if (n.length !== r.length) return !1;
                    for (r = 0; r < n.length; r++) {
                        var a = n[r];
                        if (!d.call(t, a) || !or(e[a], t[a])) return !1;
                    }
                    return !0;
                }
                function ur(e) {
                    for (; e && e.firstChild; ) e = e.firstChild;
                    return e;
                }
                function cr(e, t) {
                    var n,
                        r = ur(e);
                    for (e = 0; r; ) {
                        if (3 === r.nodeType) {
                            if (
                                ((n = e + r.textContent.length),
                                e <= t && n >= t)
                            )
                                return { node: r, offset: t - e };
                            e = n;
                        }
                        e: {
                            for (; r; ) {
                                if (r.nextSibling) {
                                    r = r.nextSibling;
                                    break e;
                                }
                                r = r.parentNode;
                            }
                            r = void 0;
                        }
                        r = ur(r);
                    }
                }
                function dr(e, t) {
                    return (
                        !(!e || !t) &&
                        (e === t ||
                            ((!e || 3 !== e.nodeType) &&
                                (t && 3 === t.nodeType
                                    ? dr(e, t.parentNode)
                                    : "contains" in e
                                    ? e.contains(t)
                                    : !!e.compareDocumentPosition &&
                                      !!(16 & e.compareDocumentPosition(t)))))
                    );
                }
                function fr() {
                    for (
                        var e = window, t = q();
                        t instanceof e.HTMLIFrameElement;

                    ) {
                        try {
                            var n =
                                "string" ===
                                typeof t.contentWindow.location.href;
                        } catch (r) {
                            n = !1;
                        }
                        if (!n) break;
                        t = q((e = t.contentWindow).document);
                    }
                    return t;
                }
                function pr(e) {
                    var t = e && e.nodeName && e.nodeName.toLowerCase();
                    return (
                        t &&
                        (("input" === t &&
                            ("text" === e.type ||
                                "search" === e.type ||
                                "tel" === e.type ||
                                "url" === e.type ||
                                "password" === e.type)) ||
                            "textarea" === t ||
                            "true" === e.contentEditable)
                    );
                }
                function mr(e) {
                    var t = fr(),
                        n = e.focusedElem,
                        r = e.selectionRange;
                    if (
                        t !== n &&
                        n &&
                        n.ownerDocument &&
                        dr(n.ownerDocument.documentElement, n)
                    ) {
                        if (null !== r && pr(n))
                            if (
                                ((t = r.start),
                                void 0 === (e = r.end) && (e = t),
                                "selectionStart" in n)
                            )
                                (n.selectionStart = t),
                                    (n.selectionEnd = Math.min(
                                        e,
                                        n.value.length,
                                    ));
                            else if (
                                (e =
                                    ((t = n.ownerDocument || document) &&
                                        t.defaultView) ||
                                    window).getSelection
                            ) {
                                e = e.getSelection();
                                var a = n.textContent.length,
                                    l = Math.min(r.start, a);
                                (r = void 0 === r.end ? l : Math.min(r.end, a)),
                                    !e.extend &&
                                        l > r &&
                                        ((a = r), (r = l), (l = a)),
                                    (a = cr(n, l));
                                var i = cr(n, r);
                                a &&
                                    i &&
                                    (1 !== e.rangeCount ||
                                        e.anchorNode !== a.node ||
                                        e.anchorOffset !== a.offset ||
                                        e.focusNode !== i.node ||
                                        e.focusOffset !== i.offset) &&
                                    ((t = t.createRange()).setStart(
                                        a.node,
                                        a.offset,
                                    ),
                                    e.removeAllRanges(),
                                    l > r
                                        ? (e.addRange(t),
                                          e.extend(i.node, i.offset))
                                        : (t.setEnd(i.node, i.offset),
                                          e.addRange(t)));
                            }
                        for (t = [], e = n; (e = e.parentNode); )
                            1 === e.nodeType &&
                                t.push({
                                    element: e,
                                    left: e.scrollLeft,
                                    top: e.scrollTop,
                                });
                        for (
                            "function" === typeof n.focus && n.focus(), n = 0;
                            n < t.length;
                            n++
                        )
                            ((e = t[n]).element.scrollLeft = e.left),
                                (e.element.scrollTop = e.top);
                    }
                }
                var hr =
                        c &&
                        "documentMode" in document &&
                        11 >= document.documentMode,
                    vr = null,
                    gr = null,
                    yr = null,
                    br = !1;
                function xr(e, t, n) {
                    var r =
                        n.window === n
                            ? n.document
                            : 9 === n.nodeType
                            ? n
                            : n.ownerDocument;
                    br ||
                        null == vr ||
                        vr !== q(r) ||
                        ("selectionStart" in (r = vr) && pr(r)
                            ? (r = {
                                  start: r.selectionStart,
                                  end: r.selectionEnd,
                              })
                            : (r = {
                                  anchorNode: (r = (
                                      (r.ownerDocument &&
                                          r.ownerDocument.defaultView) ||
                                      window
                                  ).getSelection()).anchorNode,
                                  anchorOffset: r.anchorOffset,
                                  focusNode: r.focusNode,
                                  focusOffset: r.focusOffset,
                              }),
                        (yr && sr(yr, r)) ||
                            ((yr = r),
                            0 < (r = Wr(gr, "onSelect")).length &&
                                ((t = new cn("onSelect", "select", null, t, n)),
                                e.push({ event: t, listeners: r }),
                                (t.target = vr))));
                }
                function wr(e, t) {
                    var n = {};
                    return (
                        (n[e.toLowerCase()] = t.toLowerCase()),
                        (n["Webkit" + e] = "webkit" + t),
                        (n["Moz" + e] = "moz" + t),
                        n
                    );
                }
                var jr = {
                        animationend: wr("Animation", "AnimationEnd"),
                        animationiteration: wr(
                            "Animation",
                            "AnimationIteration",
                        ),
                        animationstart: wr("Animation", "AnimationStart"),
                        transitionend: wr("Transition", "TransitionEnd"),
                    },
                    kr = {},
                    Sr = {};
                function Nr(e) {
                    if (kr[e]) return kr[e];
                    if (!jr[e]) return e;
                    var t,
                        n = jr[e];
                    for (t in n)
                        if (n.hasOwnProperty(t) && t in Sr)
                            return (kr[e] = n[t]);
                    return e;
                }
                c &&
                    ((Sr = document.createElement("div").style),
                    "AnimationEvent" in window ||
                        (delete jr.animationend.animation,
                        delete jr.animationiteration.animation,
                        delete jr.animationstart.animation),
                    "TransitionEvent" in window ||
                        delete jr.transitionend.transition);
                var Cr = Nr("animationend"),
                    Er = Nr("animationiteration"),
                    Lr = Nr("animationstart"),
                    _r = Nr("transitionend"),
                    Pr = new Map(),
                    Or =
                        "abort auxClick cancel canPlay canPlayThrough click close contextMenu copy cut drag dragEnd dragEnter dragExit dragLeave dragOver dragStart drop durationChange emptied encrypted ended error gotPointerCapture input invalid keyDown keyPress keyUp load loadedData loadedMetadata loadStart lostPointerCapture mouseDown mouseMove mouseOut mouseOver mouseUp paste pause play playing pointerCancel pointerDown pointerMove pointerOut pointerOver pointerUp progress rateChange reset resize seeked seeking stalled submit suspend timeUpdate touchCancel touchEnd touchStart volumeChange scroll toggle touchMove waiting wheel".split(
                            " ",
                        );
                function zr(e, t) {
                    Pr.set(e, t), s(t, [e]);
                }
                for (var Mr = 0; Mr < Or.length; Mr++) {
                    var Rr = Or[Mr];
                    zr(
                        Rr.toLowerCase(),
                        "on" + (Rr[0].toUpperCase() + Rr.slice(1)),
                    );
                }
                zr(Cr, "onAnimationEnd"),
                    zr(Er, "onAnimationIteration"),
                    zr(Lr, "onAnimationStart"),
                    zr("dblclick", "onDoubleClick"),
                    zr("focusin", "onFocus"),
                    zr("focusout", "onBlur"),
                    zr(_r, "onTransitionEnd"),
                    u("onMouseEnter", ["mouseout", "mouseover"]),
                    u("onMouseLeave", ["mouseout", "mouseover"]),
                    u("onPointerEnter", ["pointerout", "pointerover"]),
                    u("onPointerLeave", ["pointerout", "pointerover"]),
                    s(
                        "onChange",
                        "change click focusin focusout input keydown keyup selectionchange".split(
                            " ",
                        ),
                    ),
                    s(
                        "onSelect",
                        "focusout contextmenu dragend focusin keydown keyup mousedown mouseup selectionchange".split(
                            " ",
                        ),
                    ),
                    s("onBeforeInput", [
                        "compositionend",
                        "keypress",
                        "textInput",
                        "paste",
                    ]),
                    s(
                        "onCompositionEnd",
                        "compositionend focusout keydown keypress keyup mousedown".split(
                            " ",
                        ),
                    ),
                    s(
                        "onCompositionStart",
                        "compositionstart focusout keydown keypress keyup mousedown".split(
                            " ",
                        ),
                    ),
                    s(
                        "onCompositionUpdate",
                        "compositionupdate focusout keydown keypress keyup mousedown".split(
                            " ",
                        ),
                    );
                var Tr =
                        "abort canplay canplaythrough durationchange emptied encrypted ended error loadeddata loadedmetadata loadstart pause play playing progress ratechange resize seeked seeking stalled suspend timeupdate volumechange waiting".split(
                            " ",
                        ),
                    Fr = new Set(
                        "cancel close invalid load scroll toggle"
                            .split(" ")
                            .concat(Tr),
                    );
                function Ir(e, t, n) {
                    var r = e.type || "unknown-event";
                    (e.currentTarget = n),
                        (function (e, t, n, r, a, i, o, s, u) {
                            if ((Ae.apply(this, arguments), Fe)) {
                                if (!Fe) throw Error(l(198));
                                var c = Ie;
                                (Fe = !1),
                                    (Ie = null),
                                    De || ((De = !0), (Ue = c));
                            }
                        })(r, t, void 0, e),
                        (e.currentTarget = null);
                }
                function Dr(e, t) {
                    t = 0 !== (4 & t);
                    for (var n = 0; n < e.length; n++) {
                        var r = e[n],
                            a = r.event;
                        r = r.listeners;
                        e: {
                            var l = void 0;
                            if (t)
                                for (var i = r.length - 1; 0 <= i; i--) {
                                    var o = r[i],
                                        s = o.instance,
                                        u = o.currentTarget;
                                    if (
                                        ((o = o.listener),
                                        s !== l && a.isPropagationStopped())
                                    )
                                        break e;
                                    Ir(a, o, u), (l = s);
                                }
                            else
                                for (i = 0; i < r.length; i++) {
                                    if (
                                        ((s = (o = r[i]).instance),
                                        (u = o.currentTarget),
                                        (o = o.listener),
                                        s !== l && a.isPropagationStopped())
                                    )
                                        break e;
                                    Ir(a, o, u), (l = s);
                                }
                        }
                    }
                    if (De) throw ((e = Ue), (De = !1), (Ue = null), e);
                }
                function Ur(e, t) {
                    var n = t[ha];
                    void 0 === n && (n = t[ha] = new Set());
                    var r = e + "__bubble";
                    n.has(r) || ($r(t, e, 2, !1), n.add(r));
                }
                function Br(e, t, n) {
                    var r = 0;
                    t && (r |= 4), $r(n, e, r, t);
                }
                var Ar =
                    "_reactListening" + Math.random().toString(36).slice(2);
                function Vr(e) {
                    if (!e[Ar]) {
                        (e[Ar] = !0),
                            i.forEach(function (t) {
                                "selectionchange" !== t &&
                                    (Fr.has(t) || Br(t, !1, e), Br(t, !0, e));
                            });
                        var t = 9 === e.nodeType ? e : e.ownerDocument;
                        null === t ||
                            t[Ar] ||
                            ((t[Ar] = !0), Br("selectionchange", !1, t));
                    }
                }
                function $r(e, t, n, r) {
                    switch (Yt(t)) {
                        case 1:
                            var a = Kt;
                            break;
                        case 4:
                            a = Wt;
                            break;
                        default:
                            a = Qt;
                    }
                    (n = a.bind(null, t, n, e)),
                        (a = void 0),
                        !Me ||
                            ("touchstart" !== t &&
                                "touchmove" !== t &&
                                "wheel" !== t) ||
                            (a = !0),
                        r
                            ? void 0 !== a
                                ? e.addEventListener(t, n, {
                                      capture: !0,
                                      passive: a,
                                  })
                                : e.addEventListener(t, n, !0)
                            : void 0 !== a
                            ? e.addEventListener(t, n, { passive: a })
                            : e.addEventListener(t, n, !1);
                }
                function Hr(e, t, n, r, a) {
                    var l = r;
                    if (0 === (1 & t) && 0 === (2 & t) && null !== r)
                        e: for (;;) {
                            if (null === r) return;
                            var i = r.tag;
                            if (3 === i || 4 === i) {
                                var o = r.stateNode.containerInfo;
                                if (
                                    o === a ||
                                    (8 === o.nodeType && o.parentNode === a)
                                )
                                    break;
                                if (4 === i)
                                    for (i = r.return; null !== i; ) {
                                        var s = i.tag;
                                        if (
                                            (3 === s || 4 === s) &&
                                            ((s = i.stateNode.containerInfo) ===
                                                a ||
                                                (8 === s.nodeType &&
                                                    s.parentNode === a))
                                        )
                                            return;
                                        i = i.return;
                                    }
                                for (; null !== o; ) {
                                    if (null === (i = ya(o))) return;
                                    if (5 === (s = i.tag) || 6 === s) {
                                        r = l = i;
                                        continue e;
                                    }
                                    o = o.parentNode;
                                }
                            }
                            r = r.return;
                        }
                    Oe(function () {
                        var r = l,
                            a = we(n),
                            i = [];
                        e: {
                            var o = Pr.get(e);
                            if (void 0 !== o) {
                                var s = cn,
                                    u = e;
                                switch (e) {
                                    case "keypress":
                                        if (0 === tn(n)) break e;
                                    case "keydown":
                                    case "keyup":
                                        s = En;
                                        break;
                                    case "focusin":
                                        (u = "focus"), (s = vn);
                                        break;
                                    case "focusout":
                                        (u = "blur"), (s = vn);
                                        break;
                                    case "beforeblur":
                                    case "afterblur":
                                        s = vn;
                                        break;
                                    case "click":
                                        if (2 === n.button) break e;
                                    case "auxclick":
                                    case "dblclick":
                                    case "mousedown":
                                    case "mousemove":
                                    case "mouseup":
                                    case "mouseout":
                                    case "mouseover":
                                    case "contextmenu":
                                        s = mn;
                                        break;
                                    case "drag":
                                    case "dragend":
                                    case "dragenter":
                                    case "dragexit":
                                    case "dragleave":
                                    case "dragover":
                                    case "dragstart":
                                    case "drop":
                                        s = hn;
                                        break;
                                    case "touchcancel":
                                    case "touchend":
                                    case "touchmove":
                                    case "touchstart":
                                        s = _n;
                                        break;
                                    case Cr:
                                    case Er:
                                    case Lr:
                                        s = gn;
                                        break;
                                    case _r:
                                        s = Pn;
                                        break;
                                    case "scroll":
                                        s = fn;
                                        break;
                                    case "wheel":
                                        s = zn;
                                        break;
                                    case "copy":
                                    case "cut":
                                    case "paste":
                                        s = bn;
                                        break;
                                    case "gotpointercapture":
                                    case "lostpointercapture":
                                    case "pointercancel":
                                    case "pointerdown":
                                    case "pointermove":
                                    case "pointerout":
                                    case "pointerover":
                                    case "pointerup":
                                        s = Ln;
                                }
                                var c = 0 !== (4 & t),
                                    d = !c && "scroll" === e,
                                    f = c
                                        ? null !== o
                                            ? o + "Capture"
                                            : null
                                        : o;
                                c = [];
                                for (var p, m = r; null !== m; ) {
                                    var h = (p = m).stateNode;
                                    if (
                                        (5 === p.tag &&
                                            null !== h &&
                                            ((p = h),
                                            null !== f &&
                                                null != (h = ze(m, f)) &&
                                                c.push(Kr(m, h, p))),
                                        d)
                                    )
                                        break;
                                    m = m.return;
                                }
                                0 < c.length &&
                                    ((o = new s(o, u, null, n, a)),
                                    i.push({ event: o, listeners: c }));
                            }
                        }
                        if (0 === (7 & t)) {
                            if (
                                ((s = "mouseout" === e || "pointerout" === e),
                                (!(o =
                                    "mouseover" === e || "pointerover" === e) ||
                                    n === xe ||
                                    !(u = n.relatedTarget || n.fromElement) ||
                                    (!ya(u) && !u[ma])) &&
                                    (s || o) &&
                                    ((o =
                                        a.window === a
                                            ? a
                                            : (o = a.ownerDocument)
                                            ? o.defaultView || o.parentWindow
                                            : window),
                                    s
                                        ? ((s = r),
                                          null !==
                                              (u = (u =
                                                  n.relatedTarget ||
                                                  n.toElement)
                                                  ? ya(u)
                                                  : null) &&
                                              (u !== (d = Ve(u)) ||
                                                  (5 !== u.tag &&
                                                      6 !== u.tag)) &&
                                              (u = null))
                                        : ((s = null), (u = r)),
                                    s !== u))
                            ) {
                                if (
                                    ((c = mn),
                                    (h = "onMouseLeave"),
                                    (f = "onMouseEnter"),
                                    (m = "mouse"),
                                    ("pointerout" !== e &&
                                        "pointerover" !== e) ||
                                        ((c = Ln),
                                        (h = "onPointerLeave"),
                                        (f = "onPointerEnter"),
                                        (m = "pointer")),
                                    (d = null == s ? o : xa(s)),
                                    (p = null == u ? o : xa(u)),
                                    ((o = new c(
                                        h,
                                        m + "leave",
                                        s,
                                        n,
                                        a,
                                    )).target = d),
                                    (o.relatedTarget = p),
                                    (h = null),
                                    ya(a) === r &&
                                        (((c = new c(
                                            f,
                                            m + "enter",
                                            u,
                                            n,
                                            a,
                                        )).target = p),
                                        (c.relatedTarget = d),
                                        (h = c)),
                                    (d = h),
                                    s && u)
                                )
                                    e: {
                                        for (
                                            f = u, m = 0, p = c = s;
                                            p;
                                            p = Qr(p)
                                        )
                                            m++;
                                        for (p = 0, h = f; h; h = Qr(h)) p++;
                                        for (; 0 < m - p; ) (c = Qr(c)), m--;
                                        for (; 0 < p - m; ) (f = Qr(f)), p--;
                                        for (; m--; ) {
                                            if (
                                                c === f ||
                                                (null !== f &&
                                                    c === f.alternate)
                                            )
                                                break e;
                                            (c = Qr(c)), (f = Qr(f));
                                        }
                                        c = null;
                                    }
                                else c = null;
                                null !== s && qr(i, o, s, c, !1),
                                    null !== u &&
                                        null !== d &&
                                        qr(i, d, u, c, !0);
                            }
                            if (
                                "select" ===
                                    (s =
                                        (o = r ? xa(r) : window).nodeName &&
                                        o.nodeName.toLowerCase()) ||
                                ("input" === s && "file" === o.type)
                            )
                                var v = Yn;
                            else if (Hn(o))
                                if (Xn) v = ir;
                                else {
                                    v = ar;
                                    var g = rr;
                                }
                            else
                                (s = o.nodeName) &&
                                    "input" === s.toLowerCase() &&
                                    ("checkbox" === o.type ||
                                        "radio" === o.type) &&
                                    (v = lr);
                            switch (
                                (v && (v = v(e, r))
                                    ? Kn(i, v, n, a)
                                    : (g && g(e, o, r),
                                      "focusout" === e &&
                                          (g = o._wrapperState) &&
                                          g.controlled &&
                                          "number" === o.type &&
                                          ee(o, "number", o.value)),
                                (g = r ? xa(r) : window),
                                e)
                            ) {
                                case "focusin":
                                    (Hn(g) || "true" === g.contentEditable) &&
                                        ((vr = g), (gr = r), (yr = null));
                                    break;
                                case "focusout":
                                    yr = gr = vr = null;
                                    break;
                                case "mousedown":
                                    br = !0;
                                    break;
                                case "contextmenu":
                                case "mouseup":
                                case "dragend":
                                    (br = !1), xr(i, n, a);
                                    break;
                                case "selectionchange":
                                    if (hr) break;
                                case "keydown":
                                case "keyup":
                                    xr(i, n, a);
                            }
                            var y;
                            if (Rn)
                                e: {
                                    switch (e) {
                                        case "compositionstart":
                                            var b = "onCompositionStart";
                                            break e;
                                        case "compositionend":
                                            b = "onCompositionEnd";
                                            break e;
                                        case "compositionupdate":
                                            b = "onCompositionUpdate";
                                            break e;
                                    }
                                    b = void 0;
                                }
                            else
                                Vn
                                    ? Bn(e, n) && (b = "onCompositionEnd")
                                    : "keydown" === e &&
                                      229 === n.keyCode &&
                                      (b = "onCompositionStart");
                            b &&
                                (In &&
                                    "ko" !== n.locale &&
                                    (Vn || "onCompositionStart" !== b
                                        ? "onCompositionEnd" === b &&
                                          Vn &&
                                          (y = en())
                                        : ((Zt =
                                              "value" in (Xt = a)
                                                  ? Xt.value
                                                  : Xt.textContent),
                                          (Vn = !0))),
                                0 < (g = Wr(r, b)).length &&
                                    ((b = new xn(b, e, null, n, a)),
                                    i.push({ event: b, listeners: g }),
                                    y
                                        ? (b.data = y)
                                        : null !== (y = An(n)) &&
                                          (b.data = y))),
                                (y = Fn
                                    ? (function (e, t) {
                                          switch (e) {
                                              case "compositionend":
                                                  return An(t);
                                              case "keypress":
                                                  return 32 !== t.which
                                                      ? null
                                                      : ((Un = !0), Dn);
                                              case "textInput":
                                                  return (e = t.data) === Dn &&
                                                      Un
                                                      ? null
                                                      : e;
                                              default:
                                                  return null;
                                          }
                                      })(e, n)
                                    : (function (e, t) {
                                          if (Vn)
                                              return "compositionend" === e ||
                                                  (!Rn && Bn(e, t))
                                                  ? ((e = en()),
                                                    (Jt = Zt = Xt = null),
                                                    (Vn = !1),
                                                    e)
                                                  : null;
                                          switch (e) {
                                              case "paste":
                                              default:
                                                  return null;
                                              case "keypress":
                                                  if (
                                                      !(
                                                          t.ctrlKey ||
                                                          t.altKey ||
                                                          t.metaKey
                                                      ) ||
                                                      (t.ctrlKey && t.altKey)
                                                  ) {
                                                      if (
                                                          t.char &&
                                                          1 < t.char.length
                                                      )
                                                          return t.char;
                                                      if (t.which)
                                                          return String.fromCharCode(
                                                              t.which,
                                                          );
                                                  }
                                                  return null;
                                              case "compositionend":
                                                  return In && "ko" !== t.locale
                                                      ? null
                                                      : t.data;
                                          }
                                      })(e, n)) &&
                                    0 < (r = Wr(r, "onBeforeInput")).length &&
                                    ((a = new xn(
                                        "onBeforeInput",
                                        "beforeinput",
                                        null,
                                        n,
                                        a,
                                    )),
                                    i.push({ event: a, listeners: r }),
                                    (a.data = y));
                        }
                        Dr(i, t);
                    });
                }
                function Kr(e, t, n) {
                    return { instance: e, listener: t, currentTarget: n };
                }
                function Wr(e, t) {
                    for (var n = t + "Capture", r = []; null !== e; ) {
                        var a = e,
                            l = a.stateNode;
                        5 === a.tag &&
                            null !== l &&
                            ((a = l),
                            null != (l = ze(e, n)) && r.unshift(Kr(e, l, a)),
                            null != (l = ze(e, t)) && r.push(Kr(e, l, a))),
                            (e = e.return);
                    }
                    return r;
                }
                function Qr(e) {
                    if (null === e) return null;
                    do {
                        e = e.return;
                    } while (e && 5 !== e.tag);
                    return e || null;
                }
                function qr(e, t, n, r, a) {
                    for (
                        var l = t._reactName, i = [];
                        null !== n && n !== r;

                    ) {
                        var o = n,
                            s = o.alternate,
                            u = o.stateNode;
                        if (null !== s && s === r) break;
                        5 === o.tag &&
                            null !== u &&
                            ((o = u),
                            a
                                ? null != (s = ze(n, l)) &&
                                  i.unshift(Kr(n, s, o))
                                : a ||
                                  (null != (s = ze(n, l)) &&
                                      i.push(Kr(n, s, o)))),
                            (n = n.return);
                    }
                    0 !== i.length && e.push({ event: t, listeners: i });
                }
                var Gr = /\r\n?/g,
                    Yr = /\u0000|\uFFFD/g;
                function Xr(e) {
                    return ("string" === typeof e ? e : "" + e)
                        .replace(Gr, "\n")
                        .replace(Yr, "");
                }
                function Zr(e, t, n) {
                    if (((t = Xr(t)), Xr(e) !== t && n)) throw Error(l(425));
                }
                function Jr() {}
                var ea = null,
                    ta = null;
                function na(e, t) {
                    return (
                        "textarea" === e ||
                        "noscript" === e ||
                        "string" === typeof t.children ||
                        "number" === typeof t.children ||
                        ("object" === typeof t.dangerouslySetInnerHTML &&
                            null !== t.dangerouslySetInnerHTML &&
                            null != t.dangerouslySetInnerHTML.__html)
                    );
                }
                var ra = "function" === typeof setTimeout ? setTimeout : void 0,
                    aa =
                        "function" === typeof clearTimeout
                            ? clearTimeout
                            : void 0,
                    la = "function" === typeof Promise ? Promise : void 0,
                    ia =
                        "function" === typeof queueMicrotask
                            ? queueMicrotask
                            : "undefined" !== typeof la
                            ? function (e) {
                                  return la.resolve(null).then(e).catch(oa);
                              }
                            : ra;
                function oa(e) {
                    setTimeout(function () {
                        throw e;
                    });
                }
                function sa(e, t) {
                    var n = t,
                        r = 0;
                    do {
                        var a = n.nextSibling;
                        if ((e.removeChild(n), a && 8 === a.nodeType))
                            if ("/$" === (n = a.data)) {
                                if (0 === r)
                                    return e.removeChild(a), void Vt(t);
                                r--;
                            } else
                                ("$" !== n && "$?" !== n && "$!" !== n) || r++;
                        n = a;
                    } while (n);
                    Vt(t);
                }
                function ua(e) {
                    for (; null != e; e = e.nextSibling) {
                        var t = e.nodeType;
                        if (1 === t || 3 === t) break;
                        if (8 === t) {
                            if (
                                "$" === (t = e.data) ||
                                "$!" === t ||
                                "$?" === t
                            )
                                break;
                            if ("/$" === t) return null;
                        }
                    }
                    return e;
                }
                function ca(e) {
                    e = e.previousSibling;
                    for (var t = 0; e; ) {
                        if (8 === e.nodeType) {
                            var n = e.data;
                            if ("$" === n || "$!" === n || "$?" === n) {
                                if (0 === t) return e;
                                t--;
                            } else "/$" === n && t++;
                        }
                        e = e.previousSibling;
                    }
                    return null;
                }
                var da = Math.random().toString(36).slice(2),
                    fa = "__reactFiber$" + da,
                    pa = "__reactProps$" + da,
                    ma = "__reactContainer$" + da,
                    ha = "__reactEvents$" + da,
                    va = "__reactListeners$" + da,
                    ga = "__reactHandles$" + da;
                function ya(e) {
                    var t = e[fa];
                    if (t) return t;
                    for (var n = e.parentNode; n; ) {
                        if ((t = n[ma] || n[fa])) {
                            if (
                                ((n = t.alternate),
                                null !== t.child ||
                                    (null !== n && null !== n.child))
                            )
                                for (e = ca(e); null !== e; ) {
                                    if ((n = e[fa])) return n;
                                    e = ca(e);
                                }
                            return t;
                        }
                        n = (e = n).parentNode;
                    }
                    return null;
                }
                function ba(e) {
                    return !(e = e[fa] || e[ma]) ||
                        (5 !== e.tag &&
                            6 !== e.tag &&
                            13 !== e.tag &&
                            3 !== e.tag)
                        ? null
                        : e;
                }
                function xa(e) {
                    if (5 === e.tag || 6 === e.tag) return e.stateNode;
                    throw Error(l(33));
                }
                function wa(e) {
                    return e[pa] || null;
                }
                var ja = [],
                    ka = -1;
                function Sa(e) {
                    return { current: e };
                }
                function Na(e) {
                    0 > ka || ((e.current = ja[ka]), (ja[ka] = null), ka--);
                }
                function Ca(e, t) {
                    ka++, (ja[ka] = e.current), (e.current = t);
                }
                var Ea = {},
                    La = Sa(Ea),
                    _a = Sa(!1),
                    Pa = Ea;
                function Oa(e, t) {
                    var n = e.type.contextTypes;
                    if (!n) return Ea;
                    var r = e.stateNode;
                    if (
                        r &&
                        r.__reactInternalMemoizedUnmaskedChildContext === t
                    )
                        return r.__reactInternalMemoizedMaskedChildContext;
                    var a,
                        l = {};
                    for (a in n) l[a] = t[a];
                    return (
                        r &&
                            (((e =
                                e.stateNode).__reactInternalMemoizedUnmaskedChildContext =
                                t),
                            (e.__reactInternalMemoizedMaskedChildContext = l)),
                        l
                    );
                }
                function za(e) {
                    return null !== (e = e.childContextTypes) && void 0 !== e;
                }
                function Ma() {
                    Na(_a), Na(La);
                }
                function Ra(e, t, n) {
                    if (La.current !== Ea) throw Error(l(168));
                    Ca(La, t), Ca(_a, n);
                }
                function Ta(e, t, n) {
                    var r = e.stateNode;
                    if (
                        ((t = t.childContextTypes),
                        "function" !== typeof r.getChildContext)
                    )
                        return n;
                    for (var a in (r = r.getChildContext()))
                        if (!(a in t))
                            throw Error(l(108, $(e) || "Unknown", a));
                    return I({}, n, r);
                }
                function Fa(e) {
                    return (
                        (e =
                            ((e = e.stateNode) &&
                                e.__reactInternalMemoizedMergedChildContext) ||
                            Ea),
                        (Pa = La.current),
                        Ca(La, e),
                        Ca(_a, _a.current),
                        !0
                    );
                }
                function Ia(e, t, n) {
                    var r = e.stateNode;
                    if (!r) throw Error(l(169));
                    n
                        ? ((e = Ta(e, t, Pa)),
                          (r.__reactInternalMemoizedMergedChildContext = e),
                          Na(_a),
                          Na(La),
                          Ca(La, e))
                        : Na(_a),
                        Ca(_a, n);
                }
                var Da = null,
                    Ua = !1,
                    Ba = !1;
                function Aa(e) {
                    null === Da ? (Da = [e]) : Da.push(e);
                }
                function Va() {
                    if (!Ba && null !== Da) {
                        Ba = !0;
                        var e = 0,
                            t = bt;
                        try {
                            var n = Da;
                            for (bt = 1; e < n.length; e++) {
                                var r = n[e];
                                do {
                                    r = r(!0);
                                } while (null !== r);
                            }
                            (Da = null), (Ua = !1);
                        } catch (a) {
                            throw (
                                (null !== Da && (Da = Da.slice(e + 1)),
                                Qe(Je, Va),
                                a)
                            );
                        } finally {
                            (bt = t), (Ba = !1);
                        }
                    }
                    return null;
                }
                var $a = [],
                    Ha = 0,
                    Ka = null,
                    Wa = 0,
                    Qa = [],
                    qa = 0,
                    Ga = null,
                    Ya = 1,
                    Xa = "";
                function Za(e, t) {
                    ($a[Ha++] = Wa), ($a[Ha++] = Ka), (Ka = e), (Wa = t);
                }
                function Ja(e, t, n) {
                    (Qa[qa++] = Ya), (Qa[qa++] = Xa), (Qa[qa++] = Ga), (Ga = e);
                    var r = Ya;
                    e = Xa;
                    var a = 32 - it(r) - 1;
                    (r &= ~(1 << a)), (n += 1);
                    var l = 32 - it(t) + a;
                    if (30 < l) {
                        var i = a - (a % 5);
                        (l = (r & ((1 << i) - 1)).toString(32)),
                            (r >>= i),
                            (a -= i),
                            (Ya = (1 << (32 - it(t) + a)) | (n << a) | r),
                            (Xa = l + e);
                    } else (Ya = (1 << l) | (n << a) | r), (Xa = e);
                }
                function el(e) {
                    null !== e.return && (Za(e, 1), Ja(e, 1, 0));
                }
                function tl(e) {
                    for (; e === Ka; )
                        (Ka = $a[--Ha]),
                            ($a[Ha] = null),
                            (Wa = $a[--Ha]),
                            ($a[Ha] = null);
                    for (; e === Ga; )
                        (Ga = Qa[--qa]),
                            (Qa[qa] = null),
                            (Xa = Qa[--qa]),
                            (Qa[qa] = null),
                            (Ya = Qa[--qa]),
                            (Qa[qa] = null);
                }
                var nl = null,
                    rl = null,
                    al = !1,
                    ll = null;
                function il(e, t) {
                    var n = zu(5, null, null, 0);
                    (n.elementType = "DELETED"),
                        (n.stateNode = t),
                        (n.return = e),
                        null === (t = e.deletions)
                            ? ((e.deletions = [n]), (e.flags |= 16))
                            : t.push(n);
                }
                function ol(e, t) {
                    switch (e.tag) {
                        case 5:
                            var n = e.type;
                            return (
                                null !==
                                    (t =
                                        1 !== t.nodeType ||
                                        n.toLowerCase() !==
                                            t.nodeName.toLowerCase()
                                            ? null
                                            : t) &&
                                ((e.stateNode = t),
                                (nl = e),
                                (rl = ua(t.firstChild)),
                                !0)
                            );
                        case 6:
                            return (
                                null !==
                                    (t =
                                        "" === e.pendingProps ||
                                        3 !== t.nodeType
                                            ? null
                                            : t) &&
                                ((e.stateNode = t), (nl = e), (rl = null), !0)
                            );
                        case 13:
                            return (
                                null !== (t = 8 !== t.nodeType ? null : t) &&
                                ((n =
                                    null !== Ga
                                        ? { id: Ya, overflow: Xa }
                                        : null),
                                (e.memoizedState = {
                                    dehydrated: t,
                                    treeContext: n,
                                    retryLane: 1073741824,
                                }),
                                ((n = zu(18, null, null, 0)).stateNode = t),
                                (n.return = e),
                                (e.child = n),
                                (nl = e),
                                (rl = null),
                                !0)
                            );
                        default:
                            return !1;
                    }
                }
                function sl(e) {
                    return 0 !== (1 & e.mode) && 0 === (128 & e.flags);
                }
                function ul(e) {
                    if (al) {
                        var t = rl;
                        if (t) {
                            var n = t;
                            if (!ol(e, t)) {
                                if (sl(e)) throw Error(l(418));
                                t = ua(n.nextSibling);
                                var r = nl;
                                t && ol(e, t)
                                    ? il(r, n)
                                    : ((e.flags = (-4097 & e.flags) | 2),
                                      (al = !1),
                                      (nl = e));
                            }
                        } else {
                            if (sl(e)) throw Error(l(418));
                            (e.flags = (-4097 & e.flags) | 2),
                                (al = !1),
                                (nl = e);
                        }
                    }
                }
                function cl(e) {
                    for (
                        e = e.return;
                        null !== e &&
                        5 !== e.tag &&
                        3 !== e.tag &&
                        13 !== e.tag;

                    )
                        e = e.return;
                    nl = e;
                }
                function dl(e) {
                    if (e !== nl) return !1;
                    if (!al) return cl(e), (al = !0), !1;
                    var t;
                    if (
                        ((t = 3 !== e.tag) &&
                            !(t = 5 !== e.tag) &&
                            (t =
                                "head" !== (t = e.type) &&
                                "body" !== t &&
                                !na(e.type, e.memoizedProps)),
                        t && (t = rl))
                    ) {
                        if (sl(e)) throw (fl(), Error(l(418)));
                        for (; t; ) il(e, t), (t = ua(t.nextSibling));
                    }
                    if ((cl(e), 13 === e.tag)) {
                        if (
                            !(e =
                                null !== (e = e.memoizedState)
                                    ? e.dehydrated
                                    : null)
                        )
                            throw Error(l(317));
                        e: {
                            for (e = e.nextSibling, t = 0; e; ) {
                                if (8 === e.nodeType) {
                                    var n = e.data;
                                    if ("/$" === n) {
                                        if (0 === t) {
                                            rl = ua(e.nextSibling);
                                            break e;
                                        }
                                        t--;
                                    } else
                                        ("$" !== n &&
                                            "$!" !== n &&
                                            "$?" !== n) ||
                                            t++;
                                }
                                e = e.nextSibling;
                            }
                            rl = null;
                        }
                    } else rl = nl ? ua(e.stateNode.nextSibling) : null;
                    return !0;
                }
                function fl() {
                    for (var e = rl; e; ) e = ua(e.nextSibling);
                }
                function pl() {
                    (rl = nl = null), (al = !1);
                }
                function ml(e) {
                    null === ll ? (ll = [e]) : ll.push(e);
                }
                var hl = x.ReactCurrentBatchConfig;
                function vl(e, t) {
                    if (e && e.defaultProps) {
                        for (var n in ((t = I({}, t)), (e = e.defaultProps)))
                            void 0 === t[n] && (t[n] = e[n]);
                        return t;
                    }
                    return t;
                }
                var gl = Sa(null),
                    yl = null,
                    bl = null,
                    xl = null;
                function wl() {
                    xl = bl = yl = null;
                }
                function jl(e) {
                    var t = gl.current;
                    Na(gl), (e._currentValue = t);
                }
                function kl(e, t, n) {
                    for (; null !== e; ) {
                        var r = e.alternate;
                        if (
                            ((e.childLanes & t) !== t
                                ? ((e.childLanes |= t),
                                  null !== r && (r.childLanes |= t))
                                : null !== r &&
                                  (r.childLanes & t) !== t &&
                                  (r.childLanes |= t),
                            e === n)
                        )
                            break;
                        e = e.return;
                    }
                }
                function Sl(e, t) {
                    (yl = e),
                        (xl = bl = null),
                        null !== (e = e.dependencies) &&
                            null !== e.firstContext &&
                            (0 !== (e.lanes & t) && (xo = !0),
                            (e.firstContext = null));
                }
                function Nl(e) {
                    var t = e._currentValue;
                    if (xl !== e)
                        if (
                            ((e = { context: e, memoizedValue: t, next: null }),
                            null === bl)
                        ) {
                            if (null === yl) throw Error(l(308));
                            (bl = e),
                                (yl.dependencies = {
                                    lanes: 0,
                                    firstContext: e,
                                });
                        } else bl = bl.next = e;
                    return t;
                }
                var Cl = null;
                function El(e) {
                    null === Cl ? (Cl = [e]) : Cl.push(e);
                }
                function Ll(e, t, n, r) {
                    var a = t.interleaved;
                    return (
                        null === a
                            ? ((n.next = n), El(t))
                            : ((n.next = a.next), (a.next = n)),
                        (t.interleaved = n),
                        _l(e, r)
                    );
                }
                function _l(e, t) {
                    e.lanes |= t;
                    var n = e.alternate;
                    for (
                        null !== n && (n.lanes |= t), n = e, e = e.return;
                        null !== e;

                    )
                        (e.childLanes |= t),
                            null !== (n = e.alternate) && (n.childLanes |= t),
                            (n = e),
                            (e = e.return);
                    return 3 === n.tag ? n.stateNode : null;
                }
                var Pl = !1;
                function Ol(e) {
                    e.updateQueue = {
                        baseState: e.memoizedState,
                        firstBaseUpdate: null,
                        lastBaseUpdate: null,
                        shared: { pending: null, interleaved: null, lanes: 0 },
                        effects: null,
                    };
                }
                function zl(e, t) {
                    (e = e.updateQueue),
                        t.updateQueue === e &&
                            (t.updateQueue = {
                                baseState: e.baseState,
                                firstBaseUpdate: e.firstBaseUpdate,
                                lastBaseUpdate: e.lastBaseUpdate,
                                shared: e.shared,
                                effects: e.effects,
                            });
                }
                function Ml(e, t) {
                    return {
                        eventTime: e,
                        lane: t,
                        tag: 0,
                        payload: null,
                        callback: null,
                        next: null,
                    };
                }
                function Rl(e, t, n) {
                    var r = e.updateQueue;
                    if (null === r) return null;
                    if (((r = r.shared), 0 !== (2 & _s))) {
                        var a = r.pending;
                        return (
                            null === a
                                ? (t.next = t)
                                : ((t.next = a.next), (a.next = t)),
                            (r.pending = t),
                            _l(e, n)
                        );
                    }
                    return (
                        null === (a = r.interleaved)
                            ? ((t.next = t), El(r))
                            : ((t.next = a.next), (a.next = t)),
                        (r.interleaved = t),
                        _l(e, n)
                    );
                }
                function Tl(e, t, n) {
                    if (
                        null !== (t = t.updateQueue) &&
                        ((t = t.shared), 0 !== (4194240 & n))
                    ) {
                        var r = t.lanes;
                        (n |= r &= e.pendingLanes), (t.lanes = n), yt(e, n);
                    }
                }
                function Fl(e, t) {
                    var n = e.updateQueue,
                        r = e.alternate;
                    if (null !== r && n === (r = r.updateQueue)) {
                        var a = null,
                            l = null;
                        if (null !== (n = n.firstBaseUpdate)) {
                            do {
                                var i = {
                                    eventTime: n.eventTime,
                                    lane: n.lane,
                                    tag: n.tag,
                                    payload: n.payload,
                                    callback: n.callback,
                                    next: null,
                                };
                                null === l ? (a = l = i) : (l = l.next = i),
                                    (n = n.next);
                            } while (null !== n);
                            null === l ? (a = l = t) : (l = l.next = t);
                        } else a = l = t;
                        return (
                            (n = {
                                baseState: r.baseState,
                                firstBaseUpdate: a,
                                lastBaseUpdate: l,
                                shared: r.shared,
                                effects: r.effects,
                            }),
                            void (e.updateQueue = n)
                        );
                    }
                    null === (e = n.lastBaseUpdate)
                        ? (n.firstBaseUpdate = t)
                        : (e.next = t),
                        (n.lastBaseUpdate = t);
                }
                function Il(e, t, n, r) {
                    var a = e.updateQueue;
                    Pl = !1;
                    var l = a.firstBaseUpdate,
                        i = a.lastBaseUpdate,
                        o = a.shared.pending;
                    if (null !== o) {
                        a.shared.pending = null;
                        var s = o,
                            u = s.next;
                        (s.next = null),
                            null === i ? (l = u) : (i.next = u),
                            (i = s);
                        var c = e.alternate;
                        null !== c &&
                            (o = (c = c.updateQueue).lastBaseUpdate) !== i &&
                            (null === o
                                ? (c.firstBaseUpdate = u)
                                : (o.next = u),
                            (c.lastBaseUpdate = s));
                    }
                    if (null !== l) {
                        var d = a.baseState;
                        for (i = 0, c = u = s = null, o = l; ; ) {
                            var f = o.lane,
                                p = o.eventTime;
                            if ((r & f) === f) {
                                null !== c &&
                                    (c = c.next =
                                        {
                                            eventTime: p,
                                            lane: 0,
                                            tag: o.tag,
                                            payload: o.payload,
                                            callback: o.callback,
                                            next: null,
                                        });
                                e: {
                                    var m = e,
                                        h = o;
                                    switch (((f = t), (p = n), h.tag)) {
                                        case 1:
                                            if (
                                                "function" ===
                                                typeof (m = h.payload)
                                            ) {
                                                d = m.call(p, d, f);
                                                break e;
                                            }
                                            d = m;
                                            break e;
                                        case 3:
                                            m.flags = (-65537 & m.flags) | 128;
                                        case 0:
                                            if (
                                                null ===
                                                    (f =
                                                        "function" ===
                                                        typeof (m = h.payload)
                                                            ? m.call(p, d, f)
                                                            : m) ||
                                                void 0 === f
                                            )
                                                break e;
                                            d = I({}, d, f);
                                            break e;
                                        case 2:
                                            Pl = !0;
                                    }
                                }
                                null !== o.callback &&
                                    0 !== o.lane &&
                                    ((e.flags |= 64),
                                    null === (f = a.effects)
                                        ? (a.effects = [o])
                                        : f.push(o));
                            } else
                                (p = {
                                    eventTime: p,
                                    lane: f,
                                    tag: o.tag,
                                    payload: o.payload,
                                    callback: o.callback,
                                    next: null,
                                }),
                                    null === c
                                        ? ((u = c = p), (s = d))
                                        : (c = c.next = p),
                                    (i |= f);
                            if (null === (o = o.next)) {
                                if (null === (o = a.shared.pending)) break;
                                (o = (f = o).next),
                                    (f.next = null),
                                    (a.lastBaseUpdate = f),
                                    (a.shared.pending = null);
                            }
                        }
                        if (
                            (null === c && (s = d),
                            (a.baseState = s),
                            (a.firstBaseUpdate = u),
                            (a.lastBaseUpdate = c),
                            null !== (t = a.shared.interleaved))
                        ) {
                            a = t;
                            do {
                                (i |= a.lane), (a = a.next);
                            } while (a !== t);
                        } else null === l && (a.shared.lanes = 0);
                        (Is |= i), (e.lanes = i), (e.memoizedState = d);
                    }
                }
                function Dl(e, t, n) {
                    if (((e = t.effects), (t.effects = null), null !== e))
                        for (t = 0; t < e.length; t++) {
                            var r = e[t],
                                a = r.callback;
                            if (null !== a) {
                                if (
                                    ((r.callback = null),
                                    (r = n),
                                    "function" !== typeof a)
                                )
                                    throw Error(l(191, a));
                                a.call(r);
                            }
                        }
                }
                var Ul = new r.Component().refs;
                function Bl(e, t, n, r) {
                    (n =
                        null === (n = n(r, (t = e.memoizedState))) ||
                        void 0 === n
                            ? t
                            : I({}, t, n)),
                        (e.memoizedState = n),
                        0 === e.lanes && (e.updateQueue.baseState = n);
                }
                var Al = {
                    isMounted: function (e) {
                        return !!(e = e._reactInternals) && Ve(e) === e;
                    },
                    enqueueSetState: function (e, t, n) {
                        e = e._reactInternals;
                        var r = tu(),
                            a = nu(e),
                            l = Ml(r, a);
                        (l.payload = t),
                            void 0 !== n && null !== n && (l.callback = n),
                            null !== (t = Rl(e, l, a)) &&
                                (ru(t, e, a, r), Tl(t, e, a));
                    },
                    enqueueReplaceState: function (e, t, n) {
                        e = e._reactInternals;
                        var r = tu(),
                            a = nu(e),
                            l = Ml(r, a);
                        (l.tag = 1),
                            (l.payload = t),
                            void 0 !== n && null !== n && (l.callback = n),
                            null !== (t = Rl(e, l, a)) &&
                                (ru(t, e, a, r), Tl(t, e, a));
                    },
                    enqueueForceUpdate: function (e, t) {
                        e = e._reactInternals;
                        var n = tu(),
                            r = nu(e),
                            a = Ml(n, r);
                        (a.tag = 2),
                            void 0 !== t && null !== t && (a.callback = t),
                            null !== (t = Rl(e, a, r)) &&
                                (ru(t, e, r, n), Tl(t, e, r));
                    },
                };
                function Vl(e, t, n, r, a, l, i) {
                    return "function" ===
                        typeof (e = e.stateNode).shouldComponentUpdate
                        ? e.shouldComponentUpdate(r, l, i)
                        : !t.prototype ||
                              !t.prototype.isPureReactComponent ||
                              !sr(n, r) ||
                              !sr(a, l);
                }
                function $l(e, t, n) {
                    var r = !1,
                        a = Ea,
                        l = t.contextType;
                    return (
                        "object" === typeof l && null !== l
                            ? (l = Nl(l))
                            : ((a = za(t) ? Pa : La.current),
                              (l = (r =
                                  null !== (r = t.contextTypes) && void 0 !== r)
                                  ? Oa(e, a)
                                  : Ea)),
                        (t = new t(n, l)),
                        (e.memoizedState =
                            null !== t.state && void 0 !== t.state
                                ? t.state
                                : null),
                        (t.updater = Al),
                        (e.stateNode = t),
                        (t._reactInternals = e),
                        r &&
                            (((e =
                                e.stateNode).__reactInternalMemoizedUnmaskedChildContext =
                                a),
                            (e.__reactInternalMemoizedMaskedChildContext = l)),
                        t
                    );
                }
                function Hl(e, t, n, r) {
                    (e = t.state),
                        "function" === typeof t.componentWillReceiveProps &&
                            t.componentWillReceiveProps(n, r),
                        "function" ===
                            typeof t.UNSAFE_componentWillReceiveProps &&
                            t.UNSAFE_componentWillReceiveProps(n, r),
                        t.state !== e &&
                            Al.enqueueReplaceState(t, t.state, null);
                }
                function Kl(e, t, n, r) {
                    var a = e.stateNode;
                    (a.props = n),
                        (a.state = e.memoizedState),
                        (a.refs = Ul),
                        Ol(e);
                    var l = t.contextType;
                    "object" === typeof l && null !== l
                        ? (a.context = Nl(l))
                        : ((l = za(t) ? Pa : La.current),
                          (a.context = Oa(e, l))),
                        (a.state = e.memoizedState),
                        "function" ===
                            typeof (l = t.getDerivedStateFromProps) &&
                            (Bl(e, t, l, n), (a.state = e.memoizedState)),
                        "function" === typeof t.getDerivedStateFromProps ||
                            "function" === typeof a.getSnapshotBeforeUpdate ||
                            ("function" !==
                                typeof a.UNSAFE_componentWillMount &&
                                "function" !== typeof a.componentWillMount) ||
                            ((t = a.state),
                            "function" === typeof a.componentWillMount &&
                                a.componentWillMount(),
                            "function" === typeof a.UNSAFE_componentWillMount &&
                                a.UNSAFE_componentWillMount(),
                            t !== a.state &&
                                Al.enqueueReplaceState(a, a.state, null),
                            Il(e, n, a, r),
                            (a.state = e.memoizedState)),
                        "function" === typeof a.componentDidMount &&
                            (e.flags |= 4194308);
                }
                function Wl(e, t, n) {
                    if (
                        null !== (e = n.ref) &&
                        "function" !== typeof e &&
                        "object" !== typeof e
                    ) {
                        if (n._owner) {
                            if ((n = n._owner)) {
                                if (1 !== n.tag) throw Error(l(309));
                                var r = n.stateNode;
                            }
                            if (!r) throw Error(l(147, e));
                            var a = r,
                                i = "" + e;
                            return null !== t &&
                                null !== t.ref &&
                                "function" === typeof t.ref &&
                                t.ref._stringRef === i
                                ? t.ref
                                : ((t = function (e) {
                                      var t = a.refs;
                                      t === Ul && (t = a.refs = {}),
                                          null === e ? delete t[i] : (t[i] = e);
                                  }),
                                  (t._stringRef = i),
                                  t);
                        }
                        if ("string" !== typeof e) throw Error(l(284));
                        if (!n._owner) throw Error(l(290, e));
                    }
                    return e;
                }
                function Ql(e, t) {
                    throw (
                        ((e = Object.prototype.toString.call(t)),
                        Error(
                            l(
                                31,
                                "[object Object]" === e
                                    ? "object with keys {" +
                                          Object.keys(t).join(", ") +
                                          "}"
                                    : e,
                            ),
                        ))
                    );
                }
                function ql(e) {
                    return (0, e._init)(e._payload);
                }
                function Gl(e) {
                    function t(t, n) {
                        if (e) {
                            var r = t.deletions;
                            null === r
                                ? ((t.deletions = [n]), (t.flags |= 16))
                                : r.push(n);
                        }
                    }
                    function n(n, r) {
                        if (!e) return null;
                        for (; null !== r; ) t(n, r), (r = r.sibling);
                        return null;
                    }
                    function r(e, t) {
                        for (e = new Map(); null !== t; )
                            null !== t.key
                                ? e.set(t.key, t)
                                : e.set(t.index, t),
                                (t = t.sibling);
                        return e;
                    }
                    function a(e, t) {
                        return (
                            ((e = Ru(e, t)).index = 0), (e.sibling = null), e
                        );
                    }
                    function i(t, n, r) {
                        return (
                            (t.index = r),
                            e
                                ? null !== (r = t.alternate)
                                    ? (r = r.index) < n
                                        ? ((t.flags |= 2), n)
                                        : r
                                    : ((t.flags |= 2), n)
                                : ((t.flags |= 1048576), n)
                        );
                    }
                    function o(t) {
                        return e && null === t.alternate && (t.flags |= 2), t;
                    }
                    function s(e, t, n, r) {
                        return null === t || 6 !== t.tag
                            ? (((t = Du(n, e.mode, r)).return = e), t)
                            : (((t = a(t, n)).return = e), t);
                    }
                    function u(e, t, n, r) {
                        var l = n.type;
                        return l === k
                            ? d(e, t, n.props.children, r, n.key)
                            : null !== t &&
                              (t.elementType === l ||
                                  ("object" === typeof l &&
                                      null !== l &&
                                      l.$$typeof === z &&
                                      ql(l) === t.type))
                            ? (((r = a(t, n.props)).ref = Wl(e, t, n)),
                              (r.return = e),
                              r)
                            : (((r = Tu(
                                  n.type,
                                  n.key,
                                  n.props,
                                  null,
                                  e.mode,
                                  r,
                              )).ref = Wl(e, t, n)),
                              (r.return = e),
                              r);
                    }
                    function c(e, t, n, r) {
                        return null === t ||
                            4 !== t.tag ||
                            t.stateNode.containerInfo !== n.containerInfo ||
                            t.stateNode.implementation !== n.implementation
                            ? (((t = Uu(n, e.mode, r)).return = e), t)
                            : (((t = a(t, n.children || [])).return = e), t);
                    }
                    function d(e, t, n, r, l) {
                        return null === t || 7 !== t.tag
                            ? (((t = Fu(n, e.mode, r, l)).return = e), t)
                            : (((t = a(t, n)).return = e), t);
                    }
                    function f(e, t, n) {
                        if (
                            ("string" === typeof t && "" !== t) ||
                            "number" === typeof t
                        )
                            return ((t = Du("" + t, e.mode, n)).return = e), t;
                        if ("object" === typeof t && null !== t) {
                            switch (t.$$typeof) {
                                case w:
                                    return (
                                        ((n = Tu(
                                            t.type,
                                            t.key,
                                            t.props,
                                            null,
                                            e.mode,
                                            n,
                                        )).ref = Wl(e, null, t)),
                                        (n.return = e),
                                        n
                                    );
                                case j:
                                    return (
                                        ((t = Uu(t, e.mode, n)).return = e), t
                                    );
                                case z:
                                    return f(e, (0, t._init)(t._payload), n);
                            }
                            if (te(t) || T(t))
                                return (
                                    ((t = Fu(t, e.mode, n, null)).return = e), t
                                );
                            Ql(e, t);
                        }
                        return null;
                    }
                    function p(e, t, n, r) {
                        var a = null !== t ? t.key : null;
                        if (
                            ("string" === typeof n && "" !== n) ||
                            "number" === typeof n
                        )
                            return null !== a ? null : s(e, t, "" + n, r);
                        if ("object" === typeof n && null !== n) {
                            switch (n.$$typeof) {
                                case w:
                                    return n.key === a ? u(e, t, n, r) : null;
                                case j:
                                    return n.key === a ? c(e, t, n, r) : null;
                                case z:
                                    return p(
                                        e,
                                        t,
                                        (a = n._init)(n._payload),
                                        r,
                                    );
                            }
                            if (te(n) || T(n))
                                return null !== a ? null : d(e, t, n, r, null);
                            Ql(e, n);
                        }
                        return null;
                    }
                    function m(e, t, n, r, a) {
                        if (
                            ("string" === typeof r && "" !== r) ||
                            "number" === typeof r
                        )
                            return s(t, (e = e.get(n) || null), "" + r, a);
                        if ("object" === typeof r && null !== r) {
                            switch (r.$$typeof) {
                                case w:
                                    return u(
                                        t,
                                        (e =
                                            e.get(null === r.key ? n : r.key) ||
                                            null),
                                        r,
                                        a,
                                    );
                                case j:
                                    return c(
                                        t,
                                        (e =
                                            e.get(null === r.key ? n : r.key) ||
                                            null),
                                        r,
                                        a,
                                    );
                                case z:
                                    return m(
                                        e,
                                        t,
                                        n,
                                        (0, r._init)(r._payload),
                                        a,
                                    );
                            }
                            if (te(r) || T(r))
                                return d(t, (e = e.get(n) || null), r, a, null);
                            Ql(t, r);
                        }
                        return null;
                    }
                    function h(a, l, o, s) {
                        for (
                            var u = null,
                                c = null,
                                d = l,
                                h = (l = 0),
                                v = null;
                            null !== d && h < o.length;
                            h++
                        ) {
                            d.index > h
                                ? ((v = d), (d = null))
                                : (v = d.sibling);
                            var g = p(a, d, o[h], s);
                            if (null === g) {
                                null === d && (d = v);
                                break;
                            }
                            e && d && null === g.alternate && t(a, d),
                                (l = i(g, l, h)),
                                null === c ? (u = g) : (c.sibling = g),
                                (c = g),
                                (d = v);
                        }
                        if (h === o.length) return n(a, d), al && Za(a, h), u;
                        if (null === d) {
                            for (; h < o.length; h++)
                                null !== (d = f(a, o[h], s)) &&
                                    ((l = i(d, l, h)),
                                    null === c ? (u = d) : (c.sibling = d),
                                    (c = d));
                            return al && Za(a, h), u;
                        }
                        for (d = r(a, d); h < o.length; h++)
                            null !== (v = m(d, a, h, o[h], s)) &&
                                (e &&
                                    null !== v.alternate &&
                                    d.delete(null === v.key ? h : v.key),
                                (l = i(v, l, h)),
                                null === c ? (u = v) : (c.sibling = v),
                                (c = v));
                        return (
                            e &&
                                d.forEach(function (e) {
                                    return t(a, e);
                                }),
                            al && Za(a, h),
                            u
                        );
                    }
                    function v(a, o, s, u) {
                        var c = T(s);
                        if ("function" !== typeof c) throw Error(l(150));
                        if (null == (s = c.call(s))) throw Error(l(151));
                        for (
                            var d = (c = null),
                                h = o,
                                v = (o = 0),
                                g = null,
                                y = s.next();
                            null !== h && !y.done;
                            v++, y = s.next()
                        ) {
                            h.index > v
                                ? ((g = h), (h = null))
                                : (g = h.sibling);
                            var b = p(a, h, y.value, u);
                            if (null === b) {
                                null === h && (h = g);
                                break;
                            }
                            e && h && null === b.alternate && t(a, h),
                                (o = i(b, o, v)),
                                null === d ? (c = b) : (d.sibling = b),
                                (d = b),
                                (h = g);
                        }
                        if (y.done) return n(a, h), al && Za(a, v), c;
                        if (null === h) {
                            for (; !y.done; v++, y = s.next())
                                null !== (y = f(a, y.value, u)) &&
                                    ((o = i(y, o, v)),
                                    null === d ? (c = y) : (d.sibling = y),
                                    (d = y));
                            return al && Za(a, v), c;
                        }
                        for (h = r(a, h); !y.done; v++, y = s.next())
                            null !== (y = m(h, a, v, y.value, u)) &&
                                (e &&
                                    null !== y.alternate &&
                                    h.delete(null === y.key ? v : y.key),
                                (o = i(y, o, v)),
                                null === d ? (c = y) : (d.sibling = y),
                                (d = y));
                        return (
                            e &&
                                h.forEach(function (e) {
                                    return t(a, e);
                                }),
                            al && Za(a, v),
                            c
                        );
                    }
                    return function e(r, l, i, s) {
                        if (
                            ("object" === typeof i &&
                                null !== i &&
                                i.type === k &&
                                null === i.key &&
                                (i = i.props.children),
                            "object" === typeof i && null !== i)
                        ) {
                            switch (i.$$typeof) {
                                case w:
                                    e: {
                                        for (
                                            var u = i.key, c = l;
                                            null !== c;

                                        ) {
                                            if (c.key === u) {
                                                if ((u = i.type) === k) {
                                                    if (7 === c.tag) {
                                                        n(r, c.sibling),
                                                            ((l = a(
                                                                c,
                                                                i.props
                                                                    .children,
                                                            )).return = r),
                                                            (r = l);
                                                        break e;
                                                    }
                                                } else if (
                                                    c.elementType === u ||
                                                    ("object" === typeof u &&
                                                        null !== u &&
                                                        u.$$typeof === z &&
                                                        ql(u) === c.type)
                                                ) {
                                                    n(r, c.sibling),
                                                        ((l = a(
                                                            c,
                                                            i.props,
                                                        )).ref = Wl(r, c, i)),
                                                        (l.return = r),
                                                        (r = l);
                                                    break e;
                                                }
                                                n(r, c);
                                                break;
                                            }
                                            t(r, c), (c = c.sibling);
                                        }
                                        i.type === k
                                            ? (((l = Fu(
                                                  i.props.children,
                                                  r.mode,
                                                  s,
                                                  i.key,
                                              )).return = r),
                                              (r = l))
                                            : (((s = Tu(
                                                  i.type,
                                                  i.key,
                                                  i.props,
                                                  null,
                                                  r.mode,
                                                  s,
                                              )).ref = Wl(r, l, i)),
                                              (s.return = r),
                                              (r = s));
                                    }
                                    return o(r);
                                case j:
                                    e: {
                                        for (c = i.key; null !== l; ) {
                                            if (l.key === c) {
                                                if (
                                                    4 === l.tag &&
                                                    l.stateNode
                                                        .containerInfo ===
                                                        i.containerInfo &&
                                                    l.stateNode
                                                        .implementation ===
                                                        i.implementation
                                                ) {
                                                    n(r, l.sibling),
                                                        ((l = a(
                                                            l,
                                                            i.children || [],
                                                        )).return = r),
                                                        (r = l);
                                                    break e;
                                                }
                                                n(r, l);
                                                break;
                                            }
                                            t(r, l), (l = l.sibling);
                                        }
                                        ((l = Uu(i, r.mode, s)).return = r),
                                            (r = l);
                                    }
                                    return o(r);
                                case z:
                                    return e(
                                        r,
                                        l,
                                        (c = i._init)(i._payload),
                                        s,
                                    );
                            }
                            if (te(i)) return h(r, l, i, s);
                            if (T(i)) return v(r, l, i, s);
                            Ql(r, i);
                        }
                        return ("string" === typeof i && "" !== i) ||
                            "number" === typeof i
                            ? ((i = "" + i),
                              null !== l && 6 === l.tag
                                  ? (n(r, l.sibling),
                                    ((l = a(l, i)).return = r),
                                    (r = l))
                                  : (n(r, l),
                                    ((l = Du(i, r.mode, s)).return = r),
                                    (r = l)),
                              o(r))
                            : n(r, l);
                    };
                }
                var Yl = Gl(!0),
                    Xl = Gl(!1),
                    Zl = {},
                    Jl = Sa(Zl),
                    ei = Sa(Zl),
                    ti = Sa(Zl);
                function ni(e) {
                    if (e === Zl) throw Error(l(174));
                    return e;
                }
                function ri(e, t) {
                    switch (
                        (Ca(ti, t), Ca(ei, e), Ca(Jl, Zl), (e = t.nodeType))
                    ) {
                        case 9:
                        case 11:
                            t = (t = t.documentElement)
                                ? t.namespaceURI
                                : se(null, "");
                            break;
                        default:
                            t = se(
                                (t =
                                    (e = 8 === e ? t.parentNode : t)
                                        .namespaceURI || null),
                                (e = e.tagName),
                            );
                    }
                    Na(Jl), Ca(Jl, t);
                }
                function ai() {
                    Na(Jl), Na(ei), Na(ti);
                }
                function li(e) {
                    ni(ti.current);
                    var t = ni(Jl.current),
                        n = se(t, e.type);
                    t !== n && (Ca(ei, e), Ca(Jl, n));
                }
                function ii(e) {
                    ei.current === e && (Na(Jl), Na(ei));
                }
                var oi = Sa(0);
                function si(e) {
                    for (var t = e; null !== t; ) {
                        if (13 === t.tag) {
                            var n = t.memoizedState;
                            if (
                                null !== n &&
                                (null === (n = n.dehydrated) ||
                                    "$?" === n.data ||
                                    "$!" === n.data)
                            )
                                return t;
                        } else if (
                            19 === t.tag &&
                            void 0 !== t.memoizedProps.revealOrder
                        ) {
                            if (0 !== (128 & t.flags)) return t;
                        } else if (null !== t.child) {
                            (t.child.return = t), (t = t.child);
                            continue;
                        }
                        if (t === e) break;
                        for (; null === t.sibling; ) {
                            if (null === t.return || t.return === e)
                                return null;
                            t = t.return;
                        }
                        (t.sibling.return = t.return), (t = t.sibling);
                    }
                    return null;
                }
                var ui = [];
                function ci() {
                    for (var e = 0; e < ui.length; e++)
                        ui[e]._workInProgressVersionPrimary = null;
                    ui.length = 0;
                }
                var di = x.ReactCurrentDispatcher,
                    fi = x.ReactCurrentBatchConfig,
                    pi = 0,
                    mi = null,
                    hi = null,
                    vi = null,
                    gi = !1,
                    yi = !1,
                    bi = 0,
                    xi = 0;
                function wi() {
                    throw Error(l(321));
                }
                function ji(e, t) {
                    if (null === t) return !1;
                    for (var n = 0; n < t.length && n < e.length; n++)
                        if (!or(e[n], t[n])) return !1;
                    return !0;
                }
                function ki(e, t, n, r, a, i) {
                    if (
                        ((pi = i),
                        (mi = t),
                        (t.memoizedState = null),
                        (t.updateQueue = null),
                        (t.lanes = 0),
                        (di.current =
                            null === e || null === e.memoizedState ? io : oo),
                        (e = n(r, a)),
                        yi)
                    ) {
                        i = 0;
                        do {
                            if (((yi = !1), (bi = 0), 25 <= i))
                                throw Error(l(301));
                            (i += 1),
                                (vi = hi = null),
                                (t.updateQueue = null),
                                (di.current = so),
                                (e = n(r, a));
                        } while (yi);
                    }
                    if (
                        ((di.current = lo),
                        (t = null !== hi && null !== hi.next),
                        (pi = 0),
                        (vi = hi = mi = null),
                        (gi = !1),
                        t)
                    )
                        throw Error(l(300));
                    return e;
                }
                function Si() {
                    var e = 0 !== bi;
                    return (bi = 0), e;
                }
                function Ni() {
                    var e = {
                        memoizedState: null,
                        baseState: null,
                        baseQueue: null,
                        queue: null,
                        next: null,
                    };
                    return (
                        null === vi
                            ? (mi.memoizedState = vi = e)
                            : (vi = vi.next = e),
                        vi
                    );
                }
                function Ci() {
                    if (null === hi) {
                        var e = mi.alternate;
                        e = null !== e ? e.memoizedState : null;
                    } else e = hi.next;
                    var t = null === vi ? mi.memoizedState : vi.next;
                    if (null !== t) (vi = t), (hi = e);
                    else {
                        if (null === e) throw Error(l(310));
                        (e = {
                            memoizedState: (hi = e).memoizedState,
                            baseState: hi.baseState,
                            baseQueue: hi.baseQueue,
                            queue: hi.queue,
                            next: null,
                        }),
                            null === vi
                                ? (mi.memoizedState = vi = e)
                                : (vi = vi.next = e);
                    }
                    return vi;
                }
                function Ei(e, t) {
                    return "function" === typeof t ? t(e) : t;
                }
                function Li(e) {
                    var t = Ci(),
                        n = t.queue;
                    if (null === n) throw Error(l(311));
                    n.lastRenderedReducer = e;
                    var r = hi,
                        a = r.baseQueue,
                        i = n.pending;
                    if (null !== i) {
                        if (null !== a) {
                            var o = a.next;
                            (a.next = i.next), (i.next = o);
                        }
                        (r.baseQueue = a = i), (n.pending = null);
                    }
                    if (null !== a) {
                        (i = a.next), (r = r.baseState);
                        var s = (o = null),
                            u = null,
                            c = i;
                        do {
                            var d = c.lane;
                            if ((pi & d) === d)
                                null !== u &&
                                    (u = u.next =
                                        {
                                            lane: 0,
                                            action: c.action,
                                            hasEagerState: c.hasEagerState,
                                            eagerState: c.eagerState,
                                            next: null,
                                        }),
                                    (r = c.hasEagerState
                                        ? c.eagerState
                                        : e(r, c.action));
                            else {
                                var f = {
                                    lane: d,
                                    action: c.action,
                                    hasEagerState: c.hasEagerState,
                                    eagerState: c.eagerState,
                                    next: null,
                                };
                                null === u
                                    ? ((s = u = f), (o = r))
                                    : (u = u.next = f),
                                    (mi.lanes |= d),
                                    (Is |= d);
                            }
                            c = c.next;
                        } while (null !== c && c !== i);
                        null === u ? (o = r) : (u.next = s),
                            or(r, t.memoizedState) || (xo = !0),
                            (t.memoizedState = r),
                            (t.baseState = o),
                            (t.baseQueue = u),
                            (n.lastRenderedState = r);
                    }
                    if (null !== (e = n.interleaved)) {
                        a = e;
                        do {
                            (i = a.lane),
                                (mi.lanes |= i),
                                (Is |= i),
                                (a = a.next);
                        } while (a !== e);
                    } else null === a && (n.lanes = 0);
                    return [t.memoizedState, n.dispatch];
                }
                function _i(e) {
                    var t = Ci(),
                        n = t.queue;
                    if (null === n) throw Error(l(311));
                    n.lastRenderedReducer = e;
                    var r = n.dispatch,
                        a = n.pending,
                        i = t.memoizedState;
                    if (null !== a) {
                        n.pending = null;
                        var o = (a = a.next);
                        do {
                            (i = e(i, o.action)), (o = o.next);
                        } while (o !== a);
                        or(i, t.memoizedState) || (xo = !0),
                            (t.memoizedState = i),
                            null === t.baseQueue && (t.baseState = i),
                            (n.lastRenderedState = i);
                    }
                    return [i, r];
                }
                function Pi() {}
                function Oi(e, t) {
                    var n = mi,
                        r = Ci(),
                        a = t(),
                        i = !or(r.memoizedState, a);
                    if (
                        (i && ((r.memoizedState = a), (xo = !0)),
                        (r = r.queue),
                        $i(Ri.bind(null, n, r, e), [e]),
                        r.getSnapshot !== t ||
                            i ||
                            (null !== vi && 1 & vi.memoizedState.tag))
                    ) {
                        if (
                            ((n.flags |= 2048),
                            Di(9, Mi.bind(null, n, r, a, t), void 0, null),
                            null === Ps)
                        )
                            throw Error(l(349));
                        0 !== (30 & pi) || zi(n, t, a);
                    }
                    return a;
                }
                function zi(e, t, n) {
                    (e.flags |= 16384),
                        (e = { getSnapshot: t, value: n }),
                        null === (t = mi.updateQueue)
                            ? ((t = { lastEffect: null, stores: null }),
                              (mi.updateQueue = t),
                              (t.stores = [e]))
                            : null === (n = t.stores)
                            ? (t.stores = [e])
                            : n.push(e);
                }
                function Mi(e, t, n, r) {
                    (t.value = n), (t.getSnapshot = r), Ti(t) && Fi(e);
                }
                function Ri(e, t, n) {
                    return n(function () {
                        Ti(t) && Fi(e);
                    });
                }
                function Ti(e) {
                    var t = e.getSnapshot;
                    e = e.value;
                    try {
                        var n = t();
                        return !or(e, n);
                    } catch (r) {
                        return !0;
                    }
                }
                function Fi(e) {
                    var t = _l(e, 1);
                    null !== t && ru(t, e, 1, -1);
                }
                function Ii(e) {
                    var t = Ni();
                    return (
                        "function" === typeof e && (e = e()),
                        (t.memoizedState = t.baseState = e),
                        (e = {
                            pending: null,
                            interleaved: null,
                            lanes: 0,
                            dispatch: null,
                            lastRenderedReducer: Ei,
                            lastRenderedState: e,
                        }),
                        (t.queue = e),
                        (e = e.dispatch = to.bind(null, mi, e)),
                        [t.memoizedState, e]
                    );
                }
                function Di(e, t, n, r) {
                    return (
                        (e = {
                            tag: e,
                            create: t,
                            destroy: n,
                            deps: r,
                            next: null,
                        }),
                        null === (t = mi.updateQueue)
                            ? ((t = { lastEffect: null, stores: null }),
                              (mi.updateQueue = t),
                              (t.lastEffect = e.next = e))
                            : null === (n = t.lastEffect)
                            ? (t.lastEffect = e.next = e)
                            : ((r = n.next),
                              (n.next = e),
                              (e.next = r),
                              (t.lastEffect = e)),
                        e
                    );
                }
                function Ui() {
                    return Ci().memoizedState;
                }
                function Bi(e, t, n, r) {
                    var a = Ni();
                    (mi.flags |= e),
                        (a.memoizedState = Di(
                            1 | t,
                            n,
                            void 0,
                            void 0 === r ? null : r,
                        ));
                }
                function Ai(e, t, n, r) {
                    var a = Ci();
                    r = void 0 === r ? null : r;
                    var l = void 0;
                    if (null !== hi) {
                        var i = hi.memoizedState;
                        if (((l = i.destroy), null !== r && ji(r, i.deps)))
                            return void (a.memoizedState = Di(t, n, l, r));
                    }
                    (mi.flags |= e), (a.memoizedState = Di(1 | t, n, l, r));
                }
                function Vi(e, t) {
                    return Bi(8390656, 8, e, t);
                }
                function $i(e, t) {
                    return Ai(2048, 8, e, t);
                }
                function Hi(e, t) {
                    return Ai(4, 2, e, t);
                }
                function Ki(e, t) {
                    return Ai(4, 4, e, t);
                }
                function Wi(e, t) {
                    return "function" === typeof t
                        ? ((e = e()),
                          t(e),
                          function () {
                              t(null);
                          })
                        : null !== t && void 0 !== t
                        ? ((e = e()),
                          (t.current = e),
                          function () {
                              t.current = null;
                          })
                        : void 0;
                }
                function Qi(e, t, n) {
                    return (
                        (n = null !== n && void 0 !== n ? n.concat([e]) : null),
                        Ai(4, 4, Wi.bind(null, t, e), n)
                    );
                }
                function qi() {}
                function Gi(e, t) {
                    var n = Ci();
                    t = void 0 === t ? null : t;
                    var r = n.memoizedState;
                    return null !== r && null !== t && ji(t, r[1])
                        ? r[0]
                        : ((n.memoizedState = [e, t]), e);
                }
                function Yi(e, t) {
                    var n = Ci();
                    t = void 0 === t ? null : t;
                    var r = n.memoizedState;
                    return null !== r && null !== t && ji(t, r[1])
                        ? r[0]
                        : ((e = e()), (n.memoizedState = [e, t]), e);
                }
                function Xi(e, t, n) {
                    return 0 === (21 & pi)
                        ? (e.baseState && ((e.baseState = !1), (xo = !0)),
                          (e.memoizedState = n))
                        : (or(n, t) ||
                              ((n = ht()),
                              (mi.lanes |= n),
                              (Is |= n),
                              (e.baseState = !0)),
                          t);
                }
                function Zi(e, t) {
                    var n = bt;
                    (bt = 0 !== n && 4 > n ? n : 4), e(!0);
                    var r = fi.transition;
                    fi.transition = {};
                    try {
                        e(!1), t();
                    } finally {
                        (bt = n), (fi.transition = r);
                    }
                }
                function Ji() {
                    return Ci().memoizedState;
                }
                function eo(e, t, n) {
                    var r = nu(e);
                    if (
                        ((n = {
                            lane: r,
                            action: n,
                            hasEagerState: !1,
                            eagerState: null,
                            next: null,
                        }),
                        no(e))
                    )
                        ro(t, n);
                    else if (null !== (n = Ll(e, t, n, r))) {
                        ru(n, e, r, tu()), ao(n, t, r);
                    }
                }
                function to(e, t, n) {
                    var r = nu(e),
                        a = {
                            lane: r,
                            action: n,
                            hasEagerState: !1,
                            eagerState: null,
                            next: null,
                        };
                    if (no(e)) ro(t, a);
                    else {
                        var l = e.alternate;
                        if (
                            0 === e.lanes &&
                            (null === l || 0 === l.lanes) &&
                            null !== (l = t.lastRenderedReducer)
                        )
                            try {
                                var i = t.lastRenderedState,
                                    o = l(i, n);
                                if (
                                    ((a.hasEagerState = !0),
                                    (a.eagerState = o),
                                    or(o, i))
                                ) {
                                    var s = t.interleaved;
                                    return (
                                        null === s
                                            ? ((a.next = a), El(t))
                                            : ((a.next = s.next), (s.next = a)),
                                        void (t.interleaved = a)
                                    );
                                }
                            } catch (u) {}
                        null !== (n = Ll(e, t, a, r)) &&
                            (ru(n, e, r, (a = tu())), ao(n, t, r));
                    }
                }
                function no(e) {
                    var t = e.alternate;
                    return e === mi || (null !== t && t === mi);
                }
                function ro(e, t) {
                    yi = gi = !0;
                    var n = e.pending;
                    null === n
                        ? (t.next = t)
                        : ((t.next = n.next), (n.next = t)),
                        (e.pending = t);
                }
                function ao(e, t, n) {
                    if (0 !== (4194240 & n)) {
                        var r = t.lanes;
                        (n |= r &= e.pendingLanes), (t.lanes = n), yt(e, n);
                    }
                }
                var lo = {
                        readContext: Nl,
                        useCallback: wi,
                        useContext: wi,
                        useEffect: wi,
                        useImperativeHandle: wi,
                        useInsertionEffect: wi,
                        useLayoutEffect: wi,
                        useMemo: wi,
                        useReducer: wi,
                        useRef: wi,
                        useState: wi,
                        useDebugValue: wi,
                        useDeferredValue: wi,
                        useTransition: wi,
                        useMutableSource: wi,
                        useSyncExternalStore: wi,
                        useId: wi,
                        unstable_isNewReconciler: !1,
                    },
                    io = {
                        readContext: Nl,
                        useCallback: function (e, t) {
                            return (
                                (Ni().memoizedState = [
                                    e,
                                    void 0 === t ? null : t,
                                ]),
                                e
                            );
                        },
                        useContext: Nl,
                        useEffect: Vi,
                        useImperativeHandle: function (e, t, n) {
                            return (
                                (n =
                                    null !== n && void 0 !== n
                                        ? n.concat([e])
                                        : null),
                                Bi(4194308, 4, Wi.bind(null, t, e), n)
                            );
                        },
                        useLayoutEffect: function (e, t) {
                            return Bi(4194308, 4, e, t);
                        },
                        useInsertionEffect: function (e, t) {
                            return Bi(4, 2, e, t);
                        },
                        useMemo: function (e, t) {
                            var n = Ni();
                            return (
                                (t = void 0 === t ? null : t),
                                (e = e()),
                                (n.memoizedState = [e, t]),
                                e
                            );
                        },
                        useReducer: function (e, t, n) {
                            var r = Ni();
                            return (
                                (t = void 0 !== n ? n(t) : t),
                                (r.memoizedState = r.baseState = t),
                                (e = {
                                    pending: null,
                                    interleaved: null,
                                    lanes: 0,
                                    dispatch: null,
                                    lastRenderedReducer: e,
                                    lastRenderedState: t,
                                }),
                                (r.queue = e),
                                (e = e.dispatch = eo.bind(null, mi, e)),
                                [r.memoizedState, e]
                            );
                        },
                        useRef: function (e) {
                            return (
                                (e = { current: e }), (Ni().memoizedState = e)
                            );
                        },
                        useState: Ii,
                        useDebugValue: qi,
                        useDeferredValue: function (e) {
                            return (Ni().memoizedState = e);
                        },
                        useTransition: function () {
                            var e = Ii(!1),
                                t = e[0];
                            return (
                                (e = Zi.bind(null, e[1])),
                                (Ni().memoizedState = e),
                                [t, e]
                            );
                        },
                        useMutableSource: function () {},
                        useSyncExternalStore: function (e, t, n) {
                            var r = mi,
                                a = Ni();
                            if (al) {
                                if (void 0 === n) throw Error(l(407));
                                n = n();
                            } else {
                                if (((n = t()), null === Ps))
                                    throw Error(l(349));
                                0 !== (30 & pi) || zi(r, t, n);
                            }
                            a.memoizedState = n;
                            var i = { value: n, getSnapshot: t };
                            return (
                                (a.queue = i),
                                Vi(Ri.bind(null, r, i, e), [e]),
                                (r.flags |= 2048),
                                Di(9, Mi.bind(null, r, i, n, t), void 0, null),
                                n
                            );
                        },
                        useId: function () {
                            var e = Ni(),
                                t = Ps.identifierPrefix;
                            if (al) {
                                var n = Xa;
                                (t =
                                    ":" +
                                    t +
                                    "R" +
                                    (n =
                                        (
                                            Ya & ~(1 << (32 - it(Ya) - 1))
                                        ).toString(32) + n)),
                                    0 < (n = bi++) &&
                                        (t += "H" + n.toString(32)),
                                    (t += ":");
                            } else
                                t =
                                    ":" +
                                    t +
                                    "r" +
                                    (n = xi++).toString(32) +
                                    ":";
                            return (e.memoizedState = t);
                        },
                        unstable_isNewReconciler: !1,
                    },
                    oo = {
                        readContext: Nl,
                        useCallback: Gi,
                        useContext: Nl,
                        useEffect: $i,
                        useImperativeHandle: Qi,
                        useInsertionEffect: Hi,
                        useLayoutEffect: Ki,
                        useMemo: Yi,
                        useReducer: Li,
                        useRef: Ui,
                        useState: function () {
                            return Li(Ei);
                        },
                        useDebugValue: qi,
                        useDeferredValue: function (e) {
                            return Xi(Ci(), hi.memoizedState, e);
                        },
                        useTransition: function () {
                            return [Li(Ei)[0], Ci().memoizedState];
                        },
                        useMutableSource: Pi,
                        useSyncExternalStore: Oi,
                        useId: Ji,
                        unstable_isNewReconciler: !1,
                    },
                    so = {
                        readContext: Nl,
                        useCallback: Gi,
                        useContext: Nl,
                        useEffect: $i,
                        useImperativeHandle: Qi,
                        useInsertionEffect: Hi,
                        useLayoutEffect: Ki,
                        useMemo: Yi,
                        useReducer: _i,
                        useRef: Ui,
                        useState: function () {
                            return _i(Ei);
                        },
                        useDebugValue: qi,
                        useDeferredValue: function (e) {
                            var t = Ci();
                            return null === hi
                                ? (t.memoizedState = e)
                                : Xi(t, hi.memoizedState, e);
                        },
                        useTransition: function () {
                            return [_i(Ei)[0], Ci().memoizedState];
                        },
                        useMutableSource: Pi,
                        useSyncExternalStore: Oi,
                        useId: Ji,
                        unstable_isNewReconciler: !1,
                    };
                function uo(e, t) {
                    try {
                        var n = "",
                            r = t;
                        do {
                            (n += A(r)), (r = r.return);
                        } while (r);
                        var a = n;
                    } catch (l) {
                        a =
                            "\nerror generating stack: " +
                            l.message +
                            "\n" +
                            l.stack;
                    }
                    return { value: e, source: t, stack: a, digest: null };
                }
                function co(e, t, n) {
                    return {
                        value: e,
                        source: null,
                        stack: null != n ? n : null,
                        digest: null != t ? t : null,
                    };
                }
                function fo(e, t) {
                    try {
                        console.error(t.value);
                    } catch (n) {
                        setTimeout(function () {
                            throw n;
                        });
                    }
                }
                var po = "function" === typeof WeakMap ? WeakMap : Map;
                function mo(e, t, n) {
                    ((n = Ml(-1, n)).tag = 3), (n.payload = { element: null });
                    var r = t.value;
                    return (
                        (n.callback = function () {
                            Ks || ((Ks = !0), (Ws = r)), fo(0, t);
                        }),
                        n
                    );
                }
                function ho(e, t, n) {
                    (n = Ml(-1, n)).tag = 3;
                    var r = e.type.getDerivedStateFromError;
                    if ("function" === typeof r) {
                        var a = t.value;
                        (n.payload = function () {
                            return r(a);
                        }),
                            (n.callback = function () {
                                fo(0, t);
                            });
                    }
                    var l = e.stateNode;
                    return (
                        null !== l &&
                            "function" === typeof l.componentDidCatch &&
                            (n.callback = function () {
                                fo(0, t),
                                    "function" !== typeof r &&
                                        (null === Qs
                                            ? (Qs = new Set([this]))
                                            : Qs.add(this));
                                var e = t.stack;
                                this.componentDidCatch(t.value, {
                                    componentStack: null !== e ? e : "",
                                });
                            }),
                        n
                    );
                }
                function vo(e, t, n) {
                    var r = e.pingCache;
                    if (null === r) {
                        r = e.pingCache = new po();
                        var a = new Set();
                        r.set(t, a);
                    } else
                        void 0 === (a = r.get(t)) &&
                            ((a = new Set()), r.set(t, a));
                    a.has(n) ||
                        (a.add(n), (e = Cu.bind(null, e, t, n)), t.then(e, e));
                }
                function go(e) {
                    do {
                        var t;
                        if (
                            ((t = 13 === e.tag) &&
                                (t =
                                    null === (t = e.memoizedState) ||
                                    null !== t.dehydrated),
                            t)
                        )
                            return e;
                        e = e.return;
                    } while (null !== e);
                    return null;
                }
                function yo(e, t, n, r, a) {
                    return 0 === (1 & e.mode)
                        ? (e === t
                              ? (e.flags |= 65536)
                              : ((e.flags |= 128),
                                (n.flags |= 131072),
                                (n.flags &= -52805),
                                1 === n.tag &&
                                    (null === n.alternate
                                        ? (n.tag = 17)
                                        : (((t = Ml(-1, 1)).tag = 2),
                                          Rl(n, t, 1))),
                                (n.lanes |= 1)),
                          e)
                        : ((e.flags |= 65536), (e.lanes = a), e);
                }
                var bo = x.ReactCurrentOwner,
                    xo = !1;
                function wo(e, t, n, r) {
                    t.child =
                        null === e ? Xl(t, null, n, r) : Yl(t, e.child, n, r);
                }
                function jo(e, t, n, r, a) {
                    n = n.render;
                    var l = t.ref;
                    return (
                        Sl(t, a),
                        (r = ki(e, t, n, r, l, a)),
                        (n = Si()),
                        null === e || xo
                            ? (al && n && el(t),
                              (t.flags |= 1),
                              wo(e, t, r, a),
                              t.child)
                            : ((t.updateQueue = e.updateQueue),
                              (t.flags &= -2053),
                              (e.lanes &= ~a),
                              Ko(e, t, a))
                    );
                }
                function ko(e, t, n, r, a) {
                    if (null === e) {
                        var l = n.type;
                        return "function" !== typeof l ||
                            Mu(l) ||
                            void 0 !== l.defaultProps ||
                            null !== n.compare ||
                            void 0 !== n.defaultProps
                            ? (((e = Tu(n.type, null, r, t, t.mode, a)).ref =
                                  t.ref),
                              (e.return = t),
                              (t.child = e))
                            : ((t.tag = 15), (t.type = l), So(e, t, l, r, a));
                    }
                    if (((l = e.child), 0 === (e.lanes & a))) {
                        var i = l.memoizedProps;
                        if (
                            (n = null !== (n = n.compare) ? n : sr)(i, r) &&
                            e.ref === t.ref
                        )
                            return Ko(e, t, a);
                    }
                    return (
                        (t.flags |= 1),
                        ((e = Ru(l, r)).ref = t.ref),
                        (e.return = t),
                        (t.child = e)
                    );
                }
                function So(e, t, n, r, a) {
                    if (null !== e) {
                        var l = e.memoizedProps;
                        if (sr(l, r) && e.ref === t.ref) {
                            if (
                                ((xo = !1),
                                (t.pendingProps = r = l),
                                0 === (e.lanes & a))
                            )
                                return (t.lanes = e.lanes), Ko(e, t, a);
                            0 !== (131072 & e.flags) && (xo = !0);
                        }
                    }
                    return Eo(e, t, n, r, a);
                }
                function No(e, t, n) {
                    var r = t.pendingProps,
                        a = r.children,
                        l = null !== e ? e.memoizedState : null;
                    if ("hidden" === r.mode)
                        if (0 === (1 & t.mode))
                            (t.memoizedState = {
                                baseLanes: 0,
                                cachePool: null,
                                transitions: null,
                            }),
                                Ca(Rs, Ms),
                                (Ms |= n);
                        else {
                            if (0 === (1073741824 & n))
                                return (
                                    (e = null !== l ? l.baseLanes | n : n),
                                    (t.lanes = t.childLanes = 1073741824),
                                    (t.memoizedState = {
                                        baseLanes: e,
                                        cachePool: null,
                                        transitions: null,
                                    }),
                                    (t.updateQueue = null),
                                    Ca(Rs, Ms),
                                    (Ms |= e),
                                    null
                                );
                            (t.memoizedState = {
                                baseLanes: 0,
                                cachePool: null,
                                transitions: null,
                            }),
                                (r = null !== l ? l.baseLanes : n),
                                Ca(Rs, Ms),
                                (Ms |= r);
                        }
                    else
                        null !== l
                            ? ((r = l.baseLanes | n), (t.memoizedState = null))
                            : (r = n),
                            Ca(Rs, Ms),
                            (Ms |= r);
                    return wo(e, t, a, n), t.child;
                }
                function Co(e, t) {
                    var n = t.ref;
                    ((null === e && null !== n) ||
                        (null !== e && e.ref !== n)) &&
                        ((t.flags |= 512), (t.flags |= 2097152));
                }
                function Eo(e, t, n, r, a) {
                    var l = za(n) ? Pa : La.current;
                    return (
                        (l = Oa(t, l)),
                        Sl(t, a),
                        (n = ki(e, t, n, r, l, a)),
                        (r = Si()),
                        null === e || xo
                            ? (al && r && el(t),
                              (t.flags |= 1),
                              wo(e, t, n, a),
                              t.child)
                            : ((t.updateQueue = e.updateQueue),
                              (t.flags &= -2053),
                              (e.lanes &= ~a),
                              Ko(e, t, a))
                    );
                }
                function Lo(e, t, n, r, a) {
                    if (za(n)) {
                        var l = !0;
                        Fa(t);
                    } else l = !1;
                    if ((Sl(t, a), null === t.stateNode))
                        Ho(e, t), $l(t, n, r), Kl(t, n, r, a), (r = !0);
                    else if (null === e) {
                        var i = t.stateNode,
                            o = t.memoizedProps;
                        i.props = o;
                        var s = i.context,
                            u = n.contextType;
                        "object" === typeof u && null !== u
                            ? (u = Nl(u))
                            : (u = Oa(t, (u = za(n) ? Pa : La.current)));
                        var c = n.getDerivedStateFromProps,
                            d =
                                "function" === typeof c ||
                                "function" === typeof i.getSnapshotBeforeUpdate;
                        d ||
                            ("function" !==
                                typeof i.UNSAFE_componentWillReceiveProps &&
                                "function" !==
                                    typeof i.componentWillReceiveProps) ||
                            ((o !== r || s !== u) && Hl(t, i, r, u)),
                            (Pl = !1);
                        var f = t.memoizedState;
                        (i.state = f),
                            Il(t, r, i, a),
                            (s = t.memoizedState),
                            o !== r || f !== s || _a.current || Pl
                                ? ("function" === typeof c &&
                                      (Bl(t, n, c, r), (s = t.memoizedState)),
                                  (o = Pl || Vl(t, n, o, r, f, s, u))
                                      ? (d ||
                                            ("function" !==
                                                typeof i.UNSAFE_componentWillMount &&
                                                "function" !==
                                                    typeof i.componentWillMount) ||
                                            ("function" ===
                                                typeof i.componentWillMount &&
                                                i.componentWillMount(),
                                            "function" ===
                                                typeof i.UNSAFE_componentWillMount &&
                                                i.UNSAFE_componentWillMount()),
                                        "function" ===
                                            typeof i.componentDidMount &&
                                            (t.flags |= 4194308))
                                      : ("function" ===
                                            typeof i.componentDidMount &&
                                            (t.flags |= 4194308),
                                        (t.memoizedProps = r),
                                        (t.memoizedState = s)),
                                  (i.props = r),
                                  (i.state = s),
                                  (i.context = u),
                                  (r = o))
                                : ("function" === typeof i.componentDidMount &&
                                      (t.flags |= 4194308),
                                  (r = !1));
                    } else {
                        (i = t.stateNode),
                            zl(e, t),
                            (o = t.memoizedProps),
                            (u = t.type === t.elementType ? o : vl(t.type, o)),
                            (i.props = u),
                            (d = t.pendingProps),
                            (f = i.context),
                            "object" === typeof (s = n.contextType) &&
                            null !== s
                                ? (s = Nl(s))
                                : (s = Oa(t, (s = za(n) ? Pa : La.current)));
                        var p = n.getDerivedStateFromProps;
                        (c =
                            "function" === typeof p ||
                            "function" === typeof i.getSnapshotBeforeUpdate) ||
                            ("function" !==
                                typeof i.UNSAFE_componentWillReceiveProps &&
                                "function" !==
                                    typeof i.componentWillReceiveProps) ||
                            ((o !== d || f !== s) && Hl(t, i, r, s)),
                            (Pl = !1),
                            (f = t.memoizedState),
                            (i.state = f),
                            Il(t, r, i, a);
                        var m = t.memoizedState;
                        o !== d || f !== m || _a.current || Pl
                            ? ("function" === typeof p &&
                                  (Bl(t, n, p, r), (m = t.memoizedState)),
                              (u = Pl || Vl(t, n, u, r, f, m, s) || !1)
                                  ? (c ||
                                        ("function" !==
                                            typeof i.UNSAFE_componentWillUpdate &&
                                            "function" !==
                                                typeof i.componentWillUpdate) ||
                                        ("function" ===
                                            typeof i.componentWillUpdate &&
                                            i.componentWillUpdate(r, m, s),
                                        "function" ===
                                            typeof i.UNSAFE_componentWillUpdate &&
                                            i.UNSAFE_componentWillUpdate(
                                                r,
                                                m,
                                                s,
                                            )),
                                    "function" ===
                                        typeof i.componentDidUpdate &&
                                        (t.flags |= 4),
                                    "function" ===
                                        typeof i.getSnapshotBeforeUpdate &&
                                        (t.flags |= 1024))
                                  : ("function" !==
                                        typeof i.componentDidUpdate ||
                                        (o === e.memoizedProps &&
                                            f === e.memoizedState) ||
                                        (t.flags |= 4),
                                    "function" !==
                                        typeof i.getSnapshotBeforeUpdate ||
                                        (o === e.memoizedProps &&
                                            f === e.memoizedState) ||
                                        (t.flags |= 1024),
                                    (t.memoizedProps = r),
                                    (t.memoizedState = m)),
                              (i.props = r),
                              (i.state = m),
                              (i.context = s),
                              (r = u))
                            : ("function" !== typeof i.componentDidUpdate ||
                                  (o === e.memoizedProps &&
                                      f === e.memoizedState) ||
                                  (t.flags |= 4),
                              "function" !== typeof i.getSnapshotBeforeUpdate ||
                                  (o === e.memoizedProps &&
                                      f === e.memoizedState) ||
                                  (t.flags |= 1024),
                              (r = !1));
                    }
                    return _o(e, t, n, r, l, a);
                }
                function _o(e, t, n, r, a, l) {
                    Co(e, t);
                    var i = 0 !== (128 & t.flags);
                    if (!r && !i) return a && Ia(t, n, !1), Ko(e, t, l);
                    (r = t.stateNode), (bo.current = t);
                    var o =
                        i && "function" !== typeof n.getDerivedStateFromError
                            ? null
                            : r.render();
                    return (
                        (t.flags |= 1),
                        null !== e && i
                            ? ((t.child = Yl(t, e.child, null, l)),
                              (t.child = Yl(t, null, o, l)))
                            : wo(e, t, o, l),
                        (t.memoizedState = r.state),
                        a && Ia(t, n, !0),
                        t.child
                    );
                }
                function Po(e) {
                    var t = e.stateNode;
                    t.pendingContext
                        ? Ra(
                              0,
                              t.pendingContext,
                              t.pendingContext !== t.context,
                          )
                        : t.context && Ra(0, t.context, !1),
                        ri(e, t.containerInfo);
                }
                function Oo(e, t, n, r, a) {
                    return (
                        pl(), ml(a), (t.flags |= 256), wo(e, t, n, r), t.child
                    );
                }
                var zo,
                    Mo,
                    Ro,
                    To,
                    Fo = { dehydrated: null, treeContext: null, retryLane: 0 };
                function Io(e) {
                    return { baseLanes: e, cachePool: null, transitions: null };
                }
                function Do(e, t, n) {
                    var r,
                        a = t.pendingProps,
                        i = oi.current,
                        o = !1,
                        s = 0 !== (128 & t.flags);
                    if (
                        ((r = s) ||
                            (r =
                                (null === e || null !== e.memoizedState) &&
                                0 !== (2 & i)),
                        r
                            ? ((o = !0), (t.flags &= -129))
                            : (null !== e && null === e.memoizedState) ||
                              (i |= 1),
                        Ca(oi, 1 & i),
                        null === e)
                    )
                        return (
                            ul(t),
                            null !== (e = t.memoizedState) &&
                            null !== (e = e.dehydrated)
                                ? (0 === (1 & t.mode)
                                      ? (t.lanes = 1)
                                      : "$!" === e.data
                                      ? (t.lanes = 8)
                                      : (t.lanes = 1073741824),
                                  null)
                                : ((s = a.children),
                                  (e = a.fallback),
                                  o
                                      ? ((a = t.mode),
                                        (o = t.child),
                                        (s = { mode: "hidden", children: s }),
                                        0 === (1 & a) && null !== o
                                            ? ((o.childLanes = 0),
                                              (o.pendingProps = s))
                                            : (o = Iu(s, a, 0, null)),
                                        (e = Fu(e, a, n, null)),
                                        (o.return = t),
                                        (e.return = t),
                                        (o.sibling = e),
                                        (t.child = o),
                                        (t.child.memoizedState = Io(n)),
                                        (t.memoizedState = Fo),
                                        e)
                                      : Uo(t, s))
                        );
                    if (
                        null !== (i = e.memoizedState) &&
                        null !== (r = i.dehydrated)
                    )
                        return (function (e, t, n, r, a, i, o) {
                            if (n)
                                return 256 & t.flags
                                    ? ((t.flags &= -257),
                                      Bo(e, t, o, (r = co(Error(l(422))))))
                                    : null !== t.memoizedState
                                    ? ((t.child = e.child),
                                      (t.flags |= 128),
                                      null)
                                    : ((i = r.fallback),
                                      (a = t.mode),
                                      (r = Iu(
                                          {
                                              mode: "visible",
                                              children: r.children,
                                          },
                                          a,
                                          0,
                                          null,
                                      )),
                                      ((i = Fu(i, a, o, null)).flags |= 2),
                                      (r.return = t),
                                      (i.return = t),
                                      (r.sibling = i),
                                      (t.child = r),
                                      0 !== (1 & t.mode) &&
                                          Yl(t, e.child, null, o),
                                      (t.child.memoizedState = Io(o)),
                                      (t.memoizedState = Fo),
                                      i);
                            if (0 === (1 & t.mode)) return Bo(e, t, o, null);
                            if ("$!" === a.data) {
                                if (
                                    (r = a.nextSibling && a.nextSibling.dataset)
                                )
                                    var s = r.dgst;
                                return (
                                    (r = s),
                                    Bo(
                                        e,
                                        t,
                                        o,
                                        (r = co(
                                            (i = Error(l(419))),
                                            r,
                                            void 0,
                                        )),
                                    )
                                );
                            }
                            if (((s = 0 !== (o & e.childLanes)), xo || s)) {
                                if (null !== (r = Ps)) {
                                    switch (o & -o) {
                                        case 4:
                                            a = 2;
                                            break;
                                        case 16:
                                            a = 8;
                                            break;
                                        case 64:
                                        case 128:
                                        case 256:
                                        case 512:
                                        case 1024:
                                        case 2048:
                                        case 4096:
                                        case 8192:
                                        case 16384:
                                        case 32768:
                                        case 65536:
                                        case 131072:
                                        case 262144:
                                        case 524288:
                                        case 1048576:
                                        case 2097152:
                                        case 4194304:
                                        case 8388608:
                                        case 16777216:
                                        case 33554432:
                                        case 67108864:
                                            a = 32;
                                            break;
                                        case 536870912:
                                            a = 268435456;
                                            break;
                                        default:
                                            a = 0;
                                    }
                                    0 !==
                                        (a =
                                            0 !== (a & (r.suspendedLanes | o))
                                                ? 0
                                                : a) &&
                                        a !== i.retryLane &&
                                        ((i.retryLane = a),
                                        _l(e, a),
                                        ru(r, e, a, -1));
                                }
                                return (
                                    vu(), Bo(e, t, o, (r = co(Error(l(421)))))
                                );
                            }
                            return "$?" === a.data
                                ? ((t.flags |= 128),
                                  (t.child = e.child),
                                  (t = Lu.bind(null, e)),
                                  (a._reactRetry = t),
                                  null)
                                : ((e = i.treeContext),
                                  (rl = ua(a.nextSibling)),
                                  (nl = t),
                                  (al = !0),
                                  (ll = null),
                                  null !== e &&
                                      ((Qa[qa++] = Ya),
                                      (Qa[qa++] = Xa),
                                      (Qa[qa++] = Ga),
                                      (Ya = e.id),
                                      (Xa = e.overflow),
                                      (Ga = t)),
                                  (t = Uo(t, r.children)),
                                  (t.flags |= 4096),
                                  t);
                        })(e, t, s, a, r, i, n);
                    if (o) {
                        (o = a.fallback),
                            (s = t.mode),
                            (r = (i = e.child).sibling);
                        var u = { mode: "hidden", children: a.children };
                        return (
                            0 === (1 & s) && t.child !== i
                                ? (((a = t.child).childLanes = 0),
                                  (a.pendingProps = u),
                                  (t.deletions = null))
                                : ((a = Ru(i, u)).subtreeFlags =
                                      14680064 & i.subtreeFlags),
                            null !== r
                                ? (o = Ru(r, o))
                                : ((o = Fu(o, s, n, null)).flags |= 2),
                            (o.return = t),
                            (a.return = t),
                            (a.sibling = o),
                            (t.child = a),
                            (a = o),
                            (o = t.child),
                            (s =
                                null === (s = e.child.memoizedState)
                                    ? Io(n)
                                    : {
                                          baseLanes: s.baseLanes | n,
                                          cachePool: null,
                                          transitions: s.transitions,
                                      }),
                            (o.memoizedState = s),
                            (o.childLanes = e.childLanes & ~n),
                            (t.memoizedState = Fo),
                            a
                        );
                    }
                    return (
                        (e = (o = e.child).sibling),
                        (a = Ru(o, { mode: "visible", children: a.children })),
                        0 === (1 & t.mode) && (a.lanes = n),
                        (a.return = t),
                        (a.sibling = null),
                        null !== e &&
                            (null === (n = t.deletions)
                                ? ((t.deletions = [e]), (t.flags |= 16))
                                : n.push(e)),
                        (t.child = a),
                        (t.memoizedState = null),
                        a
                    );
                }
                function Uo(e, t) {
                    return (
                        ((t = Iu(
                            { mode: "visible", children: t },
                            e.mode,
                            0,
                            null,
                        )).return = e),
                        (e.child = t)
                    );
                }
                function Bo(e, t, n, r) {
                    return (
                        null !== r && ml(r),
                        Yl(t, e.child, null, n),
                        ((e = Uo(t, t.pendingProps.children)).flags |= 2),
                        (t.memoizedState = null),
                        e
                    );
                }
                function Ao(e, t, n) {
                    e.lanes |= t;
                    var r = e.alternate;
                    null !== r && (r.lanes |= t), kl(e.return, t, n);
                }
                function Vo(e, t, n, r, a) {
                    var l = e.memoizedState;
                    null === l
                        ? (e.memoizedState = {
                              isBackwards: t,
                              rendering: null,
                              renderingStartTime: 0,
                              last: r,
                              tail: n,
                              tailMode: a,
                          })
                        : ((l.isBackwards = t),
                          (l.rendering = null),
                          (l.renderingStartTime = 0),
                          (l.last = r),
                          (l.tail = n),
                          (l.tailMode = a));
                }
                function $o(e, t, n) {
                    var r = t.pendingProps,
                        a = r.revealOrder,
                        l = r.tail;
                    if ((wo(e, t, r.children, n), 0 !== (2 & (r = oi.current))))
                        (r = (1 & r) | 2), (t.flags |= 128);
                    else {
                        if (null !== e && 0 !== (128 & e.flags))
                            e: for (e = t.child; null !== e; ) {
                                if (13 === e.tag)
                                    null !== e.memoizedState && Ao(e, n, t);
                                else if (19 === e.tag) Ao(e, n, t);
                                else if (null !== e.child) {
                                    (e.child.return = e), (e = e.child);
                                    continue;
                                }
                                if (e === t) break e;
                                for (; null === e.sibling; ) {
                                    if (null === e.return || e.return === t)
                                        break e;
                                    e = e.return;
                                }
                                (e.sibling.return = e.return), (e = e.sibling);
                            }
                        r &= 1;
                    }
                    if ((Ca(oi, r), 0 === (1 & t.mode))) t.memoizedState = null;
                    else
                        switch (a) {
                            case "forwards":
                                for (n = t.child, a = null; null !== n; )
                                    null !== (e = n.alternate) &&
                                        null === si(e) &&
                                        (a = n),
                                        (n = n.sibling);
                                null === (n = a)
                                    ? ((a = t.child), (t.child = null))
                                    : ((a = n.sibling), (n.sibling = null)),
                                    Vo(t, !1, a, n, l);
                                break;
                            case "backwards":
                                for (
                                    n = null, a = t.child, t.child = null;
                                    null !== a;

                                ) {
                                    if (
                                        null !== (e = a.alternate) &&
                                        null === si(e)
                                    ) {
                                        t.child = a;
                                        break;
                                    }
                                    (e = a.sibling),
                                        (a.sibling = n),
                                        (n = a),
                                        (a = e);
                                }
                                Vo(t, !0, n, null, l);
                                break;
                            case "together":
                                Vo(t, !1, null, null, void 0);
                                break;
                            default:
                                t.memoizedState = null;
                        }
                    return t.child;
                }
                function Ho(e, t) {
                    0 === (1 & t.mode) &&
                        null !== e &&
                        ((e.alternate = null),
                        (t.alternate = null),
                        (t.flags |= 2));
                }
                function Ko(e, t, n) {
                    if (
                        (null !== e && (t.dependencies = e.dependencies),
                        (Is |= t.lanes),
                        0 === (n & t.childLanes))
                    )
                        return null;
                    if (null !== e && t.child !== e.child) throw Error(l(153));
                    if (null !== t.child) {
                        for (
                            n = Ru((e = t.child), e.pendingProps),
                                t.child = n,
                                n.return = t;
                            null !== e.sibling;

                        )
                            (e = e.sibling),
                                ((n = n.sibling =
                                    Ru(e, e.pendingProps)).return = t);
                        n.sibling = null;
                    }
                    return t.child;
                }
                function Wo(e, t) {
                    if (!al)
                        switch (e.tailMode) {
                            case "hidden":
                                t = e.tail;
                                for (var n = null; null !== t; )
                                    null !== t.alternate && (n = t),
                                        (t = t.sibling);
                                null === n
                                    ? (e.tail = null)
                                    : (n.sibling = null);
                                break;
                            case "collapsed":
                                n = e.tail;
                                for (var r = null; null !== n; )
                                    null !== n.alternate && (r = n),
                                        (n = n.sibling);
                                null === r
                                    ? t || null === e.tail
                                        ? (e.tail = null)
                                        : (e.tail.sibling = null)
                                    : (r.sibling = null);
                        }
                }
                function Qo(e) {
                    var t =
                            null !== e.alternate &&
                            e.alternate.child === e.child,
                        n = 0,
                        r = 0;
                    if (t)
                        for (var a = e.child; null !== a; )
                            (n |= a.lanes | a.childLanes),
                                (r |= 14680064 & a.subtreeFlags),
                                (r |= 14680064 & a.flags),
                                (a.return = e),
                                (a = a.sibling);
                    else
                        for (a = e.child; null !== a; )
                            (n |= a.lanes | a.childLanes),
                                (r |= a.subtreeFlags),
                                (r |= a.flags),
                                (a.return = e),
                                (a = a.sibling);
                    return (e.subtreeFlags |= r), (e.childLanes = n), t;
                }
                function qo(e, t, n) {
                    var r = t.pendingProps;
                    switch ((tl(t), t.tag)) {
                        case 2:
                        case 16:
                        case 15:
                        case 0:
                        case 11:
                        case 7:
                        case 8:
                        case 12:
                        case 9:
                        case 14:
                            return Qo(t), null;
                        case 1:
                        case 17:
                            return za(t.type) && Ma(), Qo(t), null;
                        case 3:
                            return (
                                (r = t.stateNode),
                                ai(),
                                Na(_a),
                                Na(La),
                                ci(),
                                r.pendingContext &&
                                    ((r.context = r.pendingContext),
                                    (r.pendingContext = null)),
                                (null !== e && null !== e.child) ||
                                    (dl(t)
                                        ? (t.flags |= 4)
                                        : null === e ||
                                          (e.memoizedState.isDehydrated &&
                                              0 === (256 & t.flags)) ||
                                          ((t.flags |= 1024),
                                          null !== ll &&
                                              (ou(ll), (ll = null)))),
                                Mo(e, t),
                                Qo(t),
                                null
                            );
                        case 5:
                            ii(t);
                            var a = ni(ti.current);
                            if (
                                ((n = t.type),
                                null !== e && null != t.stateNode)
                            )
                                Ro(e, t, n, r, a),
                                    e.ref !== t.ref &&
                                        ((t.flags |= 512),
                                        (t.flags |= 2097152));
                            else {
                                if (!r) {
                                    if (null === t.stateNode)
                                        throw Error(l(166));
                                    return Qo(t), null;
                                }
                                if (((e = ni(Jl.current)), dl(t))) {
                                    (r = t.stateNode), (n = t.type);
                                    var i = t.memoizedProps;
                                    switch (
                                        ((r[fa] = t),
                                        (r[pa] = i),
                                        (e = 0 !== (1 & t.mode)),
                                        n)
                                    ) {
                                        case "dialog":
                                            Ur("cancel", r), Ur("close", r);
                                            break;
                                        case "iframe":
                                        case "object":
                                        case "embed":
                                            Ur("load", r);
                                            break;
                                        case "video":
                                        case "audio":
                                            for (a = 0; a < Tr.length; a++)
                                                Ur(Tr[a], r);
                                            break;
                                        case "source":
                                            Ur("error", r);
                                            break;
                                        case "img":
                                        case "image":
                                        case "link":
                                            Ur("error", r), Ur("load", r);
                                            break;
                                        case "details":
                                            Ur("toggle", r);
                                            break;
                                        case "input":
                                            Y(r, i), Ur("invalid", r);
                                            break;
                                        case "select":
                                            (r._wrapperState = {
                                                wasMultiple: !!i.multiple,
                                            }),
                                                Ur("invalid", r);
                                            break;
                                        case "textarea":
                                            ae(r, i), Ur("invalid", r);
                                    }
                                    for (var s in (ye(n, i), (a = null), i))
                                        if (i.hasOwnProperty(s)) {
                                            var u = i[s];
                                            "children" === s
                                                ? "string" === typeof u
                                                    ? r.textContent !== u &&
                                                      (!0 !==
                                                          i.suppressHydrationWarning &&
                                                          Zr(
                                                              r.textContent,
                                                              u,
                                                              e,
                                                          ),
                                                      (a = ["children", u]))
                                                    : "number" === typeof u &&
                                                      r.textContent !==
                                                          "" + u &&
                                                      (!0 !==
                                                          i.suppressHydrationWarning &&
                                                          Zr(
                                                              r.textContent,
                                                              u,
                                                              e,
                                                          ),
                                                      (a = [
                                                          "children",
                                                          "" + u,
                                                      ]))
                                                : o.hasOwnProperty(s) &&
                                                  null != u &&
                                                  "onScroll" === s &&
                                                  Ur("scroll", r);
                                        }
                                    switch (n) {
                                        case "input":
                                            W(r), J(r, i, !0);
                                            break;
                                        case "textarea":
                                            W(r), ie(r);
                                            break;
                                        case "select":
                                        case "option":
                                            break;
                                        default:
                                            "function" === typeof i.onClick &&
                                                (r.onclick = Jr);
                                    }
                                    (r = a),
                                        (t.updateQueue = r),
                                        null !== r && (t.flags |= 4);
                                } else {
                                    (s =
                                        9 === a.nodeType ? a : a.ownerDocument),
                                        "http://www.w3.org/1999/xhtml" === e &&
                                            (e = oe(n)),
                                        "http://www.w3.org/1999/xhtml" === e
                                            ? "script" === n
                                                ? (((e =
                                                      s.createElement(
                                                          "div",
                                                      )).innerHTML =
                                                      "<script></script>"),
                                                  (e = e.removeChild(
                                                      e.firstChild,
                                                  )))
                                                : "string" === typeof r.is
                                                ? (e = s.createElement(n, {
                                                      is: r.is,
                                                  }))
                                                : ((e = s.createElement(n)),
                                                  "select" === n &&
                                                      ((s = e),
                                                      r.multiple
                                                          ? (s.multiple = !0)
                                                          : r.size &&
                                                            (s.size = r.size)))
                                            : (e = s.createElementNS(e, n)),
                                        (e[fa] = t),
                                        (e[pa] = r),
                                        zo(e, t, !1, !1),
                                        (t.stateNode = e);
                                    e: {
                                        switch (((s = be(n, r)), n)) {
                                            case "dialog":
                                                Ur("cancel", e),
                                                    Ur("close", e),
                                                    (a = r);
                                                break;
                                            case "iframe":
                                            case "object":
                                            case "embed":
                                                Ur("load", e), (a = r);
                                                break;
                                            case "video":
                                            case "audio":
                                                for (a = 0; a < Tr.length; a++)
                                                    Ur(Tr[a], e);
                                                a = r;
                                                break;
                                            case "source":
                                                Ur("error", e), (a = r);
                                                break;
                                            case "img":
                                            case "image":
                                            case "link":
                                                Ur("error", e),
                                                    Ur("load", e),
                                                    (a = r);
                                                break;
                                            case "details":
                                                Ur("toggle", e), (a = r);
                                                break;
                                            case "input":
                                                Y(e, r),
                                                    (a = G(e, r)),
                                                    Ur("invalid", e);
                                                break;
                                            case "option":
                                            default:
                                                a = r;
                                                break;
                                            case "select":
                                                (e._wrapperState = {
                                                    wasMultiple: !!r.multiple,
                                                }),
                                                    (a = I({}, r, {
                                                        value: void 0,
                                                    })),
                                                    Ur("invalid", e);
                                                break;
                                            case "textarea":
                                                ae(e, r),
                                                    (a = re(e, r)),
                                                    Ur("invalid", e);
                                        }
                                        for (i in (ye(n, a), (u = a)))
                                            if (u.hasOwnProperty(i)) {
                                                var c = u[i];
                                                "style" === i
                                                    ? ve(e, c)
                                                    : "dangerouslySetInnerHTML" ===
                                                      i
                                                    ? null !=
                                                          (c = c
                                                              ? c.__html
                                                              : void 0) &&
                                                      de(e, c)
                                                    : "children" === i
                                                    ? "string" === typeof c
                                                        ? ("textarea" !== n ||
                                                              "" !== c) &&
                                                          fe(e, c)
                                                        : "number" ===
                                                              typeof c &&
                                                          fe(e, "" + c)
                                                    : "suppressContentEditableWarning" !==
                                                          i &&
                                                      "suppressHydrationWarning" !==
                                                          i &&
                                                      "autoFocus" !== i &&
                                                      (o.hasOwnProperty(i)
                                                          ? null != c &&
                                                            "onScroll" === i &&
                                                            Ur("scroll", e)
                                                          : null != c &&
                                                            b(e, i, c, s));
                                            }
                                        switch (n) {
                                            case "input":
                                                W(e), J(e, r, !1);
                                                break;
                                            case "textarea":
                                                W(e), ie(e);
                                                break;
                                            case "option":
                                                null != r.value &&
                                                    e.setAttribute(
                                                        "value",
                                                        "" + H(r.value),
                                                    );
                                                break;
                                            case "select":
                                                (e.multiple = !!r.multiple),
                                                    null != (i = r.value)
                                                        ? ne(
                                                              e,
                                                              !!r.multiple,
                                                              i,
                                                              !1,
                                                          )
                                                        : null !=
                                                              r.defaultValue &&
                                                          ne(
                                                              e,
                                                              !!r.multiple,
                                                              r.defaultValue,
                                                              !0,
                                                          );
                                                break;
                                            default:
                                                "function" ===
                                                    typeof a.onClick &&
                                                    (e.onclick = Jr);
                                        }
                                        switch (n) {
                                            case "button":
                                            case "input":
                                            case "select":
                                            case "textarea":
                                                r = !!r.autoFocus;
                                                break e;
                                            case "img":
                                                r = !0;
                                                break e;
                                            default:
                                                r = !1;
                                        }
                                    }
                                    r && (t.flags |= 4);
                                }
                                null !== t.ref &&
                                    ((t.flags |= 512), (t.flags |= 2097152));
                            }
                            return Qo(t), null;
                        case 6:
                            if (e && null != t.stateNode)
                                To(e, t, e.memoizedProps, r);
                            else {
                                if (
                                    "string" !== typeof r &&
                                    null === t.stateNode
                                )
                                    throw Error(l(166));
                                if (
                                    ((n = ni(ti.current)),
                                    ni(Jl.current),
                                    dl(t))
                                ) {
                                    if (
                                        ((r = t.stateNode),
                                        (n = t.memoizedProps),
                                        (r[fa] = t),
                                        (i = r.nodeValue !== n) &&
                                            null !== (e = nl))
                                    )
                                        switch (e.tag) {
                                            case 3:
                                                Zr(
                                                    r.nodeValue,
                                                    n,
                                                    0 !== (1 & e.mode),
                                                );
                                                break;
                                            case 5:
                                                !0 !==
                                                    e.memoizedProps
                                                        .suppressHydrationWarning &&
                                                    Zr(
                                                        r.nodeValue,
                                                        n,
                                                        0 !== (1 & e.mode),
                                                    );
                                        }
                                    i && (t.flags |= 4);
                                } else
                                    ((r = (
                                        9 === n.nodeType ? n : n.ownerDocument
                                    ).createTextNode(r))[fa] = t),
                                        (t.stateNode = r);
                            }
                            return Qo(t), null;
                        case 13:
                            if (
                                (Na(oi),
                                (r = t.memoizedState),
                                null === e ||
                                    (null !== e.memoizedState &&
                                        null !== e.memoizedState.dehydrated))
                            ) {
                                if (
                                    al &&
                                    null !== rl &&
                                    0 !== (1 & t.mode) &&
                                    0 === (128 & t.flags)
                                )
                                    fl(), pl(), (t.flags |= 98560), (i = !1);
                                else if (
                                    ((i = dl(t)),
                                    null !== r && null !== r.dehydrated)
                                ) {
                                    if (null === e) {
                                        if (!i) throw Error(l(318));
                                        if (
                                            !(i =
                                                null !== (i = t.memoizedState)
                                                    ? i.dehydrated
                                                    : null)
                                        )
                                            throw Error(l(317));
                                        i[fa] = t;
                                    } else
                                        pl(),
                                            0 === (128 & t.flags) &&
                                                (t.memoizedState = null),
                                            (t.flags |= 4);
                                    Qo(t), (i = !1);
                                } else
                                    null !== ll && (ou(ll), (ll = null)),
                                        (i = !0);
                                if (!i) return 65536 & t.flags ? t : null;
                            }
                            return 0 !== (128 & t.flags)
                                ? ((t.lanes = n), t)
                                : ((r = null !== r) !==
                                      (null !== e &&
                                          null !== e.memoizedState) &&
                                      r &&
                                      ((t.child.flags |= 8192),
                                      0 !== (1 & t.mode) &&
                                          (null === e || 0 !== (1 & oi.current)
                                              ? 0 === Ts && (Ts = 3)
                                              : vu())),
                                  null !== t.updateQueue && (t.flags |= 4),
                                  Qo(t),
                                  null);
                        case 4:
                            return (
                                ai(),
                                Mo(e, t),
                                null === e && Vr(t.stateNode.containerInfo),
                                Qo(t),
                                null
                            );
                        case 10:
                            return jl(t.type._context), Qo(t), null;
                        case 19:
                            if ((Na(oi), null === (i = t.memoizedState)))
                                return Qo(t), null;
                            if (
                                ((r = 0 !== (128 & t.flags)),
                                null === (s = i.rendering))
                            )
                                if (r) Wo(i, !1);
                                else {
                                    if (
                                        0 !== Ts ||
                                        (null !== e && 0 !== (128 & e.flags))
                                    )
                                        for (e = t.child; null !== e; ) {
                                            if (null !== (s = si(e))) {
                                                for (
                                                    t.flags |= 128,
                                                        Wo(i, !1),
                                                        null !==
                                                            (r =
                                                                s.updateQueue) &&
                                                            ((t.updateQueue =
                                                                r),
                                                            (t.flags |= 4)),
                                                        t.subtreeFlags = 0,
                                                        r = n,
                                                        n = t.child;
                                                    null !== n;

                                                )
                                                    (e = r),
                                                        ((i =
                                                            n).flags &= 14680066),
                                                        null ===
                                                        (s = i.alternate)
                                                            ? ((i.childLanes = 0),
                                                              (i.lanes = e),
                                                              (i.child = null),
                                                              (i.subtreeFlags = 0),
                                                              (i.memoizedProps =
                                                                  null),
                                                              (i.memoizedState =
                                                                  null),
                                                              (i.updateQueue =
                                                                  null),
                                                              (i.dependencies =
                                                                  null),
                                                              (i.stateNode =
                                                                  null))
                                                            : ((i.childLanes =
                                                                  s.childLanes),
                                                              (i.lanes =
                                                                  s.lanes),
                                                              (i.child =
                                                                  s.child),
                                                              (i.subtreeFlags = 0),
                                                              (i.deletions =
                                                                  null),
                                                              (i.memoizedProps =
                                                                  s.memoizedProps),
                                                              (i.memoizedState =
                                                                  s.memoizedState),
                                                              (i.updateQueue =
                                                                  s.updateQueue),
                                                              (i.type = s.type),
                                                              (e =
                                                                  s.dependencies),
                                                              (i.dependencies =
                                                                  null === e
                                                                      ? null
                                                                      : {
                                                                            lanes: e.lanes,
                                                                            firstContext:
                                                                                e.firstContext,
                                                                        })),
                                                        (n = n.sibling);
                                                return (
                                                    Ca(
                                                        oi,
                                                        (1 & oi.current) | 2,
                                                    ),
                                                    t.child
                                                );
                                            }
                                            e = e.sibling;
                                        }
                                    null !== i.tail &&
                                        Xe() > $s &&
                                        ((t.flags |= 128),
                                        (r = !0),
                                        Wo(i, !1),
                                        (t.lanes = 4194304));
                                }
                            else {
                                if (!r)
                                    if (null !== (e = si(s))) {
                                        if (
                                            ((t.flags |= 128),
                                            (r = !0),
                                            null !== (n = e.updateQueue) &&
                                                ((t.updateQueue = n),
                                                (t.flags |= 4)),
                                            Wo(i, !0),
                                            null === i.tail &&
                                                "hidden" === i.tailMode &&
                                                !s.alternate &&
                                                !al)
                                        )
                                            return Qo(t), null;
                                    } else
                                        2 * Xe() - i.renderingStartTime > $s &&
                                            1073741824 !== n &&
                                            ((t.flags |= 128),
                                            (r = !0),
                                            Wo(i, !1),
                                            (t.lanes = 4194304));
                                i.isBackwards
                                    ? ((s.sibling = t.child), (t.child = s))
                                    : (null !== (n = i.last)
                                          ? (n.sibling = s)
                                          : (t.child = s),
                                      (i.last = s));
                            }
                            return null !== i.tail
                                ? ((t = i.tail),
                                  (i.rendering = t),
                                  (i.tail = t.sibling),
                                  (i.renderingStartTime = Xe()),
                                  (t.sibling = null),
                                  (n = oi.current),
                                  Ca(oi, r ? (1 & n) | 2 : 1 & n),
                                  t)
                                : (Qo(t), null);
                        case 22:
                        case 23:
                            return (
                                fu(),
                                (r = null !== t.memoizedState),
                                null !== e &&
                                    (null !== e.memoizedState) !== r &&
                                    (t.flags |= 8192),
                                r && 0 !== (1 & t.mode)
                                    ? 0 !== (1073741824 & Ms) &&
                                      (Qo(t),
                                      6 & t.subtreeFlags && (t.flags |= 8192))
                                    : Qo(t),
                                null
                            );
                        case 24:
                        case 25:
                            return null;
                    }
                    throw Error(l(156, t.tag));
                }
                function Go(e, t) {
                    switch ((tl(t), t.tag)) {
                        case 1:
                            return (
                                za(t.type) && Ma(),
                                65536 & (e = t.flags)
                                    ? ((t.flags = (-65537 & e) | 128), t)
                                    : null
                            );
                        case 3:
                            return (
                                ai(),
                                Na(_a),
                                Na(La),
                                ci(),
                                0 !== (65536 & (e = t.flags)) && 0 === (128 & e)
                                    ? ((t.flags = (-65537 & e) | 128), t)
                                    : null
                            );
                        case 5:
                            return ii(t), null;
                        case 13:
                            if (
                                (Na(oi),
                                null !== (e = t.memoizedState) &&
                                    null !== e.dehydrated)
                            ) {
                                if (null === t.alternate) throw Error(l(340));
                                pl();
                            }
                            return 65536 & (e = t.flags)
                                ? ((t.flags = (-65537 & e) | 128), t)
                                : null;
                        case 19:
                            return Na(oi), null;
                        case 4:
                            return ai(), null;
                        case 10:
                            return jl(t.type._context), null;
                        case 22:
                        case 23:
                            return fu(), null;
                        default:
                            return null;
                    }
                }
                (zo = function (e, t) {
                    for (var n = t.child; null !== n; ) {
                        if (5 === n.tag || 6 === n.tag)
                            e.appendChild(n.stateNode);
                        else if (4 !== n.tag && null !== n.child) {
                            (n.child.return = n), (n = n.child);
                            continue;
                        }
                        if (n === t) break;
                        for (; null === n.sibling; ) {
                            if (null === n.return || n.return === t) return;
                            n = n.return;
                        }
                        (n.sibling.return = n.return), (n = n.sibling);
                    }
                }),
                    (Mo = function () {}),
                    (Ro = function (e, t, n, r) {
                        var a = e.memoizedProps;
                        if (a !== r) {
                            (e = t.stateNode), ni(Jl.current);
                            var l,
                                i = null;
                            switch (n) {
                                case "input":
                                    (a = G(e, a)), (r = G(e, r)), (i = []);
                                    break;
                                case "select":
                                    (a = I({}, a, { value: void 0 })),
                                        (r = I({}, r, { value: void 0 })),
                                        (i = []);
                                    break;
                                case "textarea":
                                    (a = re(e, a)), (r = re(e, r)), (i = []);
                                    break;
                                default:
                                    "function" !== typeof a.onClick &&
                                        "function" === typeof r.onClick &&
                                        (e.onclick = Jr);
                            }
                            for (c in (ye(n, r), (n = null), a))
                                if (
                                    !r.hasOwnProperty(c) &&
                                    a.hasOwnProperty(c) &&
                                    null != a[c]
                                )
                                    if ("style" === c) {
                                        var s = a[c];
                                        for (l in s)
                                            s.hasOwnProperty(l) &&
                                                (n || (n = {}), (n[l] = ""));
                                    } else
                                        "dangerouslySetInnerHTML" !== c &&
                                            "children" !== c &&
                                            "suppressContentEditableWarning" !==
                                                c &&
                                            "suppressHydrationWarning" !== c &&
                                            "autoFocus" !== c &&
                                            (o.hasOwnProperty(c)
                                                ? i || (i = [])
                                                : (i = i || []).push(c, null));
                            for (c in r) {
                                var u = r[c];
                                if (
                                    ((s = null != a ? a[c] : void 0),
                                    r.hasOwnProperty(c) &&
                                        u !== s &&
                                        (null != u || null != s))
                                )
                                    if ("style" === c)
                                        if (s) {
                                            for (l in s)
                                                !s.hasOwnProperty(l) ||
                                                    (u &&
                                                        u.hasOwnProperty(l)) ||
                                                    (n || (n = {}),
                                                    (n[l] = ""));
                                            for (l in u)
                                                u.hasOwnProperty(l) &&
                                                    s[l] !== u[l] &&
                                                    (n || (n = {}),
                                                    (n[l] = u[l]));
                                        } else
                                            n || (i || (i = []), i.push(c, n)),
                                                (n = u);
                                    else
                                        "dangerouslySetInnerHTML" === c
                                            ? ((u = u ? u.__html : void 0),
                                              (s = s ? s.__html : void 0),
                                              null != u &&
                                                  s !== u &&
                                                  (i = i || []).push(c, u))
                                            : "children" === c
                                            ? ("string" !== typeof u &&
                                                  "number" !== typeof u) ||
                                              (i = i || []).push(c, "" + u)
                                            : "suppressContentEditableWarning" !==
                                                  c &&
                                              "suppressHydrationWarning" !==
                                                  c &&
                                              (o.hasOwnProperty(c)
                                                  ? (null != u &&
                                                        "onScroll" === c &&
                                                        Ur("scroll", e),
                                                    i || s === u || (i = []))
                                                  : (i = i || []).push(c, u));
                            }
                            n && (i = i || []).push("style", n);
                            var c = i;
                            (t.updateQueue = c) && (t.flags |= 4);
                        }
                    }),
                    (To = function (e, t, n, r) {
                        n !== r && (t.flags |= 4);
                    });
                var Yo = !1,
                    Xo = !1,
                    Zo = "function" === typeof WeakSet ? WeakSet : Set,
                    Jo = null;
                function es(e, t) {
                    var n = e.ref;
                    if (null !== n)
                        if ("function" === typeof n)
                            try {
                                n(null);
                            } catch (r) {
                                Nu(e, t, r);
                            }
                        else n.current = null;
                }
                function ts(e, t, n) {
                    try {
                        n();
                    } catch (r) {
                        Nu(e, t, r);
                    }
                }
                var ns = !1;
                function rs(e, t, n) {
                    var r = t.updateQueue;
                    if (null !== (r = null !== r ? r.lastEffect : null)) {
                        var a = (r = r.next);
                        do {
                            if ((a.tag & e) === e) {
                                var l = a.destroy;
                                (a.destroy = void 0),
                                    void 0 !== l && ts(t, n, l);
                            }
                            a = a.next;
                        } while (a !== r);
                    }
                }
                function as(e, t) {
                    if (
                        null !==
                        (t = null !== (t = t.updateQueue) ? t.lastEffect : null)
                    ) {
                        var n = (t = t.next);
                        do {
                            if ((n.tag & e) === e) {
                                var r = n.create;
                                n.destroy = r();
                            }
                            n = n.next;
                        } while (n !== t);
                    }
                }
                function ls(e) {
                    var t = e.ref;
                    if (null !== t) {
                        var n = e.stateNode;
                        e.tag,
                            (e = n),
                            "function" === typeof t ? t(e) : (t.current = e);
                    }
                }
                function is(e) {
                    var t = e.alternate;
                    null !== t && ((e.alternate = null), is(t)),
                        (e.child = null),
                        (e.deletions = null),
                        (e.sibling = null),
                        5 === e.tag &&
                            null !== (t = e.stateNode) &&
                            (delete t[fa],
                            delete t[pa],
                            delete t[ha],
                            delete t[va],
                            delete t[ga]),
                        (e.stateNode = null),
                        (e.return = null),
                        (e.dependencies = null),
                        (e.memoizedProps = null),
                        (e.memoizedState = null),
                        (e.pendingProps = null),
                        (e.stateNode = null),
                        (e.updateQueue = null);
                }
                function os(e) {
                    return 5 === e.tag || 3 === e.tag || 4 === e.tag;
                }
                function ss(e) {
                    e: for (;;) {
                        for (; null === e.sibling; ) {
                            if (null === e.return || os(e.return)) return null;
                            e = e.return;
                        }
                        for (
                            e.sibling.return = e.return, e = e.sibling;
                            5 !== e.tag && 6 !== e.tag && 18 !== e.tag;

                        ) {
                            if (2 & e.flags) continue e;
                            if (null === e.child || 4 === e.tag) continue e;
                            (e.child.return = e), (e = e.child);
                        }
                        if (!(2 & e.flags)) return e.stateNode;
                    }
                }
                function us(e, t, n) {
                    var r = e.tag;
                    if (5 === r || 6 === r)
                        (e = e.stateNode),
                            t
                                ? 8 === n.nodeType
                                    ? n.parentNode.insertBefore(e, t)
                                    : n.insertBefore(e, t)
                                : (8 === n.nodeType
                                      ? (t = n.parentNode).insertBefore(e, n)
                                      : (t = n).appendChild(e),
                                  (null !== (n = n._reactRootContainer) &&
                                      void 0 !== n) ||
                                      null !== t.onclick ||
                                      (t.onclick = Jr));
                    else if (4 !== r && null !== (e = e.child))
                        for (us(e, t, n), e = e.sibling; null !== e; )
                            us(e, t, n), (e = e.sibling);
                }
                function cs(e, t, n) {
                    var r = e.tag;
                    if (5 === r || 6 === r)
                        (e = e.stateNode),
                            t ? n.insertBefore(e, t) : n.appendChild(e);
                    else if (4 !== r && null !== (e = e.child))
                        for (cs(e, t, n), e = e.sibling; null !== e; )
                            cs(e, t, n), (e = e.sibling);
                }
                var ds = null,
                    fs = !1;
                function ps(e, t, n) {
                    for (n = n.child; null !== n; )
                        ms(e, t, n), (n = n.sibling);
                }
                function ms(e, t, n) {
                    if (lt && "function" === typeof lt.onCommitFiberUnmount)
                        try {
                            lt.onCommitFiberUnmount(at, n);
                        } catch (o) {}
                    switch (n.tag) {
                        case 5:
                            Xo || es(n, t);
                        case 6:
                            var r = ds,
                                a = fs;
                            (ds = null),
                                ps(e, t, n),
                                (fs = a),
                                null !== (ds = r) &&
                                    (fs
                                        ? ((e = ds),
                                          (n = n.stateNode),
                                          8 === e.nodeType
                                              ? e.parentNode.removeChild(n)
                                              : e.removeChild(n))
                                        : ds.removeChild(n.stateNode));
                            break;
                        case 18:
                            null !== ds &&
                                (fs
                                    ? ((e = ds),
                                      (n = n.stateNode),
                                      8 === e.nodeType
                                          ? sa(e.parentNode, n)
                                          : 1 === e.nodeType && sa(e, n),
                                      Vt(e))
                                    : sa(ds, n.stateNode));
                            break;
                        case 4:
                            (r = ds),
                                (a = fs),
                                (ds = n.stateNode.containerInfo),
                                (fs = !0),
                                ps(e, t, n),
                                (ds = r),
                                (fs = a);
                            break;
                        case 0:
                        case 11:
                        case 14:
                        case 15:
                            if (
                                !Xo &&
                                null !== (r = n.updateQueue) &&
                                null !== (r = r.lastEffect)
                            ) {
                                a = r = r.next;
                                do {
                                    var l = a,
                                        i = l.destroy;
                                    (l = l.tag),
                                        void 0 !== i &&
                                            (0 !== (2 & l) || 0 !== (4 & l)) &&
                                            ts(n, t, i),
                                        (a = a.next);
                                } while (a !== r);
                            }
                            ps(e, t, n);
                            break;
                        case 1:
                            if (
                                !Xo &&
                                (es(n, t),
                                "function" ===
                                    typeof (r = n.stateNode)
                                        .componentWillUnmount)
                            )
                                try {
                                    (r.props = n.memoizedProps),
                                        (r.state = n.memoizedState),
                                        r.componentWillUnmount();
                                } catch (o) {
                                    Nu(n, t, o);
                                }
                            ps(e, t, n);
                            break;
                        case 21:
                            ps(e, t, n);
                            break;
                        case 22:
                            1 & n.mode
                                ? ((Xo = (r = Xo) || null !== n.memoizedState),
                                  ps(e, t, n),
                                  (Xo = r))
                                : ps(e, t, n);
                            break;
                        default:
                            ps(e, t, n);
                    }
                }
                function hs(e) {
                    var t = e.updateQueue;
                    if (null !== t) {
                        e.updateQueue = null;
                        var n = e.stateNode;
                        null === n && (n = e.stateNode = new Zo()),
                            t.forEach(function (t) {
                                var r = _u.bind(null, e, t);
                                n.has(t) || (n.add(t), t.then(r, r));
                            });
                    }
                }
                function vs(e, t) {
                    var n = t.deletions;
                    if (null !== n)
                        for (var r = 0; r < n.length; r++) {
                            var a = n[r];
                            try {
                                var i = e,
                                    o = t,
                                    s = o;
                                e: for (; null !== s; ) {
                                    switch (s.tag) {
                                        case 5:
                                            (ds = s.stateNode), (fs = !1);
                                            break e;
                                        case 3:
                                        case 4:
                                            (ds = s.stateNode.containerInfo),
                                                (fs = !0);
                                            break e;
                                    }
                                    s = s.return;
                                }
                                if (null === ds) throw Error(l(160));
                                ms(i, o, a), (ds = null), (fs = !1);
                                var u = a.alternate;
                                null !== u && (u.return = null),
                                    (a.return = null);
                            } catch (c) {
                                Nu(a, t, c);
                            }
                        }
                    if (12854 & t.subtreeFlags)
                        for (t = t.child; null !== t; )
                            gs(t, e), (t = t.sibling);
                }
                function gs(e, t) {
                    var n = e.alternate,
                        r = e.flags;
                    switch (e.tag) {
                        case 0:
                        case 11:
                        case 14:
                        case 15:
                            if ((vs(t, e), ys(e), 4 & r)) {
                                try {
                                    rs(3, e, e.return), as(3, e);
                                } catch (v) {
                                    Nu(e, e.return, v);
                                }
                                try {
                                    rs(5, e, e.return);
                                } catch (v) {
                                    Nu(e, e.return, v);
                                }
                            }
                            break;
                        case 1:
                            vs(t, e),
                                ys(e),
                                512 & r && null !== n && es(n, n.return);
                            break;
                        case 5:
                            if (
                                (vs(t, e),
                                ys(e),
                                512 & r && null !== n && es(n, n.return),
                                32 & e.flags)
                            ) {
                                var a = e.stateNode;
                                try {
                                    fe(a, "");
                                } catch (v) {
                                    Nu(e, e.return, v);
                                }
                            }
                            if (4 & r && null != (a = e.stateNode)) {
                                var i = e.memoizedProps,
                                    o = null !== n ? n.memoizedProps : i,
                                    s = e.type,
                                    u = e.updateQueue;
                                if (((e.updateQueue = null), null !== u))
                                    try {
                                        "input" === s &&
                                            "radio" === i.type &&
                                            null != i.name &&
                                            X(a, i),
                                            be(s, o);
                                        var c = be(s, i);
                                        for (o = 0; o < u.length; o += 2) {
                                            var d = u[o],
                                                f = u[o + 1];
                                            "style" === d
                                                ? ve(a, f)
                                                : "dangerouslySetInnerHTML" ===
                                                  d
                                                ? de(a, f)
                                                : "children" === d
                                                ? fe(a, f)
                                                : b(a, d, f, c);
                                        }
                                        switch (s) {
                                            case "input":
                                                Z(a, i);
                                                break;
                                            case "textarea":
                                                le(a, i);
                                                break;
                                            case "select":
                                                var p =
                                                    a._wrapperState.wasMultiple;
                                                a._wrapperState.wasMultiple =
                                                    !!i.multiple;
                                                var m = i.value;
                                                null != m
                                                    ? ne(a, !!i.multiple, m, !1)
                                                    : p !== !!i.multiple &&
                                                      (null != i.defaultValue
                                                          ? ne(
                                                                a,
                                                                !!i.multiple,
                                                                i.defaultValue,
                                                                !0,
                                                            )
                                                          : ne(
                                                                a,
                                                                !!i.multiple,
                                                                i.multiple
                                                                    ? []
                                                                    : "",
                                                                !1,
                                                            ));
                                        }
                                        a[pa] = i;
                                    } catch (v) {
                                        Nu(e, e.return, v);
                                    }
                            }
                            break;
                        case 6:
                            if ((vs(t, e), ys(e), 4 & r)) {
                                if (null === e.stateNode) throw Error(l(162));
                                (a = e.stateNode), (i = e.memoizedProps);
                                try {
                                    a.nodeValue = i;
                                } catch (v) {
                                    Nu(e, e.return, v);
                                }
                            }
                            break;
                        case 3:
                            if (
                                (vs(t, e),
                                ys(e),
                                4 & r &&
                                    null !== n &&
                                    n.memoizedState.isDehydrated)
                            )
                                try {
                                    Vt(t.containerInfo);
                                } catch (v) {
                                    Nu(e, e.return, v);
                                }
                            break;
                        case 4:
                        default:
                            vs(t, e), ys(e);
                            break;
                        case 13:
                            vs(t, e),
                                ys(e),
                                8192 & (a = e.child).flags &&
                                    ((i = null !== a.memoizedState),
                                    (a.stateNode.isHidden = i),
                                    !i ||
                                        (null !== a.alternate &&
                                            null !==
                                                a.alternate.memoizedState) ||
                                        (Vs = Xe())),
                                4 & r && hs(e);
                            break;
                        case 22:
                            if (
                                ((d = null !== n && null !== n.memoizedState),
                                1 & e.mode
                                    ? ((Xo = (c = Xo) || d), vs(t, e), (Xo = c))
                                    : vs(t, e),
                                ys(e),
                                8192 & r)
                            ) {
                                if (
                                    ((c = null !== e.memoizedState),
                                    (e.stateNode.isHidden = c) &&
                                        !d &&
                                        0 !== (1 & e.mode))
                                )
                                    for (Jo = e, d = e.child; null !== d; ) {
                                        for (f = Jo = d; null !== Jo; ) {
                                            switch (
                                                ((m = (p = Jo).child), p.tag)
                                            ) {
                                                case 0:
                                                case 11:
                                                case 14:
                                                case 15:
                                                    rs(4, p, p.return);
                                                    break;
                                                case 1:
                                                    es(p, p.return);
                                                    var h = p.stateNode;
                                                    if (
                                                        "function" ===
                                                        typeof h.componentWillUnmount
                                                    ) {
                                                        (r = p), (n = p.return);
                                                        try {
                                                            (t = r),
                                                                (h.props =
                                                                    t.memoizedProps),
                                                                (h.state =
                                                                    t.memoizedState),
                                                                h.componentWillUnmount();
                                                        } catch (v) {
                                                            Nu(r, n, v);
                                                        }
                                                    }
                                                    break;
                                                case 5:
                                                    es(p, p.return);
                                                    break;
                                                case 22:
                                                    if (
                                                        null !== p.memoizedState
                                                    ) {
                                                        js(f);
                                                        continue;
                                                    }
                                            }
                                            null !== m
                                                ? ((m.return = p), (Jo = m))
                                                : js(f);
                                        }
                                        d = d.sibling;
                                    }
                                e: for (d = null, f = e; ; ) {
                                    if (5 === f.tag) {
                                        if (null === d) {
                                            d = f;
                                            try {
                                                (a = f.stateNode),
                                                    c
                                                        ? "function" ===
                                                          typeof (i = a.style)
                                                              .setProperty
                                                            ? i.setProperty(
                                                                  "display",
                                                                  "none",
                                                                  "important",
                                                              )
                                                            : (i.display =
                                                                  "none")
                                                        : ((s = f.stateNode),
                                                          (o =
                                                              void 0 !==
                                                                  (u =
                                                                      f
                                                                          .memoizedProps
                                                                          .style) &&
                                                              null !== u &&
                                                              u.hasOwnProperty(
                                                                  "display",
                                                              )
                                                                  ? u.display
                                                                  : null),
                                                          (s.style.display = he(
                                                              "display",
                                                              o,
                                                          )));
                                            } catch (v) {
                                                Nu(e, e.return, v);
                                            }
                                        }
                                    } else if (6 === f.tag) {
                                        if (null === d)
                                            try {
                                                f.stateNode.nodeValue = c
                                                    ? ""
                                                    : f.memoizedProps;
                                            } catch (v) {
                                                Nu(e, e.return, v);
                                            }
                                    } else if (
                                        ((22 !== f.tag && 23 !== f.tag) ||
                                            null === f.memoizedState ||
                                            f === e) &&
                                        null !== f.child
                                    ) {
                                        (f.child.return = f), (f = f.child);
                                        continue;
                                    }
                                    if (f === e) break e;
                                    for (; null === f.sibling; ) {
                                        if (null === f.return || f.return === e)
                                            break e;
                                        d === f && (d = null), (f = f.return);
                                    }
                                    d === f && (d = null),
                                        (f.sibling.return = f.return),
                                        (f = f.sibling);
                                }
                            }
                            break;
                        case 19:
                            vs(t, e), ys(e), 4 & r && hs(e);
                        case 21:
                    }
                }
                function ys(e) {
                    var t = e.flags;
                    if (2 & t) {
                        try {
                            e: {
                                for (var n = e.return; null !== n; ) {
                                    if (os(n)) {
                                        var r = n;
                                        break e;
                                    }
                                    n = n.return;
                                }
                                throw Error(l(160));
                            }
                            switch (r.tag) {
                                case 5:
                                    var a = r.stateNode;
                                    32 & r.flags &&
                                        (fe(a, ""), (r.flags &= -33)),
                                        cs(e, ss(e), a);
                                    break;
                                case 3:
                                case 4:
                                    var i = r.stateNode.containerInfo;
                                    us(e, ss(e), i);
                                    break;
                                default:
                                    throw Error(l(161));
                            }
                        } catch (o) {
                            Nu(e, e.return, o);
                        }
                        e.flags &= -3;
                    }
                    4096 & t && (e.flags &= -4097);
                }
                function bs(e, t, n) {
                    (Jo = e), xs(e, t, n);
                }
                function xs(e, t, n) {
                    for (var r = 0 !== (1 & e.mode); null !== Jo; ) {
                        var a = Jo,
                            l = a.child;
                        if (22 === a.tag && r) {
                            var i = null !== a.memoizedState || Yo;
                            if (!i) {
                                var o = a.alternate,
                                    s =
                                        (null !== o &&
                                            null !== o.memoizedState) ||
                                        Xo;
                                o = Yo;
                                var u = Xo;
                                if (((Yo = i), (Xo = s) && !u))
                                    for (Jo = a; null !== Jo; )
                                        (s = (i = Jo).child),
                                            22 === i.tag &&
                                            null !== i.memoizedState
                                                ? ks(a)
                                                : null !== s
                                                ? ((s.return = i), (Jo = s))
                                                : ks(a);
                                for (; null !== l; )
                                    (Jo = l), xs(l, t, n), (l = l.sibling);
                                (Jo = a), (Yo = o), (Xo = u);
                            }
                            ws(e);
                        } else
                            0 !== (8772 & a.subtreeFlags) && null !== l
                                ? ((l.return = a), (Jo = l))
                                : ws(e);
                    }
                }
                function ws(e) {
                    for (; null !== Jo; ) {
                        var t = Jo;
                        if (0 !== (8772 & t.flags)) {
                            var n = t.alternate;
                            try {
                                if (0 !== (8772 & t.flags))
                                    switch (t.tag) {
                                        case 0:
                                        case 11:
                                        case 15:
                                            Xo || as(5, t);
                                            break;
                                        case 1:
                                            var r = t.stateNode;
                                            if (4 & t.flags && !Xo)
                                                if (null === n)
                                                    r.componentDidMount();
                                                else {
                                                    var a =
                                                        t.elementType === t.type
                                                            ? n.memoizedProps
                                                            : vl(
                                                                  t.type,
                                                                  n.memoizedProps,
                                                              );
                                                    r.componentDidUpdate(
                                                        a,
                                                        n.memoizedState,
                                                        r.__reactInternalSnapshotBeforeUpdate,
                                                    );
                                                }
                                            var i = t.updateQueue;
                                            null !== i && Dl(t, i, r);
                                            break;
                                        case 3:
                                            var o = t.updateQueue;
                                            if (null !== o) {
                                                if (
                                                    ((n = null),
                                                    null !== t.child)
                                                )
                                                    switch (t.child.tag) {
                                                        case 5:
                                                        case 1:
                                                            n =
                                                                t.child
                                                                    .stateNode;
                                                    }
                                                Dl(t, o, n);
                                            }
                                            break;
                                        case 5:
                                            var s = t.stateNode;
                                            if (null === n && 4 & t.flags) {
                                                n = s;
                                                var u = t.memoizedProps;
                                                switch (t.type) {
                                                    case "button":
                                                    case "input":
                                                    case "select":
                                                    case "textarea":
                                                        u.autoFocus &&
                                                            n.focus();
                                                        break;
                                                    case "img":
                                                        u.src &&
                                                            (n.src = u.src);
                                                }
                                            }
                                            break;
                                        case 6:
                                        case 4:
                                        case 12:
                                        case 19:
                                        case 17:
                                        case 21:
                                        case 22:
                                        case 23:
                                        case 25:
                                            break;
                                        case 13:
                                            if (null === t.memoizedState) {
                                                var c = t.alternate;
                                                if (null !== c) {
                                                    var d = c.memoizedState;
                                                    if (null !== d) {
                                                        var f = d.dehydrated;
                                                        null !== f && Vt(f);
                                                    }
                                                }
                                            }
                                            break;
                                        default:
                                            throw Error(l(163));
                                    }
                                Xo || (512 & t.flags && ls(t));
                            } catch (p) {
                                Nu(t, t.return, p);
                            }
                        }
                        if (t === e) {
                            Jo = null;
                            break;
                        }
                        if (null !== (n = t.sibling)) {
                            (n.return = t.return), (Jo = n);
                            break;
                        }
                        Jo = t.return;
                    }
                }
                function js(e) {
                    for (; null !== Jo; ) {
                        var t = Jo;
                        if (t === e) {
                            Jo = null;
                            break;
                        }
                        var n = t.sibling;
                        if (null !== n) {
                            (n.return = t.return), (Jo = n);
                            break;
                        }
                        Jo = t.return;
                    }
                }
                function ks(e) {
                    for (; null !== Jo; ) {
                        var t = Jo;
                        try {
                            switch (t.tag) {
                                case 0:
                                case 11:
                                case 15:
                                    var n = t.return;
                                    try {
                                        as(4, t);
                                    } catch (s) {
                                        Nu(t, n, s);
                                    }
                                    break;
                                case 1:
                                    var r = t.stateNode;
                                    if (
                                        "function" ===
                                        typeof r.componentDidMount
                                    ) {
                                        var a = t.return;
                                        try {
                                            r.componentDidMount();
                                        } catch (s) {
                                            Nu(t, a, s);
                                        }
                                    }
                                    var l = t.return;
                                    try {
                                        ls(t);
                                    } catch (s) {
                                        Nu(t, l, s);
                                    }
                                    break;
                                case 5:
                                    var i = t.return;
                                    try {
                                        ls(t);
                                    } catch (s) {
                                        Nu(t, i, s);
                                    }
                            }
                        } catch (s) {
                            Nu(t, t.return, s);
                        }
                        if (t === e) {
                            Jo = null;
                            break;
                        }
                        var o = t.sibling;
                        if (null !== o) {
                            (o.return = t.return), (Jo = o);
                            break;
                        }
                        Jo = t.return;
                    }
                }
                var Ss,
                    Ns = Math.ceil,
                    Cs = x.ReactCurrentDispatcher,
                    Es = x.ReactCurrentOwner,
                    Ls = x.ReactCurrentBatchConfig,
                    _s = 0,
                    Ps = null,
                    Os = null,
                    zs = 0,
                    Ms = 0,
                    Rs = Sa(0),
                    Ts = 0,
                    Fs = null,
                    Is = 0,
                    Ds = 0,
                    Us = 0,
                    Bs = null,
                    As = null,
                    Vs = 0,
                    $s = 1 / 0,
                    Hs = null,
                    Ks = !1,
                    Ws = null,
                    Qs = null,
                    qs = !1,
                    Gs = null,
                    Ys = 0,
                    Xs = 0,
                    Zs = null,
                    Js = -1,
                    eu = 0;
                function tu() {
                    return 0 !== (6 & _s) ? Xe() : -1 !== Js ? Js : (Js = Xe());
                }
                function nu(e) {
                    return 0 === (1 & e.mode)
                        ? 1
                        : 0 !== (2 & _s) && 0 !== zs
                        ? zs & -zs
                        : null !== hl.transition
                        ? (0 === eu && (eu = ht()), eu)
                        : 0 !== (e = bt)
                        ? e
                        : (e = void 0 === (e = window.event) ? 16 : Yt(e.type));
                }
                function ru(e, t, n, r) {
                    if (50 < Xs) throw ((Xs = 0), (Zs = null), Error(l(185)));
                    gt(e, n, r),
                        (0 !== (2 & _s) && e === Ps) ||
                            (e === Ps &&
                                (0 === (2 & _s) && (Ds |= n),
                                4 === Ts && su(e, zs)),
                            au(e, r),
                            1 === n &&
                                0 === _s &&
                                0 === (1 & t.mode) &&
                                (($s = Xe() + 500), Ua && Va()));
                }
                function au(e, t) {
                    var n = e.callbackNode;
                    !(function (e, t) {
                        for (
                            var n = e.suspendedLanes,
                                r = e.pingedLanes,
                                a = e.expirationTimes,
                                l = e.pendingLanes;
                            0 < l;

                        ) {
                            var i = 31 - it(l),
                                o = 1 << i,
                                s = a[i];
                            -1 === s
                                ? (0 !== (o & n) && 0 === (o & r)) ||
                                  (a[i] = pt(o, t))
                                : s <= t && (e.expiredLanes |= o),
                                (l &= ~o);
                        }
                    })(e, t);
                    var r = ft(e, e === Ps ? zs : 0);
                    if (0 === r)
                        null !== n && qe(n),
                            (e.callbackNode = null),
                            (e.callbackPriority = 0);
                    else if (((t = r & -r), e.callbackPriority !== t)) {
                        if ((null != n && qe(n), 1 === t))
                            0 === e.tag
                                ? (function (e) {
                                      (Ua = !0), Aa(e);
                                  })(uu.bind(null, e))
                                : Aa(uu.bind(null, e)),
                                ia(function () {
                                    0 === (6 & _s) && Va();
                                }),
                                (n = null);
                        else {
                            switch (xt(r)) {
                                case 1:
                                    n = Je;
                                    break;
                                case 4:
                                    n = et;
                                    break;
                                case 16:
                                default:
                                    n = tt;
                                    break;
                                case 536870912:
                                    n = rt;
                            }
                            n = Pu(n, lu.bind(null, e));
                        }
                        (e.callbackPriority = t), (e.callbackNode = n);
                    }
                }
                function lu(e, t) {
                    if (((Js = -1), (eu = 0), 0 !== (6 & _s)))
                        throw Error(l(327));
                    var n = e.callbackNode;
                    if (ku() && e.callbackNode !== n) return null;
                    var r = ft(e, e === Ps ? zs : 0);
                    if (0 === r) return null;
                    if (0 !== (30 & r) || 0 !== (r & e.expiredLanes) || t)
                        t = gu(e, r);
                    else {
                        t = r;
                        var a = _s;
                        _s |= 2;
                        var i = hu();
                        for (
                            (Ps === e && zs === t) ||
                            ((Hs = null), ($s = Xe() + 500), pu(e, t));
                            ;

                        )
                            try {
                                bu();
                                break;
                            } catch (s) {
                                mu(e, s);
                            }
                        wl(),
                            (Cs.current = i),
                            (_s = a),
                            null !== Os
                                ? (t = 0)
                                : ((Ps = null), (zs = 0), (t = Ts));
                    }
                    if (0 !== t) {
                        if (
                            (2 === t &&
                                0 !== (a = mt(e)) &&
                                ((r = a), (t = iu(e, a))),
                            1 === t)
                        )
                            throw (
                                ((n = Fs), pu(e, 0), su(e, r), au(e, Xe()), n)
                            );
                        if (6 === t) su(e, r);
                        else {
                            if (
                                ((a = e.current.alternate),
                                0 === (30 & r) &&
                                    !(function (e) {
                                        for (var t = e; ; ) {
                                            if (16384 & t.flags) {
                                                var n = t.updateQueue;
                                                if (
                                                    null !== n &&
                                                    null !== (n = n.stores)
                                                )
                                                    for (
                                                        var r = 0;
                                                        r < n.length;
                                                        r++
                                                    ) {
                                                        var a = n[r],
                                                            l = a.getSnapshot;
                                                        a = a.value;
                                                        try {
                                                            if (!or(l(), a))
                                                                return !1;
                                                        } catch (o) {
                                                            return !1;
                                                        }
                                                    }
                                            }
                                            if (
                                                ((n = t.child),
                                                16384 & t.subtreeFlags &&
                                                    null !== n)
                                            )
                                                (n.return = t), (t = n);
                                            else {
                                                if (t === e) break;
                                                for (; null === t.sibling; ) {
                                                    if (
                                                        null === t.return ||
                                                        t.return === e
                                                    )
                                                        return !0;
                                                    t = t.return;
                                                }
                                                (t.sibling.return = t.return),
                                                    (t = t.sibling);
                                            }
                                        }
                                        return !0;
                                    })(a) &&
                                    (2 === (t = gu(e, r)) &&
                                        0 !== (i = mt(e)) &&
                                        ((r = i), (t = iu(e, i))),
                                    1 === t))
                            )
                                throw (
                                    ((n = Fs),
                                    pu(e, 0),
                                    su(e, r),
                                    au(e, Xe()),
                                    n)
                                );
                            switch (
                                ((e.finishedWork = a), (e.finishedLanes = r), t)
                            ) {
                                case 0:
                                case 1:
                                    throw Error(l(345));
                                case 2:
                                case 5:
                                    ju(e, As, Hs);
                                    break;
                                case 3:
                                    if (
                                        (su(e, r),
                                        (130023424 & r) === r &&
                                            10 < (t = Vs + 500 - Xe()))
                                    ) {
                                        if (0 !== ft(e, 0)) break;
                                        if (
                                            ((a = e.suspendedLanes) & r) !==
                                            r
                                        ) {
                                            tu(),
                                                (e.pingedLanes |=
                                                    e.suspendedLanes & a);
                                            break;
                                        }
                                        e.timeoutHandle = ra(
                                            ju.bind(null, e, As, Hs),
                                            t,
                                        );
                                        break;
                                    }
                                    ju(e, As, Hs);
                                    break;
                                case 4:
                                    if ((su(e, r), (4194240 & r) === r)) break;
                                    for (t = e.eventTimes, a = -1; 0 < r; ) {
                                        var o = 31 - it(r);
                                        (i = 1 << o),
                                            (o = t[o]) > a && (a = o),
                                            (r &= ~i);
                                    }
                                    if (
                                        ((r = a),
                                        10 <
                                            (r =
                                                (120 > (r = Xe() - r)
                                                    ? 120
                                                    : 480 > r
                                                    ? 480
                                                    : 1080 > r
                                                    ? 1080
                                                    : 1920 > r
                                                    ? 1920
                                                    : 3e3 > r
                                                    ? 3e3
                                                    : 4320 > r
                                                    ? 4320
                                                    : 1960 * Ns(r / 1960)) - r))
                                    ) {
                                        e.timeoutHandle = ra(
                                            ju.bind(null, e, As, Hs),
                                            r,
                                        );
                                        break;
                                    }
                                    ju(e, As, Hs);
                                    break;
                                default:
                                    throw Error(l(329));
                            }
                        }
                    }
                    return (
                        au(e, Xe()),
                        e.callbackNode === n ? lu.bind(null, e) : null
                    );
                }
                function iu(e, t) {
                    var n = Bs;
                    return (
                        e.current.memoizedState.isDehydrated &&
                            (pu(e, t).flags |= 256),
                        2 !== (e = gu(e, t)) &&
                            ((t = As), (As = n), null !== t && ou(t)),
                        e
                    );
                }
                function ou(e) {
                    null === As ? (As = e) : As.push.apply(As, e);
                }
                function su(e, t) {
                    for (
                        t &= ~Us,
                            t &= ~Ds,
                            e.suspendedLanes |= t,
                            e.pingedLanes &= ~t,
                            e = e.expirationTimes;
                        0 < t;

                    ) {
                        var n = 31 - it(t),
                            r = 1 << n;
                        (e[n] = -1), (t &= ~r);
                    }
                }
                function uu(e) {
                    if (0 !== (6 & _s)) throw Error(l(327));
                    ku();
                    var t = ft(e, 0);
                    if (0 === (1 & t)) return au(e, Xe()), null;
                    var n = gu(e, t);
                    if (0 !== e.tag && 2 === n) {
                        var r = mt(e);
                        0 !== r && ((t = r), (n = iu(e, r)));
                    }
                    if (1 === n)
                        throw ((n = Fs), pu(e, 0), su(e, t), au(e, Xe()), n);
                    if (6 === n) throw Error(l(345));
                    return (
                        (e.finishedWork = e.current.alternate),
                        (e.finishedLanes = t),
                        ju(e, As, Hs),
                        au(e, Xe()),
                        null
                    );
                }
                function cu(e, t) {
                    var n = _s;
                    _s |= 1;
                    try {
                        return e(t);
                    } finally {
                        0 === (_s = n) && (($s = Xe() + 500), Ua && Va());
                    }
                }
                function du(e) {
                    null !== Gs && 0 === Gs.tag && 0 === (6 & _s) && ku();
                    var t = _s;
                    _s |= 1;
                    var n = Ls.transition,
                        r = bt;
                    try {
                        if (((Ls.transition = null), (bt = 1), e)) return e();
                    } finally {
                        (bt = r),
                            (Ls.transition = n),
                            0 === (6 & (_s = t)) && Va();
                    }
                }
                function fu() {
                    (Ms = Rs.current), Na(Rs);
                }
                function pu(e, t) {
                    (e.finishedWork = null), (e.finishedLanes = 0);
                    var n = e.timeoutHandle;
                    if (
                        (-1 !== n && ((e.timeoutHandle = -1), aa(n)),
                        null !== Os)
                    )
                        for (n = Os.return; null !== n; ) {
                            var r = n;
                            switch ((tl(r), r.tag)) {
                                case 1:
                                    null !== (r = r.type.childContextTypes) &&
                                        void 0 !== r &&
                                        Ma();
                                    break;
                                case 3:
                                    ai(), Na(_a), Na(La), ci();
                                    break;
                                case 5:
                                    ii(r);
                                    break;
                                case 4:
                                    ai();
                                    break;
                                case 13:
                                case 19:
                                    Na(oi);
                                    break;
                                case 10:
                                    jl(r.type._context);
                                    break;
                                case 22:
                                case 23:
                                    fu();
                            }
                            n = n.return;
                        }
                    if (
                        ((Ps = e),
                        (Os = e = Ru(e.current, null)),
                        (zs = Ms = t),
                        (Ts = 0),
                        (Fs = null),
                        (Us = Ds = Is = 0),
                        (As = Bs = null),
                        null !== Cl)
                    ) {
                        for (t = 0; t < Cl.length; t++)
                            if (null !== (r = (n = Cl[t]).interleaved)) {
                                n.interleaved = null;
                                var a = r.next,
                                    l = n.pending;
                                if (null !== l) {
                                    var i = l.next;
                                    (l.next = a), (r.next = i);
                                }
                                n.pending = r;
                            }
                        Cl = null;
                    }
                    return e;
                }
                function mu(e, t) {
                    for (;;) {
                        var n = Os;
                        try {
                            if ((wl(), (di.current = lo), gi)) {
                                for (var r = mi.memoizedState; null !== r; ) {
                                    var a = r.queue;
                                    null !== a && (a.pending = null),
                                        (r = r.next);
                                }
                                gi = !1;
                            }
                            if (
                                ((pi = 0),
                                (vi = hi = mi = null),
                                (yi = !1),
                                (bi = 0),
                                (Es.current = null),
                                null === n || null === n.return)
                            ) {
                                (Ts = 1), (Fs = t), (Os = null);
                                break;
                            }
                            e: {
                                var i = e,
                                    o = n.return,
                                    s = n,
                                    u = t;
                                if (
                                    ((t = zs),
                                    (s.flags |= 32768),
                                    null !== u &&
                                        "object" === typeof u &&
                                        "function" === typeof u.then)
                                ) {
                                    var c = u,
                                        d = s,
                                        f = d.tag;
                                    if (
                                        0 === (1 & d.mode) &&
                                        (0 === f || 11 === f || 15 === f)
                                    ) {
                                        var p = d.alternate;
                                        p
                                            ? ((d.updateQueue = p.updateQueue),
                                              (d.memoizedState =
                                                  p.memoizedState),
                                              (d.lanes = p.lanes))
                                            : ((d.updateQueue = null),
                                              (d.memoizedState = null));
                                    }
                                    var m = go(o);
                                    if (null !== m) {
                                        (m.flags &= -257),
                                            yo(m, o, s, 0, t),
                                            1 & m.mode && vo(i, c, t),
                                            (u = c);
                                        var h = (t = m).updateQueue;
                                        if (null === h) {
                                            var v = new Set();
                                            v.add(u), (t.updateQueue = v);
                                        } else h.add(u);
                                        break e;
                                    }
                                    if (0 === (1 & t)) {
                                        vo(i, c, t), vu();
                                        break e;
                                    }
                                    u = Error(l(426));
                                } else if (al && 1 & s.mode) {
                                    var g = go(o);
                                    if (null !== g) {
                                        0 === (65536 & g.flags) &&
                                            (g.flags |= 256),
                                            yo(g, o, s, 0, t),
                                            ml(uo(u, s));
                                        break e;
                                    }
                                }
                                (i = u = uo(u, s)),
                                    4 !== Ts && (Ts = 2),
                                    null === Bs ? (Bs = [i]) : Bs.push(i),
                                    (i = o);
                                do {
                                    switch (i.tag) {
                                        case 3:
                                            (i.flags |= 65536),
                                                (t &= -t),
                                                (i.lanes |= t),
                                                Fl(i, mo(0, u, t));
                                            break e;
                                        case 1:
                                            s = u;
                                            var y = i.type,
                                                b = i.stateNode;
                                            if (
                                                0 === (128 & i.flags) &&
                                                ("function" ===
                                                    typeof y.getDerivedStateFromError ||
                                                    (null !== b &&
                                                        "function" ===
                                                            typeof b.componentDidCatch &&
                                                        (null === Qs ||
                                                            !Qs.has(b))))
                                            ) {
                                                (i.flags |= 65536),
                                                    (t &= -t),
                                                    (i.lanes |= t),
                                                    Fl(i, ho(i, s, t));
                                                break e;
                                            }
                                    }
                                    i = i.return;
                                } while (null !== i);
                            }
                            wu(n);
                        } catch (x) {
                            (t = x),
                                Os === n && null !== n && (Os = n = n.return);
                            continue;
                        }
                        break;
                    }
                }
                function hu() {
                    var e = Cs.current;
                    return (Cs.current = lo), null === e ? lo : e;
                }
                function vu() {
                    (0 !== Ts && 3 !== Ts && 2 !== Ts) || (Ts = 4),
                        null === Ps ||
                            (0 === (268435455 & Is) &&
                                0 === (268435455 & Ds)) ||
                            su(Ps, zs);
                }
                function gu(e, t) {
                    var n = _s;
                    _s |= 2;
                    var r = hu();
                    for ((Ps === e && zs === t) || ((Hs = null), pu(e, t)); ; )
                        try {
                            yu();
                            break;
                        } catch (a) {
                            mu(e, a);
                        }
                    if ((wl(), (_s = n), (Cs.current = r), null !== Os))
                        throw Error(l(261));
                    return (Ps = null), (zs = 0), Ts;
                }
                function yu() {
                    for (; null !== Os; ) xu(Os);
                }
                function bu() {
                    for (; null !== Os && !Ge(); ) xu(Os);
                }
                function xu(e) {
                    var t = Ss(e.alternate, e, Ms);
                    (e.memoizedProps = e.pendingProps),
                        null === t ? wu(e) : (Os = t),
                        (Es.current = null);
                }
                function wu(e) {
                    var t = e;
                    do {
                        var n = t.alternate;
                        if (((e = t.return), 0 === (32768 & t.flags))) {
                            if (null !== (n = qo(n, t, Ms)))
                                return void (Os = n);
                        } else {
                            if (null !== (n = Go(n, t)))
                                return (n.flags &= 32767), void (Os = n);
                            if (null === e) return (Ts = 6), void (Os = null);
                            (e.flags |= 32768),
                                (e.subtreeFlags = 0),
                                (e.deletions = null);
                        }
                        if (null !== (t = t.sibling)) return void (Os = t);
                        Os = t = e;
                    } while (null !== t);
                    0 === Ts && (Ts = 5);
                }
                function ju(e, t, n) {
                    var r = bt,
                        a = Ls.transition;
                    try {
                        (Ls.transition = null),
                            (bt = 1),
                            (function (e, t, n, r) {
                                do {
                                    ku();
                                } while (null !== Gs);
                                if (0 !== (6 & _s)) throw Error(l(327));
                                n = e.finishedWork;
                                var a = e.finishedLanes;
                                if (null === n) return null;
                                if (
                                    ((e.finishedWork = null),
                                    (e.finishedLanes = 0),
                                    n === e.current)
                                )
                                    throw Error(l(177));
                                (e.callbackNode = null),
                                    (e.callbackPriority = 0);
                                var i = n.lanes | n.childLanes;
                                if (
                                    ((function (e, t) {
                                        var n = e.pendingLanes & ~t;
                                        (e.pendingLanes = t),
                                            (e.suspendedLanes = 0),
                                            (e.pingedLanes = 0),
                                            (e.expiredLanes &= t),
                                            (e.mutableReadLanes &= t),
                                            (e.entangledLanes &= t),
                                            (t = e.entanglements);
                                        var r = e.eventTimes;
                                        for (e = e.expirationTimes; 0 < n; ) {
                                            var a = 31 - it(n),
                                                l = 1 << a;
                                            (t[a] = 0),
                                                (r[a] = -1),
                                                (e[a] = -1),
                                                (n &= ~l);
                                        }
                                    })(e, i),
                                    e === Ps && ((Os = Ps = null), (zs = 0)),
                                    (0 === (2064 & n.subtreeFlags) &&
                                        0 === (2064 & n.flags)) ||
                                        qs ||
                                        ((qs = !0),
                                        Pu(tt, function () {
                                            return ku(), null;
                                        })),
                                    (i = 0 !== (15990 & n.flags)),
                                    0 !== (15990 & n.subtreeFlags) || i)
                                ) {
                                    (i = Ls.transition), (Ls.transition = null);
                                    var o = bt;
                                    bt = 1;
                                    var s = _s;
                                    (_s |= 4),
                                        (Es.current = null),
                                        (function (e, t) {
                                            if (((ea = Ht), pr((e = fr())))) {
                                                if ("selectionStart" in e)
                                                    var n = {
                                                        start: e.selectionStart,
                                                        end: e.selectionEnd,
                                                    };
                                                else
                                                    e: {
                                                        var r =
                                                            (n =
                                                                ((n =
                                                                    e.ownerDocument) &&
                                                                    n.defaultView) ||
                                                                window)
                                                                .getSelection &&
                                                            n.getSelection();
                                                        if (
                                                            r &&
                                                            0 !== r.rangeCount
                                                        ) {
                                                            n = r.anchorNode;
                                                            var a =
                                                                    r.anchorOffset,
                                                                i = r.focusNode;
                                                            r = r.focusOffset;
                                                            try {
                                                                n.nodeType,
                                                                    i.nodeType;
                                                            } catch (w) {
                                                                n = null;
                                                                break e;
                                                            }
                                                            var o = 0,
                                                                s = -1,
                                                                u = -1,
                                                                c = 0,
                                                                d = 0,
                                                                f = e,
                                                                p = null;
                                                            t: for (;;) {
                                                                for (
                                                                    var m;
                                                                    f !== n ||
                                                                        (0 !==
                                                                            a &&
                                                                            3 !==
                                                                                f.nodeType) ||
                                                                        (s =
                                                                            o +
                                                                            a),
                                                                        f !==
                                                                            i ||
                                                                            (0 !==
                                                                                r &&
                                                                                3 !==
                                                                                    f.nodeType) ||
                                                                            (u =
                                                                                o +
                                                                                r),
                                                                        3 ===
                                                                            f.nodeType &&
                                                                            (o +=
                                                                                f
                                                                                    .nodeValue
                                                                                    .length),
                                                                        null !==
                                                                            (m =
                                                                                f.firstChild);

                                                                )
                                                                    (p = f),
                                                                        (f = m);
                                                                for (;;) {
                                                                    if (f === e)
                                                                        break t;
                                                                    if (
                                                                        (p ===
                                                                            n &&
                                                                            ++c ===
                                                                                a &&
                                                                            (s =
                                                                                o),
                                                                        p ===
                                                                            i &&
                                                                            ++d ===
                                                                                r &&
                                                                            (u =
                                                                                o),
                                                                        null !==
                                                                            (m =
                                                                                f.nextSibling))
                                                                    )
                                                                        break;
                                                                    p = (f = p)
                                                                        .parentNode;
                                                                }
                                                                f = m;
                                                            }
                                                            n =
                                                                -1 === s ||
                                                                -1 === u
                                                                    ? null
                                                                    : {
                                                                          start: s,
                                                                          end: u,
                                                                      };
                                                        } else n = null;
                                                    }
                                                n = n || { start: 0, end: 0 };
                                            } else n = null;
                                            for (
                                                ta = {
                                                    focusedElem: e,
                                                    selectionRange: n,
                                                },
                                                    Ht = !1,
                                                    Jo = t;
                                                null !== Jo;

                                            )
                                                if (
                                                    ((e = (t = Jo).child),
                                                    0 !==
                                                        (1028 &
                                                            t.subtreeFlags) &&
                                                        null !== e)
                                                )
                                                    (e.return = t), (Jo = e);
                                                else
                                                    for (; null !== Jo; ) {
                                                        t = Jo;
                                                        try {
                                                            var h = t.alternate;
                                                            if (
                                                                0 !==
                                                                (1024 & t.flags)
                                                            )
                                                                switch (t.tag) {
                                                                    case 0:
                                                                    case 11:
                                                                    case 15:
                                                                    case 5:
                                                                    case 6:
                                                                    case 4:
                                                                    case 17:
                                                                        break;
                                                                    case 1:
                                                                        if (
                                                                            null !==
                                                                            h
                                                                        ) {
                                                                            var v =
                                                                                    h.memoizedProps,
                                                                                g =
                                                                                    h.memoizedState,
                                                                                y =
                                                                                    t.stateNode,
                                                                                b =
                                                                                    y.getSnapshotBeforeUpdate(
                                                                                        t.elementType ===
                                                                                            t.type
                                                                                            ? v
                                                                                            : vl(
                                                                                                  t.type,
                                                                                                  v,
                                                                                              ),
                                                                                        g,
                                                                                    );
                                                                            y.__reactInternalSnapshotBeforeUpdate =
                                                                                b;
                                                                        }
                                                                        break;
                                                                    case 3:
                                                                        var x =
                                                                            t
                                                                                .stateNode
                                                                                .containerInfo;
                                                                        1 ===
                                                                        x.nodeType
                                                                            ? (x.textContent =
                                                                                  "")
                                                                            : 9 ===
                                                                                  x.nodeType &&
                                                                              x.documentElement &&
                                                                              x.removeChild(
                                                                                  x.documentElement,
                                                                              );
                                                                        break;
                                                                    default:
                                                                        throw Error(
                                                                            l(
                                                                                163,
                                                                            ),
                                                                        );
                                                                }
                                                        } catch (w) {
                                                            Nu(t, t.return, w);
                                                        }
                                                        if (
                                                            null !==
                                                            (e = t.sibling)
                                                        ) {
                                                            (e.return =
                                                                t.return),
                                                                (Jo = e);
                                                            break;
                                                        }
                                                        Jo = t.return;
                                                    }
                                            (h = ns), (ns = !1);
                                        })(e, n),
                                        gs(n, e),
                                        mr(ta),
                                        (Ht = !!ea),
                                        (ta = ea = null),
                                        (e.current = n),
                                        bs(n, e, a),
                                        Ye(),
                                        (_s = s),
                                        (bt = o),
                                        (Ls.transition = i);
                                } else e.current = n;
                                if (
                                    (qs && ((qs = !1), (Gs = e), (Ys = a)),
                                    (i = e.pendingLanes),
                                    0 === i && (Qs = null),
                                    (function (e) {
                                        if (
                                            lt &&
                                            "function" ===
                                                typeof lt.onCommitFiberRoot
                                        )
                                            try {
                                                lt.onCommitFiberRoot(
                                                    at,
                                                    e,
                                                    void 0,
                                                    128 ===
                                                        (128 & e.current.flags),
                                                );
                                            } catch (t) {}
                                    })(n.stateNode),
                                    au(e, Xe()),
                                    null !== t)
                                )
                                    for (
                                        r = e.onRecoverableError, n = 0;
                                        n < t.length;
                                        n++
                                    )
                                        (a = t[n]),
                                            r(a.value, {
                                                componentStack: a.stack,
                                                digest: a.digest,
                                            });
                                if (Ks)
                                    throw ((Ks = !1), (e = Ws), (Ws = null), e);
                                0 !== (1 & Ys) && 0 !== e.tag && ku(),
                                    (i = e.pendingLanes),
                                    0 !== (1 & i)
                                        ? e === Zs
                                            ? Xs++
                                            : ((Xs = 0), (Zs = e))
                                        : (Xs = 0),
                                    Va();
                            })(e, t, n, r);
                    } finally {
                        (Ls.transition = a), (bt = r);
                    }
                    return null;
                }
                function ku() {
                    if (null !== Gs) {
                        var e = xt(Ys),
                            t = Ls.transition,
                            n = bt;
                        try {
                            if (
                                ((Ls.transition = null),
                                (bt = 16 > e ? 16 : e),
                                null === Gs)
                            )
                                var r = !1;
                            else {
                                if (
                                    ((e = Gs),
                                    (Gs = null),
                                    (Ys = 0),
                                    0 !== (6 & _s))
                                )
                                    throw Error(l(331));
                                var a = _s;
                                for (_s |= 4, Jo = e.current; null !== Jo; ) {
                                    var i = Jo,
                                        o = i.child;
                                    if (0 !== (16 & Jo.flags)) {
                                        var s = i.deletions;
                                        if (null !== s) {
                                            for (var u = 0; u < s.length; u++) {
                                                var c = s[u];
                                                for (Jo = c; null !== Jo; ) {
                                                    var d = Jo;
                                                    switch (d.tag) {
                                                        case 0:
                                                        case 11:
                                                        case 15:
                                                            rs(8, d, i);
                                                    }
                                                    var f = d.child;
                                                    if (null !== f)
                                                        (f.return = d),
                                                            (Jo = f);
                                                    else
                                                        for (; null !== Jo; ) {
                                                            var p = (d = Jo)
                                                                    .sibling,
                                                                m = d.return;
                                                            if (
                                                                (is(d), d === c)
                                                            ) {
                                                                Jo = null;
                                                                break;
                                                            }
                                                            if (null !== p) {
                                                                (p.return = m),
                                                                    (Jo = p);
                                                                break;
                                                            }
                                                            Jo = m;
                                                        }
                                                }
                                            }
                                            var h = i.alternate;
                                            if (null !== h) {
                                                var v = h.child;
                                                if (null !== v) {
                                                    h.child = null;
                                                    do {
                                                        var g = v.sibling;
                                                        (v.sibling = null),
                                                            (v = g);
                                                    } while (null !== v);
                                                }
                                            }
                                            Jo = i;
                                        }
                                    }
                                    if (
                                        0 !== (2064 & i.subtreeFlags) &&
                                        null !== o
                                    )
                                        (o.return = i), (Jo = o);
                                    else
                                        e: for (; null !== Jo; ) {
                                            if (0 !== (2048 & (i = Jo).flags))
                                                switch (i.tag) {
                                                    case 0:
                                                    case 11:
                                                    case 15:
                                                        rs(9, i, i.return);
                                                }
                                            var y = i.sibling;
                                            if (null !== y) {
                                                (y.return = i.return), (Jo = y);
                                                break e;
                                            }
                                            Jo = i.return;
                                        }
                                }
                                var b = e.current;
                                for (Jo = b; null !== Jo; ) {
                                    var x = (o = Jo).child;
                                    if (
                                        0 !== (2064 & o.subtreeFlags) &&
                                        null !== x
                                    )
                                        (x.return = o), (Jo = x);
                                    else
                                        e: for (o = b; null !== Jo; ) {
                                            if (0 !== (2048 & (s = Jo).flags))
                                                try {
                                                    switch (s.tag) {
                                                        case 0:
                                                        case 11:
                                                        case 15:
                                                            as(9, s);
                                                    }
                                                } catch (j) {
                                                    Nu(s, s.return, j);
                                                }
                                            if (s === o) {
                                                Jo = null;
                                                break e;
                                            }
                                            var w = s.sibling;
                                            if (null !== w) {
                                                (w.return = s.return), (Jo = w);
                                                break e;
                                            }
                                            Jo = s.return;
                                        }
                                }
                                if (
                                    ((_s = a),
                                    Va(),
                                    lt &&
                                        "function" ===
                                            typeof lt.onPostCommitFiberRoot)
                                )
                                    try {
                                        lt.onPostCommitFiberRoot(at, e);
                                    } catch (j) {}
                                r = !0;
                            }
                            return r;
                        } finally {
                            (bt = n), (Ls.transition = t);
                        }
                    }
                    return !1;
                }
                function Su(e, t, n) {
                    (e = Rl(e, (t = mo(0, (t = uo(n, t)), 1)), 1)),
                        (t = tu()),
                        null !== e && (gt(e, 1, t), au(e, t));
                }
                function Nu(e, t, n) {
                    if (3 === e.tag) Su(e, e, n);
                    else
                        for (; null !== t; ) {
                            if (3 === t.tag) {
                                Su(t, e, n);
                                break;
                            }
                            if (1 === t.tag) {
                                var r = t.stateNode;
                                if (
                                    "function" ===
                                        typeof t.type
                                            .getDerivedStateFromError ||
                                    ("function" ===
                                        typeof r.componentDidCatch &&
                                        (null === Qs || !Qs.has(r)))
                                ) {
                                    (t = Rl(
                                        t,
                                        (e = ho(t, (e = uo(n, e)), 1)),
                                        1,
                                    )),
                                        (e = tu()),
                                        null !== t && (gt(t, 1, e), au(t, e));
                                    break;
                                }
                            }
                            t = t.return;
                        }
                }
                function Cu(e, t, n) {
                    var r = e.pingCache;
                    null !== r && r.delete(t),
                        (t = tu()),
                        (e.pingedLanes |= e.suspendedLanes & n),
                        Ps === e &&
                            (zs & n) === n &&
                            (4 === Ts ||
                            (3 === Ts &&
                                (130023424 & zs) === zs &&
                                500 > Xe() - Vs)
                                ? pu(e, 0)
                                : (Us |= n)),
                        au(e, t);
                }
                function Eu(e, t) {
                    0 === t &&
                        (0 === (1 & e.mode)
                            ? (t = 1)
                            : ((t = ct),
                              0 === (130023424 & (ct <<= 1)) &&
                                  (ct = 4194304)));
                    var n = tu();
                    null !== (e = _l(e, t)) && (gt(e, t, n), au(e, n));
                }
                function Lu(e) {
                    var t = e.memoizedState,
                        n = 0;
                    null !== t && (n = t.retryLane), Eu(e, n);
                }
                function _u(e, t) {
                    var n = 0;
                    switch (e.tag) {
                        case 13:
                            var r = e.stateNode,
                                a = e.memoizedState;
                            null !== a && (n = a.retryLane);
                            break;
                        case 19:
                            r = e.stateNode;
                            break;
                        default:
                            throw Error(l(314));
                    }
                    null !== r && r.delete(t), Eu(e, n);
                }
                function Pu(e, t) {
                    return Qe(e, t);
                }
                function Ou(e, t, n, r) {
                    (this.tag = e),
                        (this.key = n),
                        (this.sibling =
                            this.child =
                            this.return =
                            this.stateNode =
                            this.type =
                            this.elementType =
                                null),
                        (this.index = 0),
                        (this.ref = null),
                        (this.pendingProps = t),
                        (this.dependencies =
                            this.memoizedState =
                            this.updateQueue =
                            this.memoizedProps =
                                null),
                        (this.mode = r),
                        (this.subtreeFlags = this.flags = 0),
                        (this.deletions = null),
                        (this.childLanes = this.lanes = 0),
                        (this.alternate = null);
                }
                function zu(e, t, n, r) {
                    return new Ou(e, t, n, r);
                }
                function Mu(e) {
                    return !(!(e = e.prototype) || !e.isReactComponent);
                }
                function Ru(e, t) {
                    var n = e.alternate;
                    return (
                        null === n
                            ? (((n = zu(e.tag, t, e.key, e.mode)).elementType =
                                  e.elementType),
                              (n.type = e.type),
                              (n.stateNode = e.stateNode),
                              (n.alternate = e),
                              (e.alternate = n))
                            : ((n.pendingProps = t),
                              (n.type = e.type),
                              (n.flags = 0),
                              (n.subtreeFlags = 0),
                              (n.deletions = null)),
                        (n.flags = 14680064 & e.flags),
                        (n.childLanes = e.childLanes),
                        (n.lanes = e.lanes),
                        (n.child = e.child),
                        (n.memoizedProps = e.memoizedProps),
                        (n.memoizedState = e.memoizedState),
                        (n.updateQueue = e.updateQueue),
                        (t = e.dependencies),
                        (n.dependencies =
                            null === t
                                ? null
                                : {
                                      lanes: t.lanes,
                                      firstContext: t.firstContext,
                                  }),
                        (n.sibling = e.sibling),
                        (n.index = e.index),
                        (n.ref = e.ref),
                        n
                    );
                }
                function Tu(e, t, n, r, a, i) {
                    var o = 2;
                    if (((r = e), "function" === typeof e)) Mu(e) && (o = 1);
                    else if ("string" === typeof e) o = 5;
                    else
                        e: switch (e) {
                            case k:
                                return Fu(n.children, a, i, t);
                            case S:
                                (o = 8), (a |= 8);
                                break;
                            case N:
                                return (
                                    ((e = zu(12, n, t, 2 | a)).elementType = N),
                                    (e.lanes = i),
                                    e
                                );
                            case _:
                                return (
                                    ((e = zu(13, n, t, a)).elementType = _),
                                    (e.lanes = i),
                                    e
                                );
                            case P:
                                return (
                                    ((e = zu(19, n, t, a)).elementType = P),
                                    (e.lanes = i),
                                    e
                                );
                            case M:
                                return Iu(n, a, i, t);
                            default:
                                if ("object" === typeof e && null !== e)
                                    switch (e.$$typeof) {
                                        case C:
                                            o = 10;
                                            break e;
                                        case E:
                                            o = 9;
                                            break e;
                                        case L:
                                            o = 11;
                                            break e;
                                        case O:
                                            o = 14;
                                            break e;
                                        case z:
                                            (o = 16), (r = null);
                                            break e;
                                    }
                                throw Error(
                                    l(130, null == e ? e : typeof e, ""),
                                );
                        }
                    return (
                        ((t = zu(o, n, t, a)).elementType = e),
                        (t.type = r),
                        (t.lanes = i),
                        t
                    );
                }
                function Fu(e, t, n, r) {
                    return ((e = zu(7, e, r, t)).lanes = n), e;
                }
                function Iu(e, t, n, r) {
                    return (
                        ((e = zu(22, e, r, t)).elementType = M),
                        (e.lanes = n),
                        (e.stateNode = { isHidden: !1 }),
                        e
                    );
                }
                function Du(e, t, n) {
                    return ((e = zu(6, e, null, t)).lanes = n), e;
                }
                function Uu(e, t, n) {
                    return (
                        ((t = zu(
                            4,
                            null !== e.children ? e.children : [],
                            e.key,
                            t,
                        )).lanes = n),
                        (t.stateNode = {
                            containerInfo: e.containerInfo,
                            pendingChildren: null,
                            implementation: e.implementation,
                        }),
                        t
                    );
                }
                function Bu(e, t, n, r, a) {
                    (this.tag = t),
                        (this.containerInfo = e),
                        (this.finishedWork =
                            this.pingCache =
                            this.current =
                            this.pendingChildren =
                                null),
                        (this.timeoutHandle = -1),
                        (this.callbackNode =
                            this.pendingContext =
                            this.context =
                                null),
                        (this.callbackPriority = 0),
                        (this.eventTimes = vt(0)),
                        (this.expirationTimes = vt(-1)),
                        (this.entangledLanes =
                            this.finishedLanes =
                            this.mutableReadLanes =
                            this.expiredLanes =
                            this.pingedLanes =
                            this.suspendedLanes =
                            this.pendingLanes =
                                0),
                        (this.entanglements = vt(0)),
                        (this.identifierPrefix = r),
                        (this.onRecoverableError = a),
                        (this.mutableSourceEagerHydrationData = null);
                }
                function Au(e, t, n, r, a, l, i, o, s) {
                    return (
                        (e = new Bu(e, t, n, o, s)),
                        1 === t ? ((t = 1), !0 === l && (t |= 8)) : (t = 0),
                        (l = zu(3, null, null, t)),
                        (e.current = l),
                        (l.stateNode = e),
                        (l.memoizedState = {
                            element: r,
                            isDehydrated: n,
                            cache: null,
                            transitions: null,
                            pendingSuspenseBoundaries: null,
                        }),
                        Ol(l),
                        e
                    );
                }
                function Vu(e) {
                    if (!e) return Ea;
                    e: {
                        if (Ve((e = e._reactInternals)) !== e || 1 !== e.tag)
                            throw Error(l(170));
                        var t = e;
                        do {
                            switch (t.tag) {
                                case 3:
                                    t = t.stateNode.context;
                                    break e;
                                case 1:
                                    if (za(t.type)) {
                                        t =
                                            t.stateNode
                                                .__reactInternalMemoizedMergedChildContext;
                                        break e;
                                    }
                            }
                            t = t.return;
                        } while (null !== t);
                        throw Error(l(171));
                    }
                    if (1 === e.tag) {
                        var n = e.type;
                        if (za(n)) return Ta(e, n, t);
                    }
                    return t;
                }
                function $u(e, t, n, r, a, l, i, o, s) {
                    return (
                        ((e = Au(n, r, !0, e, 0, l, 0, o, s)).context =
                            Vu(null)),
                        (n = e.current),
                        ((l = Ml((r = tu()), (a = nu(n)))).callback =
                            void 0 !== t && null !== t ? t : null),
                        Rl(n, l, a),
                        (e.current.lanes = a),
                        gt(e, a, r),
                        au(e, r),
                        e
                    );
                }
                function Hu(e, t, n, r) {
                    var a = t.current,
                        l = tu(),
                        i = nu(a);
                    return (
                        (n = Vu(n)),
                        null === t.context
                            ? (t.context = n)
                            : (t.pendingContext = n),
                        ((t = Ml(l, i)).payload = { element: e }),
                        null !== (r = void 0 === r ? null : r) &&
                            (t.callback = r),
                        null !== (e = Rl(a, t, i)) &&
                            (ru(e, a, i, l), Tl(e, a, i)),
                        i
                    );
                }
                function Ku(e) {
                    return (e = e.current).child
                        ? (e.child.tag, e.child.stateNode)
                        : null;
                }
                function Wu(e, t) {
                    if (
                        null !== (e = e.memoizedState) &&
                        null !== e.dehydrated
                    ) {
                        var n = e.retryLane;
                        e.retryLane = 0 !== n && n < t ? n : t;
                    }
                }
                function Qu(e, t) {
                    Wu(e, t), (e = e.alternate) && Wu(e, t);
                }
                Ss = function (e, t, n) {
                    if (null !== e)
                        if (e.memoizedProps !== t.pendingProps || _a.current)
                            xo = !0;
                        else {
                            if (0 === (e.lanes & n) && 0 === (128 & t.flags))
                                return (
                                    (xo = !1),
                                    (function (e, t, n) {
                                        switch (t.tag) {
                                            case 3:
                                                Po(t), pl();
                                                break;
                                            case 5:
                                                li(t);
                                                break;
                                            case 1:
                                                za(t.type) && Fa(t);
                                                break;
                                            case 4:
                                                ri(
                                                    t,
                                                    t.stateNode.containerInfo,
                                                );
                                                break;
                                            case 10:
                                                var r = t.type._context,
                                                    a = t.memoizedProps.value;
                                                Ca(gl, r._currentValue),
                                                    (r._currentValue = a);
                                                break;
                                            case 13:
                                                if (
                                                    null !==
                                                    (r = t.memoizedState)
                                                )
                                                    return null !== r.dehydrated
                                                        ? (Ca(
                                                              oi,
                                                              1 & oi.current,
                                                          ),
                                                          (t.flags |= 128),
                                                          null)
                                                        : 0 !==
                                                          (n &
                                                              t.child
                                                                  .childLanes)
                                                        ? Do(e, t, n)
                                                        : (Ca(
                                                              oi,
                                                              1 & oi.current,
                                                          ),
                                                          null !==
                                                          (e = Ko(e, t, n))
                                                              ? e.sibling
                                                              : null);
                                                Ca(oi, 1 & oi.current);
                                                break;
                                            case 19:
                                                if (
                                                    ((r =
                                                        0 !==
                                                        (n & t.childLanes)),
                                                    0 !== (128 & e.flags))
                                                ) {
                                                    if (r) return $o(e, t, n);
                                                    t.flags |= 128;
                                                }
                                                if (
                                                    (null !==
                                                        (a = t.memoizedState) &&
                                                        ((a.rendering = null),
                                                        (a.tail = null),
                                                        (a.lastEffect = null)),
                                                    Ca(oi, oi.current),
                                                    r)
                                                )
                                                    break;
                                                return null;
                                            case 22:
                                            case 23:
                                                return (
                                                    (t.lanes = 0), No(e, t, n)
                                                );
                                        }
                                        return Ko(e, t, n);
                                    })(e, t, n)
                                );
                            xo = 0 !== (131072 & e.flags);
                        }
                    else
                        (xo = !1),
                            al &&
                                0 !== (1048576 & t.flags) &&
                                Ja(t, Wa, t.index);
                    switch (((t.lanes = 0), t.tag)) {
                        case 2:
                            var r = t.type;
                            Ho(e, t), (e = t.pendingProps);
                            var a = Oa(t, La.current);
                            Sl(t, n), (a = ki(null, t, r, e, a, n));
                            var i = Si();
                            return (
                                (t.flags |= 1),
                                "object" === typeof a &&
                                null !== a &&
                                "function" === typeof a.render &&
                                void 0 === a.$$typeof
                                    ? ((t.tag = 1),
                                      (t.memoizedState = null),
                                      (t.updateQueue = null),
                                      za(r) ? ((i = !0), Fa(t)) : (i = !1),
                                      (t.memoizedState =
                                          null !== a.state && void 0 !== a.state
                                              ? a.state
                                              : null),
                                      Ol(t),
                                      (a.updater = Al),
                                      (t.stateNode = a),
                                      (a._reactInternals = t),
                                      Kl(t, r, e, n),
                                      (t = _o(null, t, r, !0, i, n)))
                                    : ((t.tag = 0),
                                      al && i && el(t),
                                      wo(null, t, a, n),
                                      (t = t.child)),
                                t
                            );
                        case 16:
                            r = t.elementType;
                            e: {
                                switch (
                                    (Ho(e, t),
                                    (e = t.pendingProps),
                                    (r = (a = r._init)(r._payload)),
                                    (t.type = r),
                                    (a = t.tag =
                                        (function (e) {
                                            if ("function" === typeof e)
                                                return Mu(e) ? 1 : 0;
                                            if (void 0 !== e && null !== e) {
                                                if ((e = e.$$typeof) === L)
                                                    return 11;
                                                if (e === O) return 14;
                                            }
                                            return 2;
                                        })(r)),
                                    (e = vl(r, e)),
                                    a)
                                ) {
                                    case 0:
                                        t = Eo(null, t, r, e, n);
                                        break e;
                                    case 1:
                                        t = Lo(null, t, r, e, n);
                                        break e;
                                    case 11:
                                        t = jo(null, t, r, e, n);
                                        break e;
                                    case 14:
                                        t = ko(null, t, r, vl(r.type, e), n);
                                        break e;
                                }
                                throw Error(l(306, r, ""));
                            }
                            return t;
                        case 0:
                            return (
                                (r = t.type),
                                (a = t.pendingProps),
                                Eo(
                                    e,
                                    t,
                                    r,
                                    (a = t.elementType === r ? a : vl(r, a)),
                                    n,
                                )
                            );
                        case 1:
                            return (
                                (r = t.type),
                                (a = t.pendingProps),
                                Lo(
                                    e,
                                    t,
                                    r,
                                    (a = t.elementType === r ? a : vl(r, a)),
                                    n,
                                )
                            );
                        case 3:
                            e: {
                                if ((Po(t), null === e)) throw Error(l(387));
                                (r = t.pendingProps),
                                    (a = (i = t.memoizedState).element),
                                    zl(e, t),
                                    Il(t, r, null, n);
                                var o = t.memoizedState;
                                if (((r = o.element), i.isDehydrated)) {
                                    if (
                                        ((i = {
                                            element: r,
                                            isDehydrated: !1,
                                            cache: o.cache,
                                            pendingSuspenseBoundaries:
                                                o.pendingSuspenseBoundaries,
                                            transitions: o.transitions,
                                        }),
                                        (t.updateQueue.baseState = i),
                                        (t.memoizedState = i),
                                        256 & t.flags)
                                    ) {
                                        t = Oo(
                                            e,
                                            t,
                                            r,
                                            n,
                                            (a = uo(Error(l(423)), t)),
                                        );
                                        break e;
                                    }
                                    if (r !== a) {
                                        t = Oo(
                                            e,
                                            t,
                                            r,
                                            n,
                                            (a = uo(Error(l(424)), t)),
                                        );
                                        break e;
                                    }
                                    for (
                                        rl = ua(
                                            t.stateNode.containerInfo
                                                .firstChild,
                                        ),
                                            nl = t,
                                            al = !0,
                                            ll = null,
                                            n = Xl(t, null, r, n),
                                            t.child = n;
                                        n;

                                    )
                                        (n.flags = (-3 & n.flags) | 4096),
                                            (n = n.sibling);
                                } else {
                                    if ((pl(), r === a)) {
                                        t = Ko(e, t, n);
                                        break e;
                                    }
                                    wo(e, t, r, n);
                                }
                                t = t.child;
                            }
                            return t;
                        case 5:
                            return (
                                li(t),
                                null === e && ul(t),
                                (r = t.type),
                                (a = t.pendingProps),
                                (i = null !== e ? e.memoizedProps : null),
                                (o = a.children),
                                na(r, a)
                                    ? (o = null)
                                    : null !== i && na(r, i) && (t.flags |= 32),
                                Co(e, t),
                                wo(e, t, o, n),
                                t.child
                            );
                        case 6:
                            return null === e && ul(t), null;
                        case 13:
                            return Do(e, t, n);
                        case 4:
                            return (
                                ri(t, t.stateNode.containerInfo),
                                (r = t.pendingProps),
                                null === e
                                    ? (t.child = Yl(t, null, r, n))
                                    : wo(e, t, r, n),
                                t.child
                            );
                        case 11:
                            return (
                                (r = t.type),
                                (a = t.pendingProps),
                                jo(
                                    e,
                                    t,
                                    r,
                                    (a = t.elementType === r ? a : vl(r, a)),
                                    n,
                                )
                            );
                        case 7:
                            return wo(e, t, t.pendingProps, n), t.child;
                        case 8:
                        case 12:
                            return (
                                wo(e, t, t.pendingProps.children, n), t.child
                            );
                        case 10:
                            e: {
                                if (
                                    ((r = t.type._context),
                                    (a = t.pendingProps),
                                    (i = t.memoizedProps),
                                    (o = a.value),
                                    Ca(gl, r._currentValue),
                                    (r._currentValue = o),
                                    null !== i)
                                )
                                    if (or(i.value, o)) {
                                        if (
                                            i.children === a.children &&
                                            !_a.current
                                        ) {
                                            t = Ko(e, t, n);
                                            break e;
                                        }
                                    } else
                                        for (
                                            null !== (i = t.child) &&
                                            (i.return = t);
                                            null !== i;

                                        ) {
                                            var s = i.dependencies;
                                            if (null !== s) {
                                                o = i.child;
                                                for (
                                                    var u = s.firstContext;
                                                    null !== u;

                                                ) {
                                                    if (u.context === r) {
                                                        if (1 === i.tag) {
                                                            (u = Ml(
                                                                -1,
                                                                n & -n,
                                                            )).tag = 2;
                                                            var c =
                                                                i.updateQueue;
                                                            if (null !== c) {
                                                                var d = (c =
                                                                    c.shared)
                                                                    .pending;
                                                                null === d
                                                                    ? (u.next =
                                                                          u)
                                                                    : ((u.next =
                                                                          d.next),
                                                                      (d.next =
                                                                          u)),
                                                                    (c.pending =
                                                                        u);
                                                            }
                                                        }
                                                        (i.lanes |= n),
                                                            null !==
                                                                (u =
                                                                    i.alternate) &&
                                                                (u.lanes |= n),
                                                            kl(i.return, n, t),
                                                            (s.lanes |= n);
                                                        break;
                                                    }
                                                    u = u.next;
                                                }
                                            } else if (10 === i.tag)
                                                o =
                                                    i.type === t.type
                                                        ? null
                                                        : i.child;
                                            else if (18 === i.tag) {
                                                if (null === (o = i.return))
                                                    throw Error(l(341));
                                                (o.lanes |= n),
                                                    null !==
                                                        (s = o.alternate) &&
                                                        (s.lanes |= n),
                                                    kl(o, n, t),
                                                    (o = i.sibling);
                                            } else o = i.child;
                                            if (null !== o) o.return = i;
                                            else
                                                for (o = i; null !== o; ) {
                                                    if (o === t) {
                                                        o = null;
                                                        break;
                                                    }
                                                    if (
                                                        null !== (i = o.sibling)
                                                    ) {
                                                        (i.return = o.return),
                                                            (o = i);
                                                        break;
                                                    }
                                                    o = o.return;
                                                }
                                            i = o;
                                        }
                                wo(e, t, a.children, n), (t = t.child);
                            }
                            return t;
                        case 9:
                            return (
                                (a = t.type),
                                (r = t.pendingProps.children),
                                Sl(t, n),
                                (r = r((a = Nl(a)))),
                                (t.flags |= 1),
                                wo(e, t, r, n),
                                t.child
                            );
                        case 14:
                            return (
                                (a = vl((r = t.type), t.pendingProps)),
                                ko(e, t, r, (a = vl(r.type, a)), n)
                            );
                        case 15:
                            return So(e, t, t.type, t.pendingProps, n);
                        case 17:
                            return (
                                (r = t.type),
                                (a = t.pendingProps),
                                (a = t.elementType === r ? a : vl(r, a)),
                                Ho(e, t),
                                (t.tag = 1),
                                za(r) ? ((e = !0), Fa(t)) : (e = !1),
                                Sl(t, n),
                                $l(t, r, a),
                                Kl(t, r, a, n),
                                _o(null, t, r, !0, e, n)
                            );
                        case 19:
                            return $o(e, t, n);
                        case 22:
                            return No(e, t, n);
                    }
                    throw Error(l(156, t.tag));
                };
                var qu =
                    "function" === typeof reportError
                        ? reportError
                        : function (e) {
                              console.error(e);
                          };
                function Gu(e) {
                    this._internalRoot = e;
                }
                function Yu(e) {
                    this._internalRoot = e;
                }
                function Xu(e) {
                    return !(
                        !e ||
                        (1 !== e.nodeType &&
                            9 !== e.nodeType &&
                            11 !== e.nodeType)
                    );
                }
                function Zu(e) {
                    return !(
                        !e ||
                        (1 !== e.nodeType &&
                            9 !== e.nodeType &&
                            11 !== e.nodeType &&
                            (8 !== e.nodeType ||
                                " react-mount-point-unstable " !== e.nodeValue))
                    );
                }
                function Ju() {}
                function ec(e, t, n, r, a) {
                    var l = n._reactRootContainer;
                    if (l) {
                        var i = l;
                        if ("function" === typeof a) {
                            var o = a;
                            a = function () {
                                var e = Ku(i);
                                o.call(e);
                            };
                        }
                        Hu(t, i, e, a);
                    } else
                        i = (function (e, t, n, r, a) {
                            if (a) {
                                if ("function" === typeof r) {
                                    var l = r;
                                    r = function () {
                                        var e = Ku(i);
                                        l.call(e);
                                    };
                                }
                                var i = $u(t, r, e, 0, null, !1, 0, "", Ju);
                                return (
                                    (e._reactRootContainer = i),
                                    (e[ma] = i.current),
                                    Vr(8 === e.nodeType ? e.parentNode : e),
                                    du(),
                                    i
                                );
                            }
                            for (; (a = e.lastChild); ) e.removeChild(a);
                            if ("function" === typeof r) {
                                var o = r;
                                r = function () {
                                    var e = Ku(s);
                                    o.call(e);
                                };
                            }
                            var s = Au(e, 0, !1, null, 0, !1, 0, "", Ju);
                            return (
                                (e._reactRootContainer = s),
                                (e[ma] = s.current),
                                Vr(8 === e.nodeType ? e.parentNode : e),
                                du(function () {
                                    Hu(t, s, n, r);
                                }),
                                s
                            );
                        })(n, t, e, a, r);
                    return Ku(i);
                }
                (Yu.prototype.render = Gu.prototype.render =
                    function (e) {
                        var t = this._internalRoot;
                        if (null === t) throw Error(l(409));
                        Hu(e, t, null, null);
                    }),
                    (Yu.prototype.unmount = Gu.prototype.unmount =
                        function () {
                            var e = this._internalRoot;
                            if (null !== e) {
                                this._internalRoot = null;
                                var t = e.containerInfo;
                                du(function () {
                                    Hu(null, e, null, null);
                                }),
                                    (t[ma] = null);
                            }
                        }),
                    (Yu.prototype.unstable_scheduleHydration = function (e) {
                        if (e) {
                            var t = St();
                            e = { blockedOn: null, target: e, priority: t };
                            for (
                                var n = 0;
                                n < Mt.length && 0 !== t && t < Mt[n].priority;
                                n++
                            );
                            Mt.splice(n, 0, e), 0 === n && It(e);
                        }
                    }),
                    (wt = function (e) {
                        switch (e.tag) {
                            case 3:
                                var t = e.stateNode;
                                if (t.current.memoizedState.isDehydrated) {
                                    var n = dt(t.pendingLanes);
                                    0 !== n &&
                                        (yt(t, 1 | n),
                                        au(t, Xe()),
                                        0 === (6 & _s) &&
                                            (($s = Xe() + 500), Va()));
                                }
                                break;
                            case 13:
                                du(function () {
                                    var t = _l(e, 1);
                                    if (null !== t) {
                                        var n = tu();
                                        ru(t, e, 1, n);
                                    }
                                }),
                                    Qu(e, 1);
                        }
                    }),
                    (jt = function (e) {
                        if (13 === e.tag) {
                            var t = _l(e, 134217728);
                            if (null !== t) ru(t, e, 134217728, tu());
                            Qu(e, 134217728);
                        }
                    }),
                    (kt = function (e) {
                        if (13 === e.tag) {
                            var t = nu(e),
                                n = _l(e, t);
                            if (null !== n) ru(n, e, t, tu());
                            Qu(e, t);
                        }
                    }),
                    (St = function () {
                        return bt;
                    }),
                    (Nt = function (e, t) {
                        var n = bt;
                        try {
                            return (bt = e), t();
                        } finally {
                            bt = n;
                        }
                    }),
                    (je = function (e, t, n) {
                        switch (t) {
                            case "input":
                                if (
                                    (Z(e, n),
                                    (t = n.name),
                                    "radio" === n.type && null != t)
                                ) {
                                    for (n = e; n.parentNode; )
                                        n = n.parentNode;
                                    for (
                                        n = n.querySelectorAll(
                                            "input[name=" +
                                                JSON.stringify("" + t) +
                                                '][type="radio"]',
                                        ),
                                            t = 0;
                                        t < n.length;
                                        t++
                                    ) {
                                        var r = n[t];
                                        if (r !== e && r.form === e.form) {
                                            var a = wa(r);
                                            if (!a) throw Error(l(90));
                                            Q(r), Z(r, a);
                                        }
                                    }
                                }
                                break;
                            case "textarea":
                                le(e, n);
                                break;
                            case "select":
                                null != (t = n.value) &&
                                    ne(e, !!n.multiple, t, !1);
                        }
                    }),
                    (Le = cu),
                    (_e = du);
                var tc = {
                        usingClientEntryPoint: !1,
                        Events: [ba, xa, wa, Ce, Ee, cu],
                    },
                    nc = {
                        findFiberByHostInstance: ya,
                        bundleType: 0,
                        version: "18.2.0",
                        rendererPackageName: "react-dom",
                    },
                    rc = {
                        bundleType: nc.bundleType,
                        version: nc.version,
                        rendererPackageName: nc.rendererPackageName,
                        rendererConfig: nc.rendererConfig,
                        overrideHookState: null,
                        overrideHookStateDeletePath: null,
                        overrideHookStateRenamePath: null,
                        overrideProps: null,
                        overridePropsDeletePath: null,
                        overridePropsRenamePath: null,
                        setErrorHandler: null,
                        setSuspenseHandler: null,
                        scheduleUpdate: null,
                        currentDispatcherRef: x.ReactCurrentDispatcher,
                        findHostInstanceByFiber: function (e) {
                            return null === (e = Ke(e)) ? null : e.stateNode;
                        },
                        findFiberByHostInstance:
                            nc.findFiberByHostInstance ||
                            function () {
                                return null;
                            },
                        findHostInstancesForRefresh: null,
                        scheduleRefresh: null,
                        scheduleRoot: null,
                        setRefreshHandler: null,
                        getCurrentFiber: null,
                        reconcilerVersion: "18.2.0-next-9e3b772b8-20220608",
                    };
                if ("undefined" !== typeof __REACT_DEVTOOLS_GLOBAL_HOOK__) {
                    var ac = __REACT_DEVTOOLS_GLOBAL_HOOK__;
                    if (!ac.isDisabled && ac.supportsFiber)
                        try {
                            (at = ac.inject(rc)), (lt = ac);
                        } catch (ce) {}
                }
                (t.__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED = tc),
                    (t.createPortal = function (e, t) {
                        var n =
                            2 < arguments.length && void 0 !== arguments[2]
                                ? arguments[2]
                                : null;
                        if (!Xu(t)) throw Error(l(200));
                        return (function (e, t, n) {
                            var r =
                                3 < arguments.length && void 0 !== arguments[3]
                                    ? arguments[3]
                                    : null;
                            return {
                                $$typeof: j,
                                key: null == r ? null : "" + r,
                                children: e,
                                containerInfo: t,
                                implementation: n,
                            };
                        })(e, t, null, n);
                    }),
                    (t.createRoot = function (e, t) {
                        if (!Xu(e)) throw Error(l(299));
                        var n = !1,
                            r = "",
                            a = qu;
                        return (
                            null !== t &&
                                void 0 !== t &&
                                (!0 === t.unstable_strictMode && (n = !0),
                                void 0 !== t.identifierPrefix &&
                                    (r = t.identifierPrefix),
                                void 0 !== t.onRecoverableError &&
                                    (a = t.onRecoverableError)),
                            (t = Au(e, 1, !1, null, 0, n, 0, r, a)),
                            (e[ma] = t.current),
                            Vr(8 === e.nodeType ? e.parentNode : e),
                            new Gu(t)
                        );
                    }),
                    (t.findDOMNode = function (e) {
                        if (null == e) return null;
                        if (1 === e.nodeType) return e;
                        var t = e._reactInternals;
                        if (void 0 === t) {
                            if ("function" === typeof e.render)
                                throw Error(l(188));
                            throw (
                                ((e = Object.keys(e).join(",")),
                                Error(l(268, e)))
                            );
                        }
                        return (e = null === (e = Ke(t)) ? null : e.stateNode);
                    }),
                    (t.flushSync = function (e) {
                        return du(e);
                    }),
                    (t.hydrate = function (e, t, n) {
                        if (!Zu(t)) throw Error(l(200));
                        return ec(null, e, t, !0, n);
                    }),
                    (t.hydrateRoot = function (e, t, n) {
                        if (!Xu(e)) throw Error(l(405));
                        var r = (null != n && n.hydratedSources) || null,
                            a = !1,
                            i = "",
                            o = qu;
                        if (
                            (null !== n &&
                                void 0 !== n &&
                                (!0 === n.unstable_strictMode && (a = !0),
                                void 0 !== n.identifierPrefix &&
                                    (i = n.identifierPrefix),
                                void 0 !== n.onRecoverableError &&
                                    (o = n.onRecoverableError)),
                            (t = $u(
                                t,
                                null,
                                e,
                                1,
                                null != n ? n : null,
                                a,
                                0,
                                i,
                                o,
                            )),
                            (e[ma] = t.current),
                            Vr(e),
                            r)
                        )
                            for (e = 0; e < r.length; e++)
                                (a = (a = (n = r[e])._getVersion)(n._source)),
                                    null == t.mutableSourceEagerHydrationData
                                        ? (t.mutableSourceEagerHydrationData = [
                                              n,
                                              a,
                                          ])
                                        : t.mutableSourceEagerHydrationData.push(
                                              n,
                                              a,
                                          );
                        return new Yu(t);
                    }),
                    (t.render = function (e, t, n) {
                        if (!Zu(t)) throw Error(l(200));
                        return ec(null, e, t, !1, n);
                    }),
                    (t.unmountComponentAtNode = function (e) {
                        if (!Zu(e)) throw Error(l(40));
                        return (
                            !!e._reactRootContainer &&
                            (du(function () {
                                ec(null, null, e, !1, function () {
                                    (e._reactRootContainer = null),
                                        (e[ma] = null);
                                });
                            }),
                            !0)
                        );
                    }),
                    (t.unstable_batchedUpdates = cu),
                    (t.unstable_renderSubtreeIntoContainer = function (
                        e,
                        t,
                        n,
                        r,
                    ) {
                        if (!Zu(n)) throw Error(l(200));
                        if (null == e || void 0 === e._reactInternals)
                            throw Error(l(38));
                        return ec(e, t, n, !1, r);
                    }),
                    (t.version = "18.2.0-next-9e3b772b8-20220608");
            },
            250: function (e, t, n) {
                var r = n(164);
                (t.createRoot = r.createRoot), (t.hydrateRoot = r.hydrateRoot);
            },
            164: function (e, t, n) {
                !(function e() {
                    if (
                        "undefined" !== typeof __REACT_DEVTOOLS_GLOBAL_HOOK__ &&
                        "function" ===
                            typeof __REACT_DEVTOOLS_GLOBAL_HOOK__.checkDCE
                    )
                        try {
                            __REACT_DEVTOOLS_GLOBAL_HOOK__.checkDCE(e);
                        } catch (t) {
                            console.error(t);
                        }
                })(),
                    (e.exports = n(463));
            },
            374: function (e, t, n) {
                var r = n(791),
                    a = Symbol.for("react.element"),
                    l = Symbol.for("react.fragment"),
                    i = Object.prototype.hasOwnProperty,
                    o =
                        r.__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED
                            .ReactCurrentOwner,
                    s = { key: !0, ref: !0, __self: !0, __source: !0 };
                function u(e, t, n) {
                    var r,
                        l = {},
                        u = null,
                        c = null;
                    for (r in (void 0 !== n && (u = "" + n),
                    void 0 !== t.key && (u = "" + t.key),
                    void 0 !== t.ref && (c = t.ref),
                    t))
                        i.call(t, r) && !s.hasOwnProperty(r) && (l[r] = t[r]);
                    if (e && e.defaultProps)
                        for (r in (t = e.defaultProps))
                            void 0 === l[r] && (l[r] = t[r]);
                    return {
                        $$typeof: a,
                        type: e,
                        key: u,
                        ref: c,
                        props: l,
                        _owner: o.current,
                    };
                }
                (t.Fragment = l), (t.jsx = u), (t.jsxs = u);
            },
            117: function (e, t) {
                var n = Symbol.for("react.element"),
                    r = Symbol.for("react.portal"),
                    a = Symbol.for("react.fragment"),
                    l = Symbol.for("react.strict_mode"),
                    i = Symbol.for("react.profiler"),
                    o = Symbol.for("react.provider"),
                    s = Symbol.for("react.context"),
                    u = Symbol.for("react.forward_ref"),
                    c = Symbol.for("react.suspense"),
                    d = Symbol.for("react.memo"),
                    f = Symbol.for("react.lazy"),
                    p = Symbol.iterator;
                var m = {
                        isMounted: function () {
                            return !1;
                        },
                        enqueueForceUpdate: function () {},
                        enqueueReplaceState: function () {},
                        enqueueSetState: function () {},
                    },
                    h = Object.assign,
                    v = {};
                function g(e, t, n) {
                    (this.props = e),
                        (this.context = t),
                        (this.refs = v),
                        (this.updater = n || m);
                }
                function y() {}
                function b(e, t, n) {
                    (this.props = e),
                        (this.context = t),
                        (this.refs = v),
                        (this.updater = n || m);
                }
                (g.prototype.isReactComponent = {}),
                    (g.prototype.setState = function (e, t) {
                        if (
                            "object" !== typeof e &&
                            "function" !== typeof e &&
                            null != e
                        )
                            throw Error(
                                "setState(...): takes an object of state variables to update or a function which returns an object of state variables.",
                            );
                        this.updater.enqueueSetState(this, e, t, "setState");
                    }),
                    (g.prototype.forceUpdate = function (e) {
                        this.updater.enqueueForceUpdate(this, e, "forceUpdate");
                    }),
                    (y.prototype = g.prototype);
                var x = (b.prototype = new y());
                (x.constructor = b),
                    h(x, g.prototype),
                    (x.isPureReactComponent = !0);
                var w = Array.isArray,
                    j = Object.prototype.hasOwnProperty,
                    k = { current: null },
                    S = { key: !0, ref: !0, __self: !0, __source: !0 };
                function N(e, t, r) {
                    var a,
                        l = {},
                        i = null,
                        o = null;
                    if (null != t)
                        for (a in (void 0 !== t.ref && (o = t.ref),
                        void 0 !== t.key && (i = "" + t.key),
                        t))
                            j.call(t, a) &&
                                !S.hasOwnProperty(a) &&
                                (l[a] = t[a]);
                    var s = arguments.length - 2;
                    if (1 === s) l.children = r;
                    else if (1 < s) {
                        for (var u = Array(s), c = 0; c < s; c++)
                            u[c] = arguments[c + 2];
                        l.children = u;
                    }
                    if (e && e.defaultProps)
                        for (a in (s = e.defaultProps))
                            void 0 === l[a] && (l[a] = s[a]);
                    return {
                        $$typeof: n,
                        type: e,
                        key: i,
                        ref: o,
                        props: l,
                        _owner: k.current,
                    };
                }
                function C(e) {
                    return (
                        "object" === typeof e && null !== e && e.$$typeof === n
                    );
                }
                var E = /\/+/g;
                function L(e, t) {
                    return "object" === typeof e && null !== e && null != e.key
                        ? (function (e) {
                              var t = { "=": "=0", ":": "=2" };
                              return (
                                  "$" +
                                  e.replace(/[=:]/g, function (e) {
                                      return t[e];
                                  })
                              );
                          })("" + e.key)
                        : t.toString(36);
                }
                function _(e, t, a, l, i) {
                    var o = typeof e;
                    ("undefined" !== o && "boolean" !== o) || (e = null);
                    var s = !1;
                    if (null === e) s = !0;
                    else
                        switch (o) {
                            case "string":
                            case "number":
                                s = !0;
                                break;
                            case "object":
                                switch (e.$$typeof) {
                                    case n:
                                    case r:
                                        s = !0;
                                }
                        }
                    if (s)
                        return (
                            (i = i((s = e))),
                            (e = "" === l ? "." + L(s, 0) : l),
                            w(i)
                                ? ((a = ""),
                                  null != e && (a = e.replace(E, "$&/") + "/"),
                                  _(i, t, a, "", function (e) {
                                      return e;
                                  }))
                                : null != i &&
                                  (C(i) &&
                                      (i = (function (e, t) {
                                          return {
                                              $$typeof: n,
                                              type: e.type,
                                              key: t,
                                              ref: e.ref,
                                              props: e.props,
                                              _owner: e._owner,
                                          };
                                      })(
                                          i,
                                          a +
                                              (!i.key || (s && s.key === i.key)
                                                  ? ""
                                                  : ("" + i.key).replace(
                                                        E,
                                                        "$&/",
                                                    ) + "/") +
                                              e,
                                      )),
                                  t.push(i)),
                            1
                        );
                    if (((s = 0), (l = "" === l ? "." : l + ":"), w(e)))
                        for (var u = 0; u < e.length; u++) {
                            var c = l + L((o = e[u]), u);
                            s += _(o, t, a, c, i);
                        }
                    else if (
                        ((c = (function (e) {
                            return null === e || "object" !== typeof e
                                ? null
                                : "function" ===
                                  typeof (e = (p && e[p]) || e["@@iterator"])
                                ? e
                                : null;
                        })(e)),
                        "function" === typeof c)
                    )
                        for (e = c.call(e), u = 0; !(o = e.next()).done; )
                            s += _((o = o.value), t, a, (c = l + L(o, u++)), i);
                    else if ("object" === o)
                        throw (
                            ((t = String(e)),
                            Error(
                                "Objects are not valid as a React child (found: " +
                                    ("[object Object]" === t
                                        ? "object with keys {" +
                                          Object.keys(e).join(", ") +
                                          "}"
                                        : t) +
                                    "). If you meant to render a collection of children, use an array instead.",
                            ))
                        );
                    return s;
                }
                function P(e, t, n) {
                    if (null == e) return e;
                    var r = [],
                        a = 0;
                    return (
                        _(e, r, "", "", function (e) {
                            return t.call(n, e, a++);
                        }),
                        r
                    );
                }
                function O(e) {
                    if (-1 === e._status) {
                        var t = e._result;
                        (t = t()).then(
                            function (t) {
                                (0 !== e._status && -1 !== e._status) ||
                                    ((e._status = 1), (e._result = t));
                            },
                            function (t) {
                                (0 !== e._status && -1 !== e._status) ||
                                    ((e._status = 2), (e._result = t));
                            },
                        ),
                            -1 === e._status &&
                                ((e._status = 0), (e._result = t));
                    }
                    if (1 === e._status) return e._result.default;
                    throw e._result;
                }
                var z = { current: null },
                    M = { transition: null },
                    R = {
                        ReactCurrentDispatcher: z,
                        ReactCurrentBatchConfig: M,
                        ReactCurrentOwner: k,
                    };
                (t.Children = {
                    map: P,
                    forEach: function (e, t, n) {
                        P(
                            e,
                            function () {
                                t.apply(this, arguments);
                            },
                            n,
                        );
                    },
                    count: function (e) {
                        var t = 0;
                        return (
                            P(e, function () {
                                t++;
                            }),
                            t
                        );
                    },
                    toArray: function (e) {
                        return (
                            P(e, function (e) {
                                return e;
                            }) || []
                        );
                    },
                    only: function (e) {
                        if (!C(e))
                            throw Error(
                                "React.Children.only expected to receive a single React element child.",
                            );
                        return e;
                    },
                }),
                    (t.Component = g),
                    (t.Fragment = a),
                    (t.Profiler = i),
                    (t.PureComponent = b),
                    (t.StrictMode = l),
                    (t.Suspense = c),
                    (t.__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED = R),
                    (t.cloneElement = function (e, t, r) {
                        if (null === e || void 0 === e)
                            throw Error(
                                "React.cloneElement(...): The argument must be a React element, but you passed " +
                                    e +
                                    ".",
                            );
                        var a = h({}, e.props),
                            l = e.key,
                            i = e.ref,
                            o = e._owner;
                        if (null != t) {
                            if (
                                (void 0 !== t.ref &&
                                    ((i = t.ref), (o = k.current)),
                                void 0 !== t.key && (l = "" + t.key),
                                e.type && e.type.defaultProps)
                            )
                                var s = e.type.defaultProps;
                            for (u in t)
                                j.call(t, u) &&
                                    !S.hasOwnProperty(u) &&
                                    (a[u] =
                                        void 0 === t[u] && void 0 !== s
                                            ? s[u]
                                            : t[u]);
                        }
                        var u = arguments.length - 2;
                        if (1 === u) a.children = r;
                        else if (1 < u) {
                            s = Array(u);
                            for (var c = 0; c < u; c++) s[c] = arguments[c + 2];
                            a.children = s;
                        }
                        return {
                            $$typeof: n,
                            type: e.type,
                            key: l,
                            ref: i,
                            props: a,
                            _owner: o,
                        };
                    }),
                    (t.createContext = function (e) {
                        return (
                            ((e = {
                                $$typeof: s,
                                _currentValue: e,
                                _currentValue2: e,
                                _threadCount: 0,
                                Provider: null,
                                Consumer: null,
                                _defaultValue: null,
                                _globalName: null,
                            }).Provider = { $$typeof: o, _context: e }),
                            (e.Consumer = e)
                        );
                    }),
                    (t.createElement = N),
                    (t.createFactory = function (e) {
                        var t = N.bind(null, e);
                        return (t.type = e), t;
                    }),
                    (t.createRef = function () {
                        return { current: null };
                    }),
                    (t.forwardRef = function (e) {
                        return { $$typeof: u, render: e };
                    }),
                    (t.isValidElement = C),
                    (t.lazy = function (e) {
                        return {
                            $$typeof: f,
                            _payload: { _status: -1, _result: e },
                            _init: O,
                        };
                    }),
                    (t.memo = function (e, t) {
                        return {
                            $$typeof: d,
                            type: e,
                            compare: void 0 === t ? null : t,
                        };
                    }),
                    (t.startTransition = function (e) {
                        var t = M.transition;
                        M.transition = {};
                        try {
                            e();
                        } finally {
                            M.transition = t;
                        }
                    }),
                    (t.unstable_act = function () {
                        throw Error(
                            "act(...) is not supported in production builds of React.",
                        );
                    }),
                    (t.useCallback = function (e, t) {
                        return z.current.useCallback(e, t);
                    }),
                    (t.useContext = function (e) {
                        return z.current.useContext(e);
                    }),
                    (t.useDebugValue = function () {}),
                    (t.useDeferredValue = function (e) {
                        return z.current.useDeferredValue(e);
                    }),
                    (t.useEffect = function (e, t) {
                        return z.current.useEffect(e, t);
                    }),
                    (t.useId = function () {
                        return z.current.useId();
                    }),
                    (t.useImperativeHandle = function (e, t, n) {
                        return z.current.useImperativeHandle(e, t, n);
                    }),
                    (t.useInsertionEffect = function (e, t) {
                        return z.current.useInsertionEffect(e, t);
                    }),
                    (t.useLayoutEffect = function (e, t) {
                        return z.current.useLayoutEffect(e, t);
                    }),
                    (t.useMemo = function (e, t) {
                        return z.current.useMemo(e, t);
                    }),
                    (t.useReducer = function (e, t, n) {
                        return z.current.useReducer(e, t, n);
                    }),
                    (t.useRef = function (e) {
                        return z.current.useRef(e);
                    }),
                    (t.useState = function (e) {
                        return z.current.useState(e);
                    }),
                    (t.useSyncExternalStore = function (e, t, n) {
                        return z.current.useSyncExternalStore(e, t, n);
                    }),
                    (t.useTransition = function () {
                        return z.current.useTransition();
                    }),
                    (t.version = "18.2.0");
            },
            791: function (e, t, n) {
                e.exports = n(117);
            },
            184: function (e, t, n) {
                e.exports = n(374);
            },
            813: function (e, t) {
                function n(e, t) {
                    var n = e.length;
                    e.push(t);
                    e: for (; 0 < n; ) {
                        var r = (n - 1) >>> 1,
                            a = e[r];
                        if (!(0 < l(a, t))) break e;
                        (e[r] = t), (e[n] = a), (n = r);
                    }
                }
                function r(e) {
                    return 0 === e.length ? null : e[0];
                }
                function a(e) {
                    if (0 === e.length) return null;
                    var t = e[0],
                        n = e.pop();
                    if (n !== t) {
                        e[0] = n;
                        e: for (var r = 0, a = e.length, i = a >>> 1; r < i; ) {
                            var o = 2 * (r + 1) - 1,
                                s = e[o],
                                u = o + 1,
                                c = e[u];
                            if (0 > l(s, n))
                                u < a && 0 > l(c, s)
                                    ? ((e[r] = c), (e[u] = n), (r = u))
                                    : ((e[r] = s), (e[o] = n), (r = o));
                            else {
                                if (!(u < a && 0 > l(c, n))) break e;
                                (e[r] = c), (e[u] = n), (r = u);
                            }
                        }
                    }
                    return t;
                }
                function l(e, t) {
                    var n = e.sortIndex - t.sortIndex;
                    return 0 !== n ? n : e.id - t.id;
                }
                if (
                    "object" === typeof performance &&
                    "function" === typeof performance.now
                ) {
                    var i = performance;
                    t.unstable_now = function () {
                        return i.now();
                    };
                } else {
                    var o = Date,
                        s = o.now();
                    t.unstable_now = function () {
                        return o.now() - s;
                    };
                }
                var u = [],
                    c = [],
                    d = 1,
                    f = null,
                    p = 3,
                    m = !1,
                    h = !1,
                    v = !1,
                    g = "function" === typeof setTimeout ? setTimeout : null,
                    y =
                        "function" === typeof clearTimeout
                            ? clearTimeout
                            : null,
                    b =
                        "undefined" !== typeof setImmediate
                            ? setImmediate
                            : null;
                function x(e) {
                    for (var t = r(c); null !== t; ) {
                        if (null === t.callback) a(c);
                        else {
                            if (!(t.startTime <= e)) break;
                            a(c), (t.sortIndex = t.expirationTime), n(u, t);
                        }
                        t = r(c);
                    }
                }
                function w(e) {
                    if (((v = !1), x(e), !h))
                        if (null !== r(u)) (h = !0), M(j);
                        else {
                            var t = r(c);
                            null !== t && R(w, t.startTime - e);
                        }
                }
                function j(e, n) {
                    (h = !1), v && ((v = !1), y(C), (C = -1)), (m = !0);
                    var l = p;
                    try {
                        for (
                            x(n), f = r(u);
                            null !== f &&
                            (!(f.expirationTime > n) || (e && !_()));

                        ) {
                            var i = f.callback;
                            if ("function" === typeof i) {
                                (f.callback = null), (p = f.priorityLevel);
                                var o = i(f.expirationTime <= n);
                                (n = t.unstable_now()),
                                    "function" === typeof o
                                        ? (f.callback = o)
                                        : f === r(u) && a(u),
                                    x(n);
                            } else a(u);
                            f = r(u);
                        }
                        if (null !== f) var s = !0;
                        else {
                            var d = r(c);
                            null !== d && R(w, d.startTime - n), (s = !1);
                        }
                        return s;
                    } finally {
                        (f = null), (p = l), (m = !1);
                    }
                }
                "undefined" !== typeof navigator &&
                    void 0 !== navigator.scheduling &&
                    void 0 !== navigator.scheduling.isInputPending &&
                    navigator.scheduling.isInputPending.bind(
                        navigator.scheduling,
                    );
                var k,
                    S = !1,
                    N = null,
                    C = -1,
                    E = 5,
                    L = -1;
                function _() {
                    return !(t.unstable_now() - L < E);
                }
                function P() {
                    if (null !== N) {
                        var e = t.unstable_now();
                        L = e;
                        var n = !0;
                        try {
                            n = N(!0, e);
                        } finally {
                            n ? k() : ((S = !1), (N = null));
                        }
                    } else S = !1;
                }
                if ("function" === typeof b)
                    k = function () {
                        b(P);
                    };
                else if ("undefined" !== typeof MessageChannel) {
                    var O = new MessageChannel(),
                        z = O.port2;
                    (O.port1.onmessage = P),
                        (k = function () {
                            z.postMessage(null);
                        });
                } else
                    k = function () {
                        g(P, 0);
                    };
                function M(e) {
                    (N = e), S || ((S = !0), k());
                }
                function R(e, n) {
                    C = g(function () {
                        e(t.unstable_now());
                    }, n);
                }
                (t.unstable_IdlePriority = 5),
                    (t.unstable_ImmediatePriority = 1),
                    (t.unstable_LowPriority = 4),
                    (t.unstable_NormalPriority = 3),
                    (t.unstable_Profiling = null),
                    (t.unstable_UserBlockingPriority = 2),
                    (t.unstable_cancelCallback = function (e) {
                        e.callback = null;
                    }),
                    (t.unstable_continueExecution = function () {
                        h || m || ((h = !0), M(j));
                    }),
                    (t.unstable_forceFrameRate = function (e) {
                        0 > e || 125 < e
                            ? console.error(
                                  "forceFrameRate takes a positive int between 0 and 125, forcing frame rates higher than 125 fps is not supported",
                              )
                            : (E = 0 < e ? Math.floor(1e3 / e) : 5);
                    }),
                    (t.unstable_getCurrentPriorityLevel = function () {
                        return p;
                    }),
                    (t.unstable_getFirstCallbackNode = function () {
                        return r(u);
                    }),
                    (t.unstable_next = function (e) {
                        switch (p) {
                            case 1:
                            case 2:
                            case 3:
                                var t = 3;
                                break;
                            default:
                                t = p;
                        }
                        var n = p;
                        p = t;
                        try {
                            return e();
                        } finally {
                            p = n;
                        }
                    }),
                    (t.unstable_pauseExecution = function () {}),
                    (t.unstable_requestPaint = function () {}),
                    (t.unstable_runWithPriority = function (e, t) {
                        switch (e) {
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                            case 5:
                                break;
                            default:
                                e = 3;
                        }
                        var n = p;
                        p = e;
                        try {
                            return t();
                        } finally {
                            p = n;
                        }
                    }),
                    (t.unstable_scheduleCallback = function (e, a, l) {
                        var i = t.unstable_now();
                        switch (
                            ("object" === typeof l && null !== l
                                ? (l =
                                      "number" === typeof (l = l.delay) && 0 < l
                                          ? i + l
                                          : i)
                                : (l = i),
                            e)
                        ) {
                            case 1:
                                var o = -1;
                                break;
                            case 2:
                                o = 250;
                                break;
                            case 5:
                                o = 1073741823;
                                break;
                            case 4:
                                o = 1e4;
                                break;
                            default:
                                o = 5e3;
                        }
                        return (
                            (e = {
                                id: d++,
                                callback: a,
                                priorityLevel: e,
                                startTime: l,
                                expirationTime: (o = l + o),
                                sortIndex: -1,
                            }),
                            l > i
                                ? ((e.sortIndex = l),
                                  n(c, e),
                                  null === r(u) &&
                                      e === r(c) &&
                                      (v ? (y(C), (C = -1)) : (v = !0),
                                      R(w, l - i)))
                                : ((e.sortIndex = o),
                                  n(u, e),
                                  h || m || ((h = !0), M(j))),
                            e
                        );
                    }),
                    (t.unstable_shouldYield = _),
                    (t.unstable_wrapCallback = function (e) {
                        var t = p;
                        return function () {
                            var n = p;
                            p = t;
                            try {
                                return e.apply(this, arguments);
                            } finally {
                                p = n;
                            }
                        };
                    });
            },
            296: function (e, t, n) {
                e.exports = n(813);
            },
        },
        t = {};
    function n(r) {
        var a = t[r];
        if (void 0 !== a) return a.exports;
        var l = (t[r] = { exports: {} });
        return e[r](l, l.exports, n), l.exports;
    }
    !(function () {
        var e,
            t = Object.getPrototypeOf
                ? function (e) {
                      return Object.getPrototypeOf(e);
                  }
                : function (e) {
                      return e.__proto__;
                  };
        n.t = function (r, a) {
            if ((1 & a && (r = this(r)), 8 & a)) return r;
            if ("object" === typeof r && r) {
                if (4 & a && r.__esModule) return r;
                if (16 & a && "function" === typeof r.then) return r;
            }
            var l = Object.create(null);
            n.r(l);
            var i = {};
            e = e || [null, t({}), t([]), t(t)];
            for (
                var o = 2 & a && r;
                "object" == typeof o && !~e.indexOf(o);
                o = t(o)
            )
                Object.getOwnPropertyNames(o).forEach(function (e) {
                    i[e] = function () {
                        return r[e];
                    };
                });
            return (
                (i.default = function () {
                    return r;
                }),
                n.d(l, i),
                l
            );
        };
    })(),
        (n.d = function (e, t) {
            for (var r in t)
                n.o(t, r) &&
                    !n.o(e, r) &&
                    Object.defineProperty(e, r, { enumerable: !0, get: t[r] });
        }),
        (n.o = function (e, t) {
            return Object.prototype.hasOwnProperty.call(e, t);
        }),
        (n.r = function (e) {
            "undefined" !== typeof Symbol &&
                Symbol.toStringTag &&
                Object.defineProperty(e, Symbol.toStringTag, {
                    value: "Module",
                }),
                Object.defineProperty(e, "__esModule", { value: !0 });
        }),
        (function () {
            var e = {};
            n.r(e),
                n.d(e, {
                    exclude: function () {
                        return Tn;
                    },
                    extract: function () {
                        return _n;
                    },
                    parse: function () {
                        return Pn;
                    },
                    parseUrl: function () {
                        return zn;
                    },
                    pick: function () {
                        return Rn;
                    },
                    stringify: function () {
                        return On;
                    },
                    stringifyUrl: function () {
                        return Mn;
                    },
                });
            var t,
                r = n(791),
                a = n.t(r, 2),
                l = n(250);
            function i(e) {
                if (Array.isArray(e)) return e;
            }
            function o(e, t) {
                (null == t || t > e.length) && (t = e.length);
                for (var n = 0, r = new Array(t); n < t; n++) r[n] = e[n];
                return r;
            }
            function s(e, t) {
                if (e) {
                    if ("string" === typeof e) return o(e, t);
                    var n = Object.prototype.toString.call(e).slice(8, -1);
                    return (
                        "Object" === n &&
                            e.constructor &&
                            (n = e.constructor.name),
                        "Map" === n || "Set" === n
                            ? Array.from(e)
                            : "Arguments" === n ||
                              /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)
                            ? o(e, t)
                            : void 0
                    );
                }
            }
            function u() {
                throw new TypeError(
                    "Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.",
                );
            }
            function c(e, t) {
                return (
                    i(e) ||
                    (function (e, t) {
                        var n =
                            null == e
                                ? null
                                : ("undefined" != typeof Symbol &&
                                      e[Symbol.iterator]) ||
                                  e["@@iterator"];
                        if (null != n) {
                            var r,
                                a,
                                l,
                                i,
                                o = [],
                                s = !0,
                                u = !1;
                            try {
                                if (((l = (n = n.call(e)).next), 0 === t)) {
                                    if (Object(n) !== n) return;
                                    s = !1;
                                } else
                                    for (
                                        ;
                                        !(s = (r = l.call(n)).done) &&
                                        (o.push(r.value), o.length !== t);
                                        s = !0
                                    );
                            } catch (e) {
                                (u = !0), (a = e);
                            } finally {
                                try {
                                    if (
                                        !s &&
                                        null != n.return &&
                                        ((i = n.return()), Object(i) !== i)
                                    )
                                        return;
                                } finally {
                                    if (u) throw a;
                                }
                            }
                            return o;
                        }
                    })(e, t) ||
                    s(e, t) ||
                    u()
                );
            }
            function d(e) {
                if (
                    ("undefined" !== typeof Symbol &&
                        null != e[Symbol.iterator]) ||
                    null != e["@@iterator"]
                )
                    return Array.from(e);
            }
            function f(e) {
                return (
                    (function (e) {
                        if (Array.isArray(e)) return o(e);
                    })(e) ||
                    d(e) ||
                    s(e) ||
                    (function () {
                        throw new TypeError(
                            "Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.",
                        );
                    })()
                );
            }
            function p(e, t) {
                if (!(e instanceof t))
                    throw new TypeError("Cannot call a class as a function");
            }
            function m(e) {
                return (
                    (m =
                        "function" == typeof Symbol &&
                        "symbol" == typeof Symbol.iterator
                            ? function (e) {
                                  return typeof e;
                              }
                            : function (e) {
                                  return e &&
                                      "function" == typeof Symbol &&
                                      e.constructor === Symbol &&
                                      e !== Symbol.prototype
                                      ? "symbol"
                                      : typeof e;
                              }),
                    m(e)
                );
            }
            function h(e) {
                var t = (function (e, t) {
                    if ("object" !== m(e) || null === e) return e;
                    var n = e[Symbol.toPrimitive];
                    if (void 0 !== n) {
                        var r = n.call(e, t || "default");
                        if ("object" !== m(r)) return r;
                        throw new TypeError(
                            "@@toPrimitive must return a primitive value.",
                        );
                    }
                    return ("string" === t ? String : Number)(e);
                })(e, "string");
                return "symbol" === m(t) ? t : String(t);
            }
            function v(e, t) {
                for (var n = 0; n < t.length; n++) {
                    var r = t[n];
                    (r.enumerable = r.enumerable || !1),
                        (r.configurable = !0),
                        "value" in r && (r.writable = !0),
                        Object.defineProperty(e, h(r.key), r);
                }
            }
            function g(e, t, n) {
                return (
                    t && v(e.prototype, t),
                    n && v(e, n),
                    Object.defineProperty(e, "prototype", { writable: !1 }),
                    e
                );
            }
            function y(e, t) {
                return (
                    (y = Object.setPrototypeOf
                        ? Object.setPrototypeOf.bind()
                        : function (e, t) {
                              return (e.__proto__ = t), e;
                          }),
                    y(e, t)
                );
            }
            function b(e, t) {
                if ("function" !== typeof t && null !== t)
                    throw new TypeError(
                        "Super expression must either be null or a function",
                    );
                (e.prototype = Object.create(t && t.prototype, {
                    constructor: { value: e, writable: !0, configurable: !0 },
                })),
                    Object.defineProperty(e, "prototype", { writable: !1 }),
                    t && y(e, t);
            }
            function x(e) {
                return (
                    (x = Object.setPrototypeOf
                        ? Object.getPrototypeOf.bind()
                        : function (e) {
                              return e.__proto__ || Object.getPrototypeOf(e);
                          }),
                    x(e)
                );
            }
            function w() {
                if ("undefined" === typeof Reflect || !Reflect.construct)
                    return !1;
                if (Reflect.construct.sham) return !1;
                if ("function" === typeof Proxy) return !0;
                try {
                    return (
                        Boolean.prototype.valueOf.call(
                            Reflect.construct(Boolean, [], function () {}),
                        ),
                        !0
                    );
                } catch (e) {
                    return !1;
                }
            }
            function j(e, t) {
                if (t && ("object" === m(t) || "function" === typeof t))
                    return t;
                if (void 0 !== t)
                    throw new TypeError(
                        "Derived constructors may only return object or undefined",
                    );
                return (function (e) {
                    if (void 0 === e)
                        throw new ReferenceError(
                            "this hasn't been initialised - super() hasn't been called",
                        );
                    return e;
                })(e);
            }
            function k(e) {
                var t = w();
                return function () {
                    var n,
                        r = x(e);
                    if (t) {
                        var a = x(this).constructor;
                        n = Reflect.construct(r, arguments, a);
                    } else n = r.apply(this, arguments);
                    return j(this, n);
                };
            }
            function S(e, t, n) {
                return (
                    (S = w()
                        ? Reflect.construct.bind()
                        : function (e, t, n) {
                              var r = [null];
                              r.push.apply(r, t);
                              var a = new (Function.bind.apply(e, r))();
                              return n && y(a, n.prototype), a;
                          }),
                    S.apply(null, arguments)
                );
            }
            function N(e) {
                var t = "function" === typeof Map ? new Map() : void 0;
                return (
                    (N = function (e) {
                        if (
                            null === e ||
                            ((n = e),
                            -1 ===
                                Function.toString
                                    .call(n)
                                    .indexOf("[native code]"))
                        )
                            return e;
                        var n;
                        if ("function" !== typeof e)
                            throw new TypeError(
                                "Super expression must either be null or a function",
                            );
                        if ("undefined" !== typeof t) {
                            if (t.has(e)) return t.get(e);
                            t.set(e, r);
                        }
                        function r() {
                            return S(e, arguments, x(this).constructor);
                        }
                        return (
                            (r.prototype = Object.create(e.prototype, {
                                constructor: {
                                    value: r,
                                    enumerable: !1,
                                    writable: !0,
                                    configurable: !0,
                                },
                            })),
                            y(r, e)
                        );
                    }),
                    N(e)
                );
            }
            function C(e, t) {
                var n =
                    ("undefined" !== typeof Symbol && e[Symbol.iterator]) ||
                    e["@@iterator"];
                if (!n) {
                    if (
                        Array.isArray(e) ||
                        (n = s(e)) ||
                        (t && e && "number" === typeof e.length)
                    ) {
                        n && (e = n);
                        var r = 0,
                            a = function () {};
                        return {
                            s: a,
                            n: function () {
                                return r >= e.length
                                    ? { done: !0 }
                                    : { done: !1, value: e[r++] };
                            },
                            e: function (e) {
                                throw e;
                            },
                            f: a,
                        };
                    }
                    throw new TypeError(
                        "Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.",
                    );
                }
                var l,
                    i = !0,
                    o = !1;
                return {
                    s: function () {
                        n = n.call(e);
                    },
                    n: function () {
                        var e = n.next();
                        return (i = e.done), e;
                    },
                    e: function (e) {
                        (o = !0), (l = e);
                    },
                    f: function () {
                        try {
                            i || null == n.return || n.return();
                        } finally {
                            if (o) throw l;
                        }
                    },
                };
            }
            function E() {
                return (
                    (E = Object.assign
                        ? Object.assign.bind()
                        : function (e) {
                              for (var t = 1; t < arguments.length; t++) {
                                  var n = arguments[t];
                                  for (var r in n)
                                      Object.prototype.hasOwnProperty.call(
                                          n,
                                          r,
                                      ) && (e[r] = n[r]);
                              }
                              return e;
                          }),
                    E.apply(this, arguments)
                );
            }
            !(function (e) {
                (e.Pop = "POP"), (e.Push = "PUSH"), (e.Replace = "REPLACE");
            })(t || (t = {}));
            var L,
                _ = "popstate";
            function P(e, t) {
                if (!1 === e || null === e || "undefined" === typeof e)
                    throw new Error(t);
            }
            function O(e, t) {
                if (!e) {
                    "undefined" !== typeof console && console.warn(t);
                    try {
                        throw new Error(t);
                    } catch (n) {}
                }
            }
            function z(e, t) {
                return { usr: e.state, key: e.key, idx: t };
            }
            function M(e, t, n, r) {
                return (
                    void 0 === n && (n = null),
                    E(
                        {
                            pathname: "string" === typeof e ? e : e.pathname,
                            search: "",
                            hash: "",
                        },
                        "string" === typeof t ? T(t) : t,
                        {
                            state: n,
                            key:
                                (t && t.key) ||
                                r ||
                                Math.random().toString(36).substr(2, 8),
                        },
                    )
                );
            }
            function R(e) {
                var t = e.pathname,
                    n = void 0 === t ? "/" : t,
                    r = e.search,
                    a = void 0 === r ? "" : r,
                    l = e.hash,
                    i = void 0 === l ? "" : l;
                return (
                    a && "?" !== a && (n += "?" === a.charAt(0) ? a : "?" + a),
                    i && "#" !== i && (n += "#" === i.charAt(0) ? i : "#" + i),
                    n
                );
            }
            function T(e) {
                var t = {};
                if (e) {
                    var n = e.indexOf("#");
                    n >= 0 && ((t.hash = e.substr(n)), (e = e.substr(0, n)));
                    var r = e.indexOf("?");
                    r >= 0 && ((t.search = e.substr(r)), (e = e.substr(0, r))),
                        e && (t.pathname = e);
                }
                return t;
            }
            function F(e, n, r, a) {
                void 0 === a && (a = {});
                var l = a,
                    i = l.window,
                    o = void 0 === i ? document.defaultView : i,
                    s = l.v5Compat,
                    u = void 0 !== s && s,
                    c = o.history,
                    d = t.Pop,
                    f = null,
                    p = m();
                function m() {
                    return (c.state || { idx: null }).idx;
                }
                function h() {
                    d = t.Pop;
                    var e = m(),
                        n = null == e ? null : e - p;
                    (p = e),
                        f && f({ action: d, location: g.location, delta: n });
                }
                function v(e) {
                    var t =
                            "null" !== o.location.origin
                                ? o.location.origin
                                : o.location.href,
                        n = "string" === typeof e ? e : R(e);
                    return (
                        P(
                            t,
                            "No window.location.(origin|href) available to create URL for href: " +
                                n,
                        ),
                        new URL(n, t)
                    );
                }
                null == p &&
                    ((p = 0), c.replaceState(E({}, c.state, { idx: p }), ""));
                var g = {
                    get action() {
                        return d;
                    },
                    get location() {
                        return e(o, c);
                    },
                    listen: function (e) {
                        if (f)
                            throw new Error(
                                "A history only accepts one active listener",
                            );
                        return (
                            o.addEventListener(_, h),
                            (f = e),
                            function () {
                                o.removeEventListener(_, h), (f = null);
                            }
                        );
                    },
                    createHref: function (e) {
                        return n(o, e);
                    },
                    createURL: v,
                    encodeLocation: function (e) {
                        var t = v(e);
                        return {
                            pathname: t.pathname,
                            search: t.search,
                            hash: t.hash,
                        };
                    },
                    push: function (e, n) {
                        d = t.Push;
                        var a = M(g.location, e, n);
                        r && r(a, e);
                        var l = z(a, (p = m() + 1)),
                            i = g.createHref(a);
                        try {
                            c.pushState(l, "", i);
                        } catch (s) {
                            if (
                                s instanceof DOMException &&
                                "DataCloneError" === s.name
                            )
                                throw s;
                            o.location.assign(i);
                        }
                        u &&
                            f &&
                            f({ action: d, location: g.location, delta: 1 });
                    },
                    replace: function (e, n) {
                        d = t.Replace;
                        var a = M(g.location, e, n);
                        r && r(a, e);
                        var l = z(a, (p = m())),
                            i = g.createHref(a);
                        c.replaceState(l, "", i),
                            u &&
                                f &&
                                f({
                                    action: d,
                                    location: g.location,
                                    delta: 0,
                                });
                    },
                    go: function (e) {
                        return c.go(e);
                    },
                };
                return g;
            }
            !(function (e) {
                (e.data = "data"),
                    (e.deferred = "deferred"),
                    (e.redirect = "redirect"),
                    (e.error = "error");
            })(L || (L = {}));
            new Set([
                "lazy",
                "caseSensitive",
                "path",
                "id",
                "index",
                "children",
            ]);
            function I(e, t, n) {
                void 0 === n && (n = "/");
                var r = X(
                    ("string" === typeof t ? T(t) : t).pathname || "/",
                    n,
                );
                if (null == r) return null;
                var a = D(e);
                !(function (e) {
                    e.sort(function (e, t) {
                        return e.score !== t.score
                            ? t.score - e.score
                            : (function (e, t) {
                                  var n =
                                      e.length === t.length &&
                                      e.slice(0, -1).every(function (e, n) {
                                          return e === t[n];
                                      });
                                  return n
                                      ? e[e.length - 1] - t[t.length - 1]
                                      : 0;
                              })(
                                  e.routesMeta.map(function (e) {
                                      return e.childrenIndex;
                                  }),
                                  t.routesMeta.map(function (e) {
                                      return e.childrenIndex;
                                  }),
                              );
                    });
                })(a);
                for (var l = null, i = 0; null == l && i < a.length; ++i)
                    l = q(a[i], Y(r));
                return l;
            }
            function D(e, t, n, r) {
                void 0 === t && (t = []),
                    void 0 === n && (n = []),
                    void 0 === r && (r = "");
                var a = function (e, a, l) {
                    var i = {
                        relativePath: void 0 === l ? e.path || "" : l,
                        caseSensitive: !0 === e.caseSensitive,
                        childrenIndex: a,
                        route: e,
                    };
                    i.relativePath.startsWith("/") &&
                        (P(
                            i.relativePath.startsWith(r),
                            'Absolute route path "' +
                                i.relativePath +
                                '" nested under path "' +
                                r +
                                '" is not valid. An absolute child route path must start with the combined path of all its parent routes.',
                        ),
                        (i.relativePath = i.relativePath.slice(r.length)));
                    var o = te([r, i.relativePath]),
                        s = n.concat(i);
                    e.children &&
                        e.children.length > 0 &&
                        (P(
                            !0 !== e.index,
                            'Index routes must not have child routes. Please remove all child routes from route path "' +
                                o +
                                '".',
                        ),
                        D(e.children, t, s, o)),
                        (null != e.path || e.index) &&
                            t.push({
                                path: o,
                                score: Q(o, e.index),
                                routesMeta: s,
                            });
                };
                return (
                    e.forEach(function (e, t) {
                        var n;
                        if (
                            "" !== e.path &&
                            null != (n = e.path) &&
                            n.includes("?")
                        ) {
                            var r,
                                l = C(U(e.path));
                            try {
                                for (l.s(); !(r = l.n()).done; ) {
                                    var i = r.value;
                                    a(e, t, i);
                                }
                            } catch (o) {
                                l.e(o);
                            } finally {
                                l.f();
                            }
                        } else a(e, t);
                    }),
                    t
                );
            }
            function U(e) {
                var t = e.split("/");
                if (0 === t.length) return [];
                var n,
                    r = i((n = t)) || d(n) || s(n) || u(),
                    a = r[0],
                    l = r.slice(1),
                    o = a.endsWith("?"),
                    c = a.replace(/\?$/, "");
                if (0 === l.length) return o ? [c, ""] : [c];
                var p = U(l.join("/")),
                    m = [];
                return (
                    m.push.apply(
                        m,
                        f(
                            p.map(function (e) {
                                return "" === e ? c : [c, e].join("/");
                            }),
                        ),
                    ),
                    o && m.push.apply(m, f(p)),
                    m.map(function (t) {
                        return e.startsWith("/") && "" === t ? "/" : t;
                    })
                );
            }
            var B = /^:\w+$/,
                A = 3,
                V = 2,
                $ = 1,
                H = 10,
                K = -2,
                W = function (e) {
                    return "*" === e;
                };
            function Q(e, t) {
                var n = e.split("/"),
                    r = n.length;
                return (
                    n.some(W) && (r += K),
                    t && (r += V),
                    n
                        .filter(function (e) {
                            return !W(e);
                        })
                        .reduce(function (e, t) {
                            return e + (B.test(t) ? A : "" === t ? $ : H);
                        }, r)
                );
            }
            function q(e, t) {
                for (
                    var n = e.routesMeta, r = {}, a = "/", l = [], i = 0;
                    i < n.length;
                    ++i
                ) {
                    var o = n[i],
                        s = i === n.length - 1,
                        u = "/" === a ? t : t.slice(a.length) || "/",
                        c = G(
                            {
                                path: o.relativePath,
                                caseSensitive: o.caseSensitive,
                                end: s,
                            },
                            u,
                        );
                    if (!c) return null;
                    Object.assign(r, c.params);
                    var d = o.route;
                    l.push({
                        params: r,
                        pathname: te([a, c.pathname]),
                        pathnameBase: ne(te([a, c.pathnameBase])),
                        route: d,
                    }),
                        "/" !== c.pathnameBase && (a = te([a, c.pathnameBase]));
                }
                return l;
            }
            function G(e, t) {
                "string" === typeof e &&
                    (e = { path: e, caseSensitive: !1, end: !0 });
                var n = (function (e, t, n) {
                        void 0 === t && (t = !1);
                        void 0 === n && (n = !0);
                        O(
                            "*" === e || !e.endsWith("*") || e.endsWith("/*"),
                            'Route path "' +
                                e +
                                '" will be treated as if it were "' +
                                e.replace(/\*$/, "/*") +
                                '" because the `*` character must always follow a `/` in the pattern. To get rid of this warning, please change the route path to "' +
                                e.replace(/\*$/, "/*") +
                                '".',
                        );
                        var r = [],
                            a =
                                "^" +
                                e
                                    .replace(/\/*\*?$/, "")
                                    .replace(/^\/*/, "/")
                                    .replace(/[\\.*+^$?{}|()[\]]/g, "\\$&")
                                    .replace(/\/:(\w+)/g, function (e, t) {
                                        return r.push(t), "/([^\\/]+)";
                                    });
                        e.endsWith("*")
                            ? (r.push("*"),
                              (a +=
                                  "*" === e || "/*" === e
                                      ? "(.*)$"
                                      : "(?:\\/(.+)|\\/*)$"))
                            : n
                            ? (a += "\\/*$")
                            : "" !== e && "/" !== e && (a += "(?:(?=\\/|$))");
                        var l = new RegExp(a, t ? void 0 : "i");
                        return [l, r];
                    })(e.path, e.caseSensitive, e.end),
                    r = c(n, 2),
                    a = r[0],
                    l = r[1],
                    i = t.match(a);
                if (!i) return null;
                var o = i[0],
                    s = o.replace(/(.)\/+$/, "$1"),
                    u = i.slice(1),
                    d = l.reduce(function (e, t, n) {
                        if ("*" === t) {
                            var r = u[n] || "";
                            s = o
                                .slice(0, o.length - r.length)
                                .replace(/(.)\/+$/, "$1");
                        }
                        return (
                            (e[t] = (function (e, t) {
                                try {
                                    return decodeURIComponent(e);
                                } catch (n) {
                                    return (
                                        O(
                                            !1,
                                            'The value for the URL param "' +
                                                t +
                                                '" will not be decoded because the string "' +
                                                e +
                                                '" is a malformed URL segment. This is probably due to a bad percent encoding (' +
                                                n +
                                                ").",
                                        ),
                                        e
                                    );
                                }
                            })(u[n] || "", t)),
                            e
                        );
                    }, {});
                return { params: d, pathname: o, pathnameBase: s, pattern: e };
            }
            function Y(e) {
                try {
                    return decodeURI(e);
                } catch (t) {
                    return (
                        O(
                            !1,
                            'The URL path "' +
                                e +
                                '" could not be decoded because it is is a malformed URL segment. This is probably due to a bad percent encoding (' +
                                t +
                                ").",
                        ),
                        e
                    );
                }
            }
            function X(e, t) {
                if ("/" === t) return e;
                if (!e.toLowerCase().startsWith(t.toLowerCase())) return null;
                var n = t.endsWith("/") ? t.length - 1 : t.length,
                    r = e.charAt(n);
                return r && "/" !== r ? null : e.slice(n) || "/";
            }
            function Z(e, t, n, r) {
                return (
                    "Cannot include a '" +
                    e +
                    "' character in a manually specified `to." +
                    t +
                    "` field [" +
                    JSON.stringify(r) +
                    "].  Please separate it out to the `to." +
                    n +
                    '` field. Alternatively you may provide the full path as a string in <Link to="..."> and the router will parse it for you.'
                );
            }
            function J(e) {
                return e.filter(function (e, t) {
                    return 0 === t || (e.route.path && e.route.path.length > 0);
                });
            }
            function ee(e, t, n, r) {
                var a;
                void 0 === r && (r = !1),
                    "string" === typeof e
                        ? (a = T(e))
                        : (P(
                              !(a = E({}, e)).pathname ||
                                  !a.pathname.includes("?"),
                              Z("?", "pathname", "search", a),
                          ),
                          P(
                              !a.pathname || !a.pathname.includes("#"),
                              Z("#", "pathname", "hash", a),
                          ),
                          P(
                              !a.search || !a.search.includes("#"),
                              Z("#", "search", "hash", a),
                          ));
                var l,
                    i = "" === e || "" === a.pathname,
                    o = i ? "/" : a.pathname;
                if (r || null == o) l = n;
                else {
                    var s = t.length - 1;
                    if (o.startsWith("..")) {
                        for (var u = o.split("/"); ".." === u[0]; )
                            u.shift(), (s -= 1);
                        a.pathname = u.join("/");
                    }
                    l = s >= 0 ? t[s] : "/";
                }
                var c = (function (e, t) {
                        void 0 === t && (t = "/");
                        var n = "string" === typeof e ? T(e) : e,
                            r = n.pathname,
                            a = n.search,
                            l = void 0 === a ? "" : a,
                            i = n.hash,
                            o = void 0 === i ? "" : i,
                            s = r
                                ? r.startsWith("/")
                                    ? r
                                    : (function (e, t) {
                                          var n = t
                                              .replace(/\/+$/, "")
                                              .split("/");
                                          return (
                                              e
                                                  .split("/")
                                                  .forEach(function (e) {
                                                      ".." === e
                                                          ? n.length > 1 &&
                                                            n.pop()
                                                          : "." !== e &&
                                                            n.push(e);
                                                  }),
                                              n.length > 1 ? n.join("/") : "/"
                                          );
                                      })(r, t)
                                : t;
                        return { pathname: s, search: re(l), hash: ae(o) };
                    })(a, l),
                    d = o && "/" !== o && o.endsWith("/"),
                    f = (i || "." === o) && n.endsWith("/");
                return (
                    c.pathname.endsWith("/") ||
                        (!d && !f) ||
                        (c.pathname += "/"),
                    c
                );
            }
            var te = function (e) {
                    return e.join("/").replace(/\/\/+/g, "/");
                },
                ne = function (e) {
                    return e.replace(/\/+$/, "").replace(/^\/*/, "/");
                },
                re = function (e) {
                    return e && "?" !== e
                        ? e.startsWith("?")
                            ? e
                            : "?" + e
                        : "";
                },
                ae = function (e) {
                    return e && "#" !== e
                        ? e.startsWith("#")
                            ? e
                            : "#" + e
                        : "";
                },
                le = (function (e) {
                    b(n, e);
                    var t = k(n);
                    function n() {
                        return p(this, n), t.apply(this, arguments);
                    }
                    return g(n);
                })(N(Error));
            function ie(e) {
                return (
                    null != e &&
                    "number" === typeof e.status &&
                    "string" === typeof e.statusText &&
                    "boolean" === typeof e.internal &&
                    "data" in e
                );
            }
            var oe = ["post", "put", "patch", "delete"],
                se = (new Set(oe), ["get"].concat(oe));
            new Set(se),
                new Set([301, 302, 303, 307, 308]),
                new Set([307, 308]);
            Symbol("deferred");
            function ue() {
                return (
                    (ue = Object.assign
                        ? Object.assign.bind()
                        : function (e) {
                              for (var t = 1; t < arguments.length; t++) {
                                  var n = arguments[t];
                                  for (var r in n)
                                      Object.prototype.hasOwnProperty.call(
                                          n,
                                          r,
                                      ) && (e[r] = n[r]);
                              }
                              return e;
                          }),
                    ue.apply(this, arguments)
                );
            }
            var ce = r.createContext(null);
            var de = r.createContext(null);
            var fe = r.createContext(null);
            var pe = r.createContext(null);
            var me = r.createContext(null);
            var he = r.createContext({
                outlet: null,
                matches: [],
                isDataRoute: !1,
            });
            var ve = r.createContext(null);
            function ge() {
                return null != r.useContext(me);
            }
            function ye() {
                return ge() || P(!1), r.useContext(me).location;
            }
            function be(e) {
                r.useContext(pe).static || r.useLayoutEffect(e);
            }
            function xe() {
                return r.useContext(he).isDataRoute
                    ? (function () {
                          var e = ze(Pe.UseNavigateStable).router,
                              t = Re(Oe.UseNavigateStable),
                              n = r.useRef(!1);
                          return (
                              be(function () {
                                  n.current = !0;
                              }),
                              r.useCallback(
                                  function (r, a) {
                                      void 0 === a && (a = {}),
                                          n.current &&
                                              ("number" === typeof r
                                                  ? e.navigate(r)
                                                  : e.navigate(
                                                        r,
                                                        ue(
                                                            { fromRouteId: t },
                                                            a,
                                                        ),
                                                    ));
                                  },
                                  [e, t],
                              )
                          );
                      })()
                    : (function () {
                          ge() || P(!1);
                          var e = r.useContext(ce),
                              t = r.useContext(pe),
                              n = t.basename,
                              a = t.navigator,
                              l = r.useContext(he).matches,
                              i = ye().pathname,
                              o = JSON.stringify(
                                  J(l).map(function (e) {
                                      return e.pathnameBase;
                                  }),
                              ),
                              s = r.useRef(!1);
                          return (
                              be(function () {
                                  s.current = !0;
                              }),
                              r.useCallback(
                                  function (t, r) {
                                      if ((void 0 === r && (r = {}), s.current))
                                          if ("number" !== typeof t) {
                                              var l = ee(
                                                  t,
                                                  JSON.parse(o),
                                                  i,
                                                  "path" === r.relative,
                                              );
                                              null == e &&
                                                  "/" !== n &&
                                                  (l.pathname =
                                                      "/" === l.pathname
                                                          ? n
                                                          : te([
                                                                n,
                                                                l.pathname,
                                                            ])),
                                                  (r.replace
                                                      ? a.replace
                                                      : a.push)(l, r.state, r);
                                          } else a.go(t);
                                  },
                                  [n, a, o, i, e],
                              )
                          );
                      })();
            }
            var we = r.createContext(null);
            function je() {
                var e = r.useContext(he).matches,
                    t = e[e.length - 1];
                return t ? t.params : {};
            }
            function ke(e, t) {
                var n = (void 0 === t ? {} : t).relative,
                    a = r.useContext(he).matches,
                    l = ye().pathname,
                    i = JSON.stringify(
                        J(a).map(function (e) {
                            return e.pathnameBase;
                        }),
                    );
                return r.useMemo(
                    function () {
                        return ee(e, JSON.parse(i), l, "path" === n);
                    },
                    [e, i, l, n],
                );
            }
            function Se(e, n, a) {
                ge() || P(!1);
                var l,
                    i = r.useContext(pe).navigator,
                    o = r.useContext(he).matches,
                    s = o[o.length - 1],
                    u = s ? s.params : {},
                    c = (s && s.pathname, s ? s.pathnameBase : "/"),
                    d = (s && s.route, ye());
                if (n) {
                    var f,
                        p = "string" === typeof n ? T(n) : n;
                    "/" === c ||
                        (null == (f = p.pathname) ? void 0 : f.startsWith(c)) ||
                        P(!1),
                        (l = p);
                } else l = d;
                var m = l.pathname || "/",
                    h = I(e, {
                        pathname: "/" === c ? m : m.slice(c.length) || "/",
                    });
                var v = _e(
                    h &&
                        h.map(function (e) {
                            return Object.assign({}, e, {
                                params: Object.assign({}, u, e.params),
                                pathname: te([
                                    c,
                                    i.encodeLocation
                                        ? i.encodeLocation(e.pathname).pathname
                                        : e.pathname,
                                ]),
                                pathnameBase:
                                    "/" === e.pathnameBase
                                        ? c
                                        : te([
                                              c,
                                              i.encodeLocation
                                                  ? i.encodeLocation(
                                                        e.pathnameBase,
                                                    ).pathname
                                                  : e.pathnameBase,
                                          ]),
                            });
                        }),
                    o,
                    a,
                );
                return n && v
                    ? r.createElement(
                          me.Provider,
                          {
                              value: {
                                  location: ue(
                                      {
                                          pathname: "/",
                                          search: "",
                                          hash: "",
                                          state: null,
                                          key: "default",
                                      },
                                      l,
                                  ),
                                  navigationType: t.Pop,
                              },
                          },
                          v,
                      )
                    : v;
            }
            function Ne() {
                var e = (function () {
                        var e,
                            t = r.useContext(ve),
                            n = Me(Oe.UseRouteError),
                            a = Re(Oe.UseRouteError);
                        if (t) return t;
                        return null == (e = n.errors) ? void 0 : e[a];
                    })(),
                    t = ie(e)
                        ? e.status + " " + e.statusText
                        : e instanceof Error
                        ? e.message
                        : JSON.stringify(e),
                    n = e instanceof Error ? e.stack : null,
                    a = "rgba(200,200,200, 0.5)",
                    l = { padding: "0.5rem", backgroundColor: a };
                return r.createElement(
                    r.Fragment,
                    null,
                    r.createElement(
                        "h2",
                        null,
                        "Unexpected Application error!",
                    ),
                    r.createElement(
                        "h3",
                        { style: { fontStyle: "italic" } },
                        t,
                    ),
                    n ? r.createElement("pre", { style: l }, n) : null,
                    null,
                );
            }
            var Ce = r.createElement(Ne, null),
                Ee = (function (e) {
                    b(n, e);
                    var t = k(n);
                    function n(e) {
                        var r;
                        return (
                            p(this, n),
                            ((r = t.call(this, e)).state = {
                                location: e.location,
                                revalidation: e.revalidation,
                                error: e.error,
                            }),
                            r
                        );
                    }
                    return (
                        g(
                            n,
                            [
                                {
                                    key: "componentDidCatch",
                                    value: function (e, t) {
                                        console.error(
                                            "React Router caught the following error during render",
                                            e,
                                            t,
                                        );
                                    },
                                },
                                {
                                    key: "render",
                                    value: function () {
                                        return this.state.error
                                            ? r.createElement(
                                                  he.Provider,
                                                  {
                                                      value: this.props
                                                          .routeContext,
                                                  },
                                                  r.createElement(ve.Provider, {
                                                      value: this.state.error,
                                                      children:
                                                          this.props.component,
                                                  }),
                                              )
                                            : this.props.children;
                                    },
                                },
                            ],
                            [
                                {
                                    key: "getDerivedStateFromError",
                                    value: function (e) {
                                        return { error: e };
                                    },
                                },
                                {
                                    key: "getDerivedStateFromProps",
                                    value: function (e, t) {
                                        return t.location !== e.location ||
                                            ("idle" !== t.revalidation &&
                                                "idle" === e.revalidation)
                                            ? {
                                                  error: e.error,
                                                  location: e.location,
                                                  revalidation: e.revalidation,
                                              }
                                            : {
                                                  error: e.error || t.error,
                                                  location: t.location,
                                                  revalidation:
                                                      e.revalidation ||
                                                      t.revalidation,
                                              };
                                    },
                                },
                            ],
                        ),
                        n
                    );
                })(r.Component);
            function Le(e) {
                var t = e.routeContext,
                    n = e.match,
                    a = e.children,
                    l = r.useContext(ce);
                return (
                    l &&
                        l.static &&
                        l.staticContext &&
                        (n.route.errorElement || n.route.ErrorBoundary) &&
                        (l.staticContext._deepestRenderedBoundaryId =
                            n.route.id),
                    r.createElement(he.Provider, { value: t }, a)
                );
            }
            function _e(e, t, n) {
                var a;
                if (
                    (void 0 === t && (t = []),
                    void 0 === n && (n = null),
                    null == e)
                ) {
                    var l;
                    if (null == (l = n) || !l.errors) return null;
                    e = n.matches;
                }
                var i = e,
                    o = null == (a = n) ? void 0 : a.errors;
                if (null != o) {
                    var s = i.findIndex(function (e) {
                        return (
                            e.route.id && (null == o ? void 0 : o[e.route.id])
                        );
                    });
                    s >= 0 || P(!1),
                        (i = i.slice(0, Math.min(i.length, s + 1)));
                }
                return i.reduceRight(function (e, a, l) {
                    var s = a.route.id
                            ? null == o
                                ? void 0
                                : o[a.route.id]
                            : null,
                        u = null;
                    n && (u = a.route.errorElement || Ce);
                    var c = t.concat(i.slice(0, l + 1)),
                        d = function () {
                            var t;
                            return (
                                (t = s
                                    ? u
                                    : a.route.Component
                                    ? r.createElement(a.route.Component, null)
                                    : a.route.element
                                    ? a.route.element
                                    : e),
                                r.createElement(Le, {
                                    match: a,
                                    routeContext: {
                                        outlet: e,
                                        matches: c,
                                        isDataRoute: null != n,
                                    },
                                    children: t,
                                })
                            );
                        };
                    return n &&
                        (a.route.ErrorBoundary ||
                            a.route.errorElement ||
                            0 === l)
                        ? r.createElement(Ee, {
                              location: n.location,
                              revalidation: n.revalidation,
                              component: u,
                              error: s,
                              children: d(),
                              routeContext: {
                                  outlet: null,
                                  matches: c,
                                  isDataRoute: !0,
                              },
                          })
                        : d();
                }, null);
            }
            var Pe = (function (e) {
                    return (
                        (e.UseBlocker = "useBlocker"),
                        (e.UseRevalidator = "useRevalidator"),
                        (e.UseNavigateStable = "useNavigate"),
                        e
                    );
                })(Pe || {}),
                Oe = (function (e) {
                    return (
                        (e.UseBlocker = "useBlocker"),
                        (e.UseLoaderData = "useLoaderData"),
                        (e.UseActionData = "useActionData"),
                        (e.UseRouteError = "useRouteError"),
                        (e.UseNavigation = "useNavigation"),
                        (e.UseRouteLoaderData = "useRouteLoaderData"),
                        (e.UseMatches = "useMatches"),
                        (e.UseRevalidator = "useRevalidator"),
                        (e.UseNavigateStable = "useNavigate"),
                        (e.UseRouteId = "useRouteId"),
                        e
                    );
                })(Oe || {});
            function ze(e) {
                var t = r.useContext(ce);
                return t || P(!1), t;
            }
            function Me(e) {
                var t = r.useContext(de);
                return t || P(!1), t;
            }
            function Re(e) {
                var t = (function (e) {
                        var t = r.useContext(he);
                        return t || P(!1), t;
                    })(),
                    n = t.matches[t.matches.length - 1];
                return n.route.id || P(!1), n.route.id;
            }
            a.startTransition;
            function Te(e) {
                return (function (e) {
                    var t = r.useContext(he).outlet;
                    return t
                        ? r.createElement(we.Provider, { value: e }, t)
                        : t;
                })(e.context);
            }
            function Fe(e) {
                P(!1);
            }
            function Ie(e) {
                var n = e.basename,
                    a = void 0 === n ? "/" : n,
                    l = e.children,
                    i = void 0 === l ? null : l,
                    o = e.location,
                    s = e.navigationType,
                    u = void 0 === s ? t.Pop : s,
                    c = e.navigator,
                    d = e.static,
                    f = void 0 !== d && d;
                ge() && P(!1);
                var p = a.replace(/^\/*/, "/"),
                    m = r.useMemo(
                        function () {
                            return { basename: p, navigator: c, static: f };
                        },
                        [p, c, f],
                    );
                "string" === typeof o && (o = T(o));
                var h = o,
                    v = h.pathname,
                    g = void 0 === v ? "/" : v,
                    y = h.search,
                    b = void 0 === y ? "" : y,
                    x = h.hash,
                    w = void 0 === x ? "" : x,
                    j = h.state,
                    k = void 0 === j ? null : j,
                    S = h.key,
                    N = void 0 === S ? "default" : S,
                    C = r.useMemo(
                        function () {
                            var e = X(g, p);
                            return null == e
                                ? null
                                : {
                                      location: {
                                          pathname: e,
                                          search: b,
                                          hash: w,
                                          state: k,
                                          key: N,
                                      },
                                      navigationType: u,
                                  };
                        },
                        [p, g, b, w, k, N, u],
                    );
                return null == C
                    ? null
                    : r.createElement(
                          pe.Provider,
                          { value: m },
                          r.createElement(me.Provider, {
                              children: i,
                              value: C,
                          }),
                      );
            }
            function De(e) {
                var t = e.children,
                    n = e.location;
                return Se(Ae(t), n);
            }
            var Ue = (function (e) {
                    return (
                        (e[(e.pending = 0)] = "pending"),
                        (e[(e.success = 1)] = "success"),
                        (e[(e.error = 2)] = "error"),
                        e
                    );
                })(Ue || {}),
                Be = new Promise(function () {});
            r.Component;
            function Ae(e, t) {
                void 0 === t && (t = []);
                var n = [];
                return (
                    r.Children.forEach(e, function (e, a) {
                        if (r.isValidElement(e)) {
                            var l = [].concat(f(t), [a]);
                            if (e.type !== r.Fragment) {
                                e.type !== Fe && P(!1),
                                    e.props.index && e.props.children && P(!1);
                                var i = {
                                    id: e.props.id || l.join("-"),
                                    caseSensitive: e.props.caseSensitive,
                                    element: e.props.element,
                                    Component: e.props.Component,
                                    index: e.props.index,
                                    path: e.props.path,
                                    loader: e.props.loader,
                                    action: e.props.action,
                                    errorElement: e.props.errorElement,
                                    ErrorBoundary: e.props.ErrorBoundary,
                                    hasErrorBoundary:
                                        null != e.props.ErrorBoundary ||
                                        null != e.props.errorElement,
                                    shouldRevalidate: e.props.shouldRevalidate,
                                    handle: e.props.handle,
                                    lazy: e.props.lazy,
                                };
                                e.props.children &&
                                    (i.children = Ae(e.props.children, l)),
                                    n.push(i);
                            } else n.push.apply(n, Ae(e.props.children, l));
                        }
                    }),
                    n
                );
            }
            function Ve() {
                return (
                    (Ve = Object.assign
                        ? Object.assign.bind()
                        : function (e) {
                              for (var t = 1; t < arguments.length; t++) {
                                  var n = arguments[t];
                                  for (var r in n)
                                      Object.prototype.hasOwnProperty.call(
                                          n,
                                          r,
                                      ) && (e[r] = n[r]);
                              }
                              return e;
                          }),
                    Ve.apply(this, arguments)
                );
            }
            function $e(e, t) {
                if (null == e) return {};
                var n,
                    r,
                    a = {},
                    l = Object.keys(e);
                for (r = 0; r < l.length; r++)
                    (n = l[r]), t.indexOf(n) >= 0 || (a[n] = e[n]);
                return a;
            }
            new Set([
                "application/x-www-form-urlencoded",
                "multipart/form-data",
                "text/plain",
            ]);
            var He = [
                "onClick",
                "relative",
                "reloadDocument",
                "replace",
                "state",
                "target",
                "to",
                "preventScrollReset",
            ];
            var Ke = a.startTransition;
            function We(e) {
                var t,
                    n = e.basename,
                    a = e.children,
                    l = e.future,
                    i = e.window,
                    o = r.useRef();
                null == o.current &&
                    (o.current =
                        (void 0 === (t = { window: i, v5Compat: !0 }) &&
                            (t = {}),
                        F(
                            function (e, t) {
                                var n = e.location;
                                return M(
                                    "",
                                    {
                                        pathname: n.pathname,
                                        search: n.search,
                                        hash: n.hash,
                                    },
                                    (t.state && t.state.usr) || null,
                                    (t.state && t.state.key) || "default",
                                );
                            },
                            function (e, t) {
                                return "string" === typeof t ? t : R(t);
                            },
                            null,
                            t,
                        )));
                var s = o.current,
                    u = c(
                        r.useState({ action: s.action, location: s.location }),
                        2,
                    ),
                    d = u[0],
                    f = u[1],
                    p = (l || {}).v7_startTransition,
                    m = r.useCallback(
                        function (e) {
                            p && Ke
                                ? Ke(function () {
                                      return f(e);
                                  })
                                : f(e);
                        },
                        [f, p],
                    );
                return (
                    r.useLayoutEffect(
                        function () {
                            return s.listen(m);
                        },
                        [s, m],
                    ),
                    r.createElement(Ie, {
                        basename: n,
                        children: a,
                        location: d.location,
                        navigationType: d.action,
                        navigator: s,
                    })
                );
            }
            var Qe =
                    "undefined" !== typeof window &&
                    "undefined" !== typeof window.document &&
                    "undefined" !== typeof window.document.createElement,
                qe = /^(?:[a-z][a-z0-9+.-]*:|\/\/)/i,
                Ge = r.forwardRef(function (e, t) {
                    var n,
                        a = e.onClick,
                        l = e.relative,
                        i = e.reloadDocument,
                        o = e.replace,
                        s = e.state,
                        u = e.target,
                        c = e.to,
                        d = e.preventScrollReset,
                        f = $e(e, He),
                        p = r.useContext(pe).basename,
                        m = !1;
                    if ("string" === typeof c && qe.test(c) && ((n = c), Qe))
                        try {
                            var h = new URL(window.location.href),
                                v = c.startsWith("//")
                                    ? new URL(h.protocol + c)
                                    : new URL(c),
                                g = X(v.pathname, p);
                            v.origin === h.origin && null != g
                                ? (c = g + v.search + v.hash)
                                : (m = !0);
                        } catch (x) {}
                    var y = (function (e, t) {
                            var n = (void 0 === t ? {} : t).relative;
                            ge() || P(!1);
                            var a = r.useContext(pe),
                                l = a.basename,
                                i = a.navigator,
                                o = ke(e, { relative: n }),
                                s = o.hash,
                                u = o.pathname,
                                c = o.search,
                                d = u;
                            return (
                                "/" !== l && (d = "/" === u ? l : te([l, u])),
                                i.createHref({
                                    pathname: d,
                                    search: c,
                                    hash: s,
                                })
                            );
                        })(c, { relative: l }),
                        b = (function (e, t) {
                            var n = void 0 === t ? {} : t,
                                a = n.target,
                                l = n.replace,
                                i = n.state,
                                o = n.preventScrollReset,
                                s = n.relative,
                                u = xe(),
                                c = ye(),
                                d = ke(e, { relative: s });
                            return r.useCallback(
                                function (t) {
                                    if (
                                        (function (e, t) {
                                            return (
                                                0 === e.button &&
                                                (!t || "_self" === t) &&
                                                !(function (e) {
                                                    return !!(
                                                        e.metaKey ||
                                                        e.altKey ||
                                                        e.ctrlKey ||
                                                        e.shiftKey
                                                    );
                                                })(e)
                                            );
                                        })(t, a)
                                    ) {
                                        t.preventDefault();
                                        var n =
                                            void 0 !== l ? l : R(c) === R(d);
                                        u(e, {
                                            replace: n,
                                            state: i,
                                            preventScrollReset: o,
                                            relative: s,
                                        });
                                    }
                                },
                                [c, u, d, l, i, a, e, o, s],
                            );
                        })(c, {
                            replace: o,
                            state: s,
                            target: u,
                            preventScrollReset: d,
                            relative: l,
                        });
                    return r.createElement(
                        "a",
                        Ve({}, f, {
                            href: n || y,
                            onClick:
                                m || i
                                    ? a
                                    : function (e) {
                                          a && a(e), e.defaultPrevented || b(e);
                                      },
                            ref: t,
                            target: u,
                        }),
                    );
                });
            var Ye, Xe;
            (function (e) {
                (e.UseScrollRestoration = "useScrollRestoration"),
                    (e.UseSubmit = "useSubmit"),
                    (e.UseSubmitFetcher = "useSubmitFetcher"),
                    (e.UseFetcher = "useFetcher");
            })(Ye || (Ye = {})),
                (function (e) {
                    (e.UseFetchers = "useFetchers"),
                        (e.UseScrollRestoration = "useScrollRestoration");
                })(Xe || (Xe = {}));
            var Ze = n(184);
            function Je() {
                return (0, Ze.jsx)("svg", {
                    className: "w-5 h-5 mr-2",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "stroke-current",
                        d: "M19 9V17.8C19 18.9201 19 19.4802 18.782 19.908C18.5903 20.2843 18.2843 20.5903 17.908 20.782C17.4802 21 16.9201 21 15.8 21H8.2C7.07989 21 6.51984 21 6.09202 20.782C5.71569 20.5903 5.40973 20.2843 5.21799 19.908C5 19.4802 5 18.9201 5 17.8V6.2C5 5.07989 5 4.51984 5.21799 4.09202C5.40973 3.71569 5.71569 3.40973 6.09202 3.21799C6.51984 3 7.0799 3 8.2 3H13M19 9L13 3M19 9H14C13.4477 9 13 8.55228 13 8V3",
                        stroke: "#000000",
                        strokeWidth: "2",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function et() {
                return (0, Ze.jsx)("svg", {
                    className: "w-5 h-5 mr-2",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "stroke-current",
                        d: "M9 17H15M9 13H15M9 9H10M13 3H8.2C7.0799 3 6.51984 3 6.09202 3.21799C5.71569 3.40973 5.40973 3.71569 5.21799 4.09202C5 4.51984 5 5.0799 5 6.2V17.8C5 18.9201 5 19.4802 5.21799 19.908C5.40973 20.2843 5.71569 20.5903 6.09202 20.782C6.51984 21 7.0799 21 8.2 21H15.8C16.9201 21 17.4802 21 17.908 20.782C18.2843 20.5903 18.5903 20.2843 18.782 19.908C19 19.4802 19 18.9201 19 17.8V9M13 3L19 9M13 3V7.4C13 7.96005 13 8.24008 13.109 8.45399C13.2049 8.64215 13.3578 8.79513 13.546 8.89101C13.7599 9 14.0399 9 14.6 9H19",
                        stroke: "#000000",
                        strokeWidth: "2",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function tt() {
                return (0, Ze.jsx)("svg", {
                    className: "w-6 h-6 mr-2",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "fill-white",
                        clipRule: "evenodd",
                        d: "m12 3.75c-4.55635 0-8.25 3.69365-8.25 8.25 0 4.5563 3.69365 8.25 8.25 8.25 4.5563 0 8.25-3.6937 8.25-8.25 0-4.55635-3.6937-8.25-8.25-8.25zm-9.75 8.25c0-5.38478 4.36522-9.75 9.75-9.75 5.3848 0 9.75 4.36522 9.75 9.75 0 5.3848-4.3652 9.75-9.75 9.75-5.38478 0-9.75-4.3652-9.75-9.75zm9.75-.75c.4142 0 .75.3358.75.75v3.5c0 .4142-.3358.75-.75.75s-.75-.3358-.75-.75v-3.5c0-.4142.3358-.75.75-.75zm0-3.25c-.5523 0-1 .44772-1 1s.4477 1 1 1h.01c.5523 0 1-.44772 1-1s-.4477-1-1-1z",
                        fill: "#000000",
                        fillRule: "evenodd",
                    }),
                });
            }
            function nt() {
                return (0, Ze.jsx)("svg", {
                    className: "w-6 h-6 mr-2",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "stroke-white",
                        d: "M10 14L13 21L20 4L3 11L6.5 12.5",
                        stroke: "#000000",
                        strokeWidth: "1.5",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function rt() {
                return (0, Ze.jsx)("svg", {
                    className: "w-6 h-6 mr-2",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "stroke-white",
                        d: "M20 10.9696L11.9628 18.5497C10.9782 19.4783 9.64274 20 8.25028 20C6.85782 20 5.52239 19.4783 4.53777 18.5497C3.55315 17.6211 3 16.3616 3 15.0483C3 13.7351 3.55315 12.4756 4.53777 11.547M14.429 6.88674L7.00403 13.8812C6.67583 14.1907 6.49144 14.6106 6.49144 15.0483C6.49144 15.4861 6.67583 15.9059 7.00403 16.2154C7.33224 16.525 7.77738 16.6989 8.24154 16.6989C8.70569 16.6989 9.15083 16.525 9.47904 16.2154L13.502 12.4254M8.55638 7.75692L12.575 3.96687C13.2314 3.34779 14.1217 3 15.05 3C15.9783 3 16.8686 3.34779 17.525 3.96687C18.1814 4.58595 18.5502 5.4256 18.5502 6.30111C18.5502 7.17662 18.1814 8.01628 17.525 8.63535L16.5 9.601",
                        stroke: "#000000",
                        strokeWidth: "1.5",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function at(e) {
                var t = e.isOpen;
                return (0, Ze.jsxs)("svg", {
                    className: "w-4 h-4 mr-2 fill-current ".concat(
                        t ? "rotate-90" : "",
                        " transition-all duration-100 shrink-0",
                    ),
                    viewBox: "0 0 48 48",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: [
                        (0, Ze.jsx)("rect", {
                            width: "48",
                            height: "48",
                            fill: "none",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M19.5,37.4l11.9-12a1.9,1.9,0,0,0,0-2.8l-11.9-12A2,2,0,0,0,16,12h0V36h0a2,2,0,0,0,3.5,1.4Z",
                        }),
                    ],
                });
            }
            function lt(e) {
                var t = e.isOpen;
                return (0, Ze.jsx)("svg", {
                    className: "w-6 h-6 ml-4 ".concat(
                        t ? "rotate-180" : "",
                        " text-white transition-all duration-150",
                    ),
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        d: "M7 10L12 15L17 10",
                        stroke: "#ffffff",
                        strokeWidth: "1.5",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function it() {
                return (0, Ze.jsxs)("svg", {
                    className: "w-4 h-4 fill-white",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: [
                        (0, Ze.jsx)("path", {
                            className: "fill-white",
                            fillRule: "evenodd",
                            clipRule: "evenodd",
                            d: "M21 8C21 6.34315 19.6569 5 18 5H10C8.34315 5 7 6.34315 7 8V20C7 21.6569 8.34315 23 10 23H18C19.6569 23 21 21.6569 21 20V8ZM19 8C19 7.44772 18.5523 7 18 7H10C9.44772 7 9 7.44772 9 8V20C9 20.5523 9.44772 21 10 21H18C18.5523 21 19 20.5523 19 20V8Z",
                            fill: "#0F0F0F",
                        }),
                        (0, Ze.jsx)("path", {
                            className: "fill-white",
                            d: "M6 3H16C16.5523 3 17 2.55228 17 2C17 1.44772 16.5523 1 16 1H6C4.34315 1 3 2.34315 3 4V18C3 18.5523 3.44772 19 4 19C4.55228 19 5 18.5523 5 18V4C5 3.44772 5.44772 3 6 3Z",
                            fill: "#0F0F0F",
                        }),
                    ],
                });
            }
            function ot(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(t),
                    viewBox: "0 0 20 20",
                    version: "1.1",
                    xmlns: "http://www.w3.org/2000/svg",
                    xmlnsXlink: "http://www.w3.org/1999/xlink",
                    children: (0, Ze.jsx)("g", {
                        stroke: "none",
                        strokeWidth: "1",
                        fillRule: "evenodd",
                        children: (0, Ze.jsx)("g", {
                            transform: "translate(-140.000000, -2159.000000)",
                            className: "fill-current",
                            children: (0, Ze.jsx)("g", {
                                transform: "translate(56.000000, 160.000000)",
                                children: (0, Ze.jsx)("path", {
                                    d: "M100.562548,2016.99998 L87.4381713,2016.99998 C86.7317804,2016.99998 86.2101535,2016.30298 86.4765813,2015.66198 C87.7127655,2012.69798 90.6169306,2010.99998 93.9998492,2010.99998 C97.3837885,2010.99998 100.287954,2012.69798 101.524138,2015.66198 C101.790566,2016.30298 101.268939,2016.99998 100.562548,2016.99998 M89.9166645,2004.99998 C89.9166645,2002.79398 91.7489936,2000.99998 93.9998492,2000.99998 C96.2517256,2000.99998 98.0830339,2002.79398 98.0830339,2004.99998 C98.0830339,2007.20598 96.2517256,2008.99998 93.9998492,2008.99998 C91.7489936,2008.99998 89.9166645,2007.20598 89.9166645,2004.99998 M103.955674,2016.63598 C103.213556,2013.27698 100.892265,2010.79798 97.837022,2009.67298 C99.4560048,2008.39598 100.400241,2006.33098 100.053171,2004.06998 C99.6509769,2001.44698 97.4235996,1999.34798 94.7348224,1999.04198 C91.0232075,1998.61898 87.8750721,2001.44898 87.8750721,2004.99998 C87.8750721,2006.88998 88.7692896,2008.57398 90.1636971,2009.67298 C87.1074334,2010.79798 84.7871636,2013.27698 84.044024,2016.63598 C83.7745338,2017.85698 84.7789973,2018.99998 86.0539717,2018.99998 L101.945727,2018.99998 C103.221722,2018.99998 104.226185,2017.85698 103.955674,2016.63598",
                                }),
                            }),
                        }),
                    }),
                });
            }
            function st() {
                return (0, Ze.jsx)("svg", {
                    xmlns: "http://www.w3.org/2000/svg",
                    className: "w-6 h-6",
                    fill: "none",
                    viewBox: "0 0 24 24",
                    stroke: "currentColor",
                    children: (0, Ze.jsx)("path", {
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                        strokeWidth: "2",
                        d: "M4 6h16M4 12h16M4 18h16",
                    }),
                });
            }
            function ut(e) {
                var t = e.size;
                return (0, Ze.jsx)("svg", {
                    style: { pointerEvents: "none" },
                    fill: "#000000",
                    className: "".concat(t, " fill-white"),
                    version: "1.1",
                    xmlns: "http://www.w3.org/2000/svg",
                    xmlnsXlink: "http://www.w3.org/1999/xlink",
                    viewBox: "0 0 460.775 460.775",
                    xmlSpace: "preserve",
                    children: (0, Ze.jsx)("path", {
                        d: "M285.08,230.397L456.218,59.27c6.076-6.077,6.076-15.911,0-21.986L423.511,4.565c-2.913-2.911-6.866-4.55-10.992-4.55 c-4.127-4.127-8.08-4.55-10.993-4.55c-4.127,0-8.08,1.639-10.993,4.55l-171.138,171.14L59.25,4.565c-2.913-2.911-6.866-4.55-10.993-4.55 c-4.126,0-8.08,1.639-10.992,4.55L4.558,37.284c-6.077,6.075-6.077,15.909,0,21.986l171.138,171.128L4.575,401.505 c-6.074,6.077-6.074,15.911,0,21.986l32.709,32.719c2.911,2.911,6.865,4.55,10.992,4.55c4.127,0,8.08-1.639,10.994-4.55 l171.117-171.12l171.118,171.12c2.913,2.911,6.866,4.55,10.993,4.55c4.128,0,8.081-1.639,10.992-4.55l32.709-32.719 c6.074-6.075,6.074-15.909,0-21.986L285.08,230.397z",
                    }),
                });
            }
            function ct(e) {
                var t = e.isOpen;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(
                        t ? "rotate-180" : "rotate-0",
                        " text-white h-10 w-16 text-sm font-medium transition duration-150",
                    ),
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        d: "M7 10L12 15L17 10",
                        stroke: "#ffffff",
                        strokeWidth: "1.5",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function dt() {
                return (0, Ze.jsxs)("svg", {
                    className: "w-80",
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: [
                        (0, Ze.jsx)("path", {
                            d: "M9 17C9.85038 16.3697 10.8846 16 12 16C13.1154 16 14.1496 16.3697 15 17",
                            stroke: "#1C274C",
                            strokeWidth: "1.5",
                            strokeLinecap: "round",
                            className: "stroke-grey-750",
                        }),
                        (0, Ze.jsx)("ellipse", {
                            cx: "15",
                            cy: "10.5",
                            rx: "1",
                            ry: "1.5",
                            fill: "#1C274C",
                            className: "fill-grey-750",
                        }),
                        (0, Ze.jsx)("ellipse", {
                            cx: "9",
                            cy: "10.5",
                            rx: "1",
                            ry: "1.5",
                            fill: "#1C274C",
                            className: "fill-grey-750",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M22 12C22 16.714 22 19.0711 20.5355 20.5355C19.0711 22 16.714 22 12 22C7.28595 22 4.92893 22 3.46447 20.5355C2 19.0711 2 16.714 2 12C2 7.28595 2 4.92893 3.46447 3.46447C4.92893 2 7.28595 2 12 2C16.714 2 19.0711 2 20.5355 3.46447C21.5093 4.43821 21.8356 5.80655 21.9449 8",
                            stroke: "#1C274C",
                            strokeWidth: "1.5",
                            strokeLinecap: "round",
                            className: "stroke-grey-750",
                        }),
                    ],
                });
            }
            function ft(e) {
                var t = e.cls;
                return (0, Ze.jsxs)("svg", {
                    className: "".concat(t, " fill-current"),
                    fill: "#000000",
                    version: "1.1",
                    xmlns: "http://www.w3.org/2000/svg",
                    xmlnsXlink: "http://www.w3.org/1999/xlink",
                    viewBox: "0 0 100 100",
                    enableBackground: "new 0 0 100 100",
                    xmlSpace: "preserve",
                    children: [
                        (0, Ze.jsx)("path", {
                            d: "M46.05,60.163H31.923c-0.836,0-1.513,0.677-1.513,1.513v21.934c0,0.836,0.677,1.513,1.513,1.513H46.05 c0.836,0,1.512-0.677,1.512-1.513V61.675C47.562,60.839,46.885,60.163,46.05,60.163z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M68.077,14.878H53.95c-0.836,0-1.513,0.677-1.513,1.513v67.218c0,0.836,0.677,1.513,1.513,1.513h14.127 c0.836,0,1.513-0.677,1.513-1.513V16.391C69.59,15.555,68.913,14.878,68.077,14.878z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M90.217,35.299H76.09c-0.836,0-1.513,0.677-1.513,1.513v46.797c0,0.836,0.677,1.513,1.513,1.513h14.126 c0.836,0,1.513-0.677,1.513-1.513V36.812C91.729,35.977,91.052,35.299,90.217,35.299z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M23.91,35.299H9.783c-0.836,0-1.513,0.677-1.513,1.513v46.797c0,0.836,0.677,1.513,1.513,1.513H23.91 c0.836,0,1.513-0.677,1.513-1.513V36.812C25.423,35.977,24.746,35.299,23.91,35.299z",
                        }),
                    ],
                });
            }
            function pt() {
                return (0, Ze.jsx)("svg", {
                    className: "w-6 h-6 mr-2",
                    xmlns: "http://www.w3.org/2000/svg",
                    viewBox: "0 0 30 30",
                    version: "1.1",
                    children: (0, Ze.jsx)("g", {
                        transform: "translate(0,-289.0625)",
                        children: (0, Ze.jsx)("path", {
                            className: "fill-current",
                            d: "M 15 3 C 8.3844276 3 3 8.38443 3 15 C 3 21.61557 8.3844276 27 15 27 C 21.615572 27 27 21.61557 27 15 C 27 8.38443 21.615572 3 15 3 z M 15 5 C 20.534692 5 25 9.46531 25 15 C 25 20.53469 20.534692 25 15 25 C 9.4653079 25 5 20.53469 5 15 C 5 9.46531 9.4653079 5 15 5 z M 15 7 C 14.446 7 14 7.446 14 8 L 14 15 C 14 15.554 14.446 16 15 16 L 22 16 C 22.554 16 23 15.554 23 15 C 23 14.446 22.554 14 22 14 L 16 14 L 16 8 C 16 7.446 15.554 7 15 7 z ",
                            transform: "translate(0,289.0625)",
                        }),
                    }),
                });
            }
            function mt() {
                return (0, Ze.jsxs)("svg", {
                    className: "w-6 h-6 mr-2 fill-current",
                    fill: "#000000",
                    version: "1.1",
                    xmlns: "http://www.w3.org/2000/svg",
                    xmlnsXlink: "http://www.w3.org/1999/xlink",
                    viewBox: "0 0 100 100",
                    enableBackground: "new 0 0 100 100",
                    xmlSpace: "preserve",
                    children: [
                        (0, Ze.jsx)("path", {
                            d: "M27.953,46.506c-1.385-2.83-2.117-6.008-2.117-9.192c0-1.743,0.252-3.534,0.768-5.468c0.231-0.87,0.521-1.702,0.847-2.509 c-1.251-0.683-2.626-1.103-4.101-1.103c-5.47,0-9.898,5.153-9.898,11.517c0,4.452,2.176,8.305,5.354,10.222L5.391,56.217 c-0.836,0.393-1.387,1.337-1.387,2.392v10.588c0,1.419,0.991,2.569,2.21,2.569h7.929V60.656c0-3.237,1.802-6.172,4.599-7.481 l10.262-4.779C28.624,47.792,28.273,47.161,27.953,46.506z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M60.137,34.801h34.092v-0.001c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761c0,0,0-0.001,0-0.001 l0-6.43h0c0-0.973-0.789-1.761-1.761-1.761c-0.002,0-0.004,0.001-0.006,0.001v-0.005H56.133c1.614,2.114,2.844,4.627,3.526,7.435 C59.874,33.168,60.03,33.999,60.137,34.801z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M95.996,66.436c0-0.973-0.789-1.761-1.761-1.761c-0.002,0-0.004,0.001-0.006,0.001v-0.005H72.007v7.095v1.994 c0,0.293-0.016,0.582-0.045,0.867h22.267v-0.001c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761l0-0.001 L95.996,66.436L95.996,66.436z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M94.235,44.762c-0.002,0-0.004,0.001-0.006,0.001v-0.005H58.944c-0.159,0.419-0.327,0.836-0.514,1.249 c-0.364,0.802-0.773,1.569-1.224,2.297l10.288,4.908c0.781,0.378,1.473,0.897,2.078,1.503h24.657v-0.001 c0.002,0,0.004,0.001,0.006,0.001c0.973,0,1.761-0.789,1.761-1.761c0,0,0-0.001,0-0.001l0-6.43h0 C95.996,45.55,95.207,44.762,94.235,44.762z",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M65.323,57.702l-11.551-5.51l-4.885-2.33c2.134-1.344,3.866-3.418,5-5.917c0.899-1.984,1.435-4.231,1.435-6.631 c0-1.348-0.213-2.627-0.512-3.863c-1.453-5.983-6.126-10.392-11.736-10.392c-5.504,0-10.106,4.251-11.648,10.065 c-0.356,1.333-0.602,2.72-0.602,4.189c0,2.552,0.596,4.93,1.609,7c1.171,2.4,2.906,4.379,5.018,5.651l-4.678,2.178l-11.926,5.554 c-1.037,0.485-1.717,1.654-1.717,2.959v11.111v1.994c0,1.756,1.224,3.181,2.735,3.181h42.417c1.511,0,2.735-1.424,2.735-3.181 v-1.994V60.656C67.019,59.355,66.349,58.198,65.323,57.702z",
                        }),
                    ],
                });
            }
            function ht() {
                return (0, Ze.jsxs)("svg", {
                    className: "w-6 h-6",
                    viewBox: "0 0 32 32",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: [
                        (0, Ze.jsx)("path", {
                            d: "M30.0014 16.3109C30.0014 15.1598 29.9061 14.3198 29.6998 13.4487H16.2871V18.6442H24.1601C24.0014 19.9354 23.1442 21.8798 21.2394 23.1864L21.2127 23.3604L25.4536 26.58L25.7474 26.6087C28.4458 24.1665 30.0014 20.5731 30.0014 16.3109Z",
                            fill: "#4285F4",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M16.2863 29.9998C20.1434 29.9998 23.3814 28.7553 25.7466 26.6086L21.2386 23.1863C20.0323 24.0108 18.4132 24.5863 16.2863 24.5863C12.5086 24.5863 9.30225 22.1441 8.15929 18.7686L7.99176 18.7825L3.58208 22.127L3.52441 22.2841C5.87359 26.8574 10.699 29.9998 16.2863 29.9998Z",
                            fill: "#34A853",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M8.15964 18.769C7.85806 17.8979 7.68352 16.9645 7.68352 16.0001C7.68352 15.0356 7.85806 14.1023 8.14377 13.2312L8.13578 13.0456L3.67083 9.64746L3.52475 9.71556C2.55654 11.6134 2.00098 13.7445 2.00098 16.0001C2.00098 18.2556 2.55654 20.3867 3.52475 22.2845L8.15964 18.769Z",
                            fill: "#FBBC05",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M16.2864 7.4133C18.9689 7.4133 20.7784 8.54885 21.8102 9.4978L25.8419 5.64C23.3658 3.38445 20.1435 2 16.2864 2C10.699 2 5.8736 5.1422 3.52441 9.71549L8.14345 13.2311C9.30229 9.85555 12.5086 7.4133 16.2864 7.4133Z",
                            fill: "#EB4335",
                        }),
                    ],
                });
            }
            function vt(e) {
                var t = e.cls;
                return (0, Ze.jsxs)("svg", {
                    className: "".concat(
                        t,
                        " text-grey-700 animate-spin-slow fill-indigo-600",
                    ),
                    viewBox: "0 0 100 101",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: [
                        (0, Ze.jsx)("path", {
                            d: "M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z",
                            fill: "currentColor",
                        }),
                        (0, Ze.jsx)("path", {
                            d: "M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z",
                            fill: "currentFill",
                        }),
                    ],
                });
            }
            function gt(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(t, " fill-current"),
                    viewBox: "0 0 20 20",
                    version: "1.1",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "fill-current",
                        d: "M 11 3.2910156 L 10.646484 3.6464844 L 4.2910156 10 L 10.646484 16.353516 L 11 16.708984 L 11.708984 16 L 11.353516 15.646484 L 5.7089844 10 L 11.353516 4.3535156 L 11.708984 4 L 11 3.2910156 z M 15 3.2910156 L 14.646484 3.6464844 L 8.2910156 10 L 14.646484 16.353516 L 15 16.708984 L 15.708984 16 L 15.353516 15.646484 L 9.7089844 10 L 15.353516 4.3535156 L 15.708984 4 L 15 3.2910156 z ",
                    }),
                });
            }
            function yt(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    fill: "#000000",
                    className: "".concat(t, " fill-current"),
                    viewBox: "0 0 56 56",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        d: "M .3321 40.0118 C .3321 41.5117 1.4337 42.5430 3.0977 42.5430 L 8.4415 42.5430 C 12.4727 42.5430 14.9103 41.3711 17.8165 37.9727 L 23.0899 31.8320 L 28.3399 37.9727 C 31.2462 41.3711 33.6603 42.5664 37.7618 42.5664 L 42.0508 42.5664 L 42.0508 47.7695 C 42.0508 49.0352 42.8476 49.8320 44.1604 49.8320 C 44.7229 49.8320 45.3085 49.6211 45.7304 49.2461 L 54.5898 41.9336 C 55.6679 41.0664 55.6448 39.6602 54.5898 38.7930 L 45.7304 31.4336 C 45.3085 31.0586 44.7229 30.8477 44.1604 30.8477 C 42.8476 30.8477 42.0508 31.6445 42.0508 32.9102 L 42.0508 37.4571 L 37.8790 37.4571 C 35.4649 37.4571 33.9649 36.6836 32.0430 34.4571 L 26.4415 27.9180 L 32.0430 21.4024 C 33.9649 19.1524 35.4649 18.3789 37.8790 18.3789 L 42.0508 18.3789 L 42.0508 23.0898 C 42.0508 24.3555 42.8476 25.1524 44.1604 25.1524 C 44.7229 25.1524 45.3085 24.9414 45.7304 24.5664 L 54.5898 17.2539 C 55.6679 16.3867 55.6448 15.0039 54.5898 14.1133 L 45.7304 6.7539 C 45.3085 6.3789 44.7229 6.1680 44.1604 6.1680 C 42.8476 6.1680 42.0508 6.9649 42.0508 8.2305 L 42.0508 13.2930 L 37.7618 13.2930 C 33.6603 13.2930 31.2462 14.4883 28.3399 17.8867 L 23.0899 24.0274 L 17.8165 17.8867 C 14.9103 14.4883 12.4727 13.2930 8.4415 13.2930 L 3.0977 13.2930 C 1.4337 13.2930 .3321 14.3242 .3321 15.8477 C .3321 17.3477 1.4571 18.4024 3.0977 18.4024 L 8.5352 18.4024 C 10.8087 18.4024 12.2384 19.1758 14.1368 21.4024 L 19.7384 27.9180 L 14.1368 34.4571 C 12.2149 36.6836 10.7852 37.4571 8.5352 37.4571 L 3.0977 37.4571 C 1.4571 37.4571 .3321 38.5118 .3321 40.0118 Z",
                    }),
                });
            }
            function bt(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(t, " fill-current"),
                    fill: "#000000",
                    viewBox: "0 0 1920 1920",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        d: "M1703.534 960c0-41.788-3.84-84.48-11.633-127.172l210.184-182.174-199.454-340.856-265.186 88.433c-66.974-55.567-143.323-99.389-223.85-128.415L1158.932 0h-397.78L706.49 269.704c-81.43 29.138-156.423 72.282-223.962 128.414l-265.073-88.32L18 650.654l210.184 182.174C220.39 875.52 216.55 918.212 216.55 960s3.84 84.48 11.633 127.172L18 1269.346l199.454 340.856 265.186-88.433c66.974 55.567 143.322 99.389 223.85 128.415L761.152 1920h397.779l54.663-269.704c81.318-29.138 156.424-72.282 223.963-128.414l265.073 88.433 199.454-340.856-210.184-182.174c7.793-42.805 11.633-85.497 11.633-127.285m-743.492 395.294c-217.976 0-395.294-177.318-395.294-395.294 0-217.976 177.318-395.294 395.294-395.294 217.977 0 395.294 177.318 395.294 395.294 0 217.976-177.317 395.294-395.294 395.294",
                        fillRule: "evenodd",
                    }),
                });
            }
            function xt(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(t, " fill-current"),
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "fill-current",
                        fillRule: "evenodd",
                        clipRule: "evenodd",
                        d: "M14.9523 6.2635L10.4523 18.2635L9.04784 17.7368L13.5478 5.73682L14.9523 6.2635ZM19.1894 12.0001L15.9698 8.78042L17.0304 7.71976L21.3108 12.0001L17.0304 16.2804L15.9698 15.2198L19.1894 12.0001ZM8.03032 15.2198L4.81065 12.0002L8.03032 8.78049L6.96966 7.71983L2.68933 12.0002L6.96966 16.2805L8.03032 15.2198Z",
                        fill: "#080341",
                    }),
                });
            }
            function wt(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(t, " fill-current"),
                    fill: "#000000",
                    viewBox: "0 0 200 200",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        d: "M114,100l49-49a9.9,9.9,0,0,0-14-14L100,86,51,37A9.9,9.9,0,0,0,37,51l49,49L37,149a9.9,9.9,0,0,0,14,14l49-49,49,49a9.9,9.9,0,0,0,14-14Z",
                    }),
                });
            }
            function jt(e) {
                var t = e.cls;
                return (0, Ze.jsx)("svg", {
                    className: "".concat(t),
                    viewBox: "0 0 24 24",
                    fill: "none",
                    xmlns: "http://www.w3.org/2000/svg",
                    children: (0, Ze.jsx)("path", {
                        className: "stroke-current",
                        d: "M4 12.6111L8.92308 17.5L20 6.5",
                        stroke: "#000000",
                        strokeWidth: "2",
                        strokeLinecap: "round",
                        strokeLinejoin: "round",
                    }),
                });
            }
            function kt(e, t) {
                return e.findIndex(function (e) {
                    return G(e, t);
                });
            }
            function St(e) {
                var t = e.name,
                    n = e.onClick;
                return (0, Ze.jsx)("li", {
                    className:
                        "cursor-pointer px-4 py-3 flex items-center hover:bg-grey-800 border-grey-750",
                    onClick: n,
                    children: t,
                });
            }
            function Nt(e) {
                var t = e.label,
                    n = e.isOpen,
                    r = e.onClick;
                return (0, Ze.jsxs)("button", {
                    className:
                        "w-full rounded-md px-3 py-2 border-1 flex items-center justify-between ".concat(
                            n
                                ? "bg-grey-750 hover:bg-grey-700 border-grey-650"
                                : "bg-grey-825 hover:bg-grey-775 border-default",
                            " transition duration-150",
                        ),
                    onClick: r,
                    children: [t, (0, Ze.jsx)(lt, { isOpen: n })],
                });
            }
            function Ct(e) {
                var t = e.initSelected,
                    n = e.itemNames,
                    a = e.button,
                    l = e.onChange,
                    i = c((0, r.useState)(t), 2),
                    o = i[0],
                    s = i[1],
                    u = c((0, r.useState)(!1), 2),
                    d = u[0],
                    f = u[1],
                    p = (0, r.useRef)(null),
                    m = n.map(function (e, t) {
                        return (0, Ze.jsx)(
                            St,
                            {
                                index: t,
                                name: e,
                                onClick: function () {
                                    l && l(t), f(!1), s(t);
                                },
                            },
                            t,
                        );
                    });
                return (
                    (0, r.useEffect)(
                        function () {
                            s(t || -1);
                        },
                        [t],
                    ),
                    (0, r.useEffect)(
                        function () {
                            l && l(o);
                        },
                        [o],
                    ),
                    (0, r.useEffect)(function () {
                        var e = function (e) {
                            p.current && !p.current.contains(e.target) && f(!1);
                        };
                        return (
                            document.addEventListener("click", e),
                            function () {
                                document.removeEventListener("click", e);
                            }
                        );
                    }, []),
                    (a = a || Nt),
                    (0, Ze.jsxs)("div", {
                        className: "relative w-full",
                        ref: p,
                        children: [
                            (0, Ze.jsx)(a, {
                                label: n[-1 === o ? 0 : o],
                                isOpen: d,
                                onClick: function () {
                                    return f(!d);
                                },
                            }),
                            (0, Ze.jsx)("div", {
                                className:
                                    "z-10 absolute overflow-hidden top-12 inset-x-0 ".concat(
                                        d
                                            ? "max-h-60 opacity-100"
                                            : "max-h-0 opacity-0",
                                        " transition-all duration-150",
                                    ),
                                children: (0, Ze.jsx)("div", {
                                    className:
                                        "rounded-md max-h-60 overflow-y-auto border-default border-1",
                                    children: (0, Ze.jsx)("ul", {
                                        className:
                                            "divide-y divide-default bg-grey-875 rounded-md",
                                        children: m,
                                    }),
                                }),
                            }),
                        ],
                    })
                );
            }
            function Et(e) {
                var t = e.routes,
                    n = e.routePatterns,
                    r = e.routeLabels,
                    a = e.button,
                    l = xe(),
                    i = ye(),
                    o = kt(n, i.pathname);
                return (0, Ze.jsx)(Ct, {
                    initSelected: o,
                    button: a,
                    itemNames: r,
                    onChange: function (e) {
                        -1 === e || G(t[e], i.pathname) || l(t[e]);
                    },
                });
            }
            var Lt = Ct,
                _t = {
                    main: "/",
                    contests: "/contests/",
                    info: "/info/",
                    archive: "/archive/",
                    submissions: "/problemset/status/",
                    problems: "/problemset/main/",
                    submission: "/submission/:id/",
                    profile: "/user/profile/:user/",
                    profileSubmissions: "/user/profile/:user/submissions/",
                    profileSettings: "/user/profile/:user/settings/",
                    problem: "/problemset/main/:problem/",
                    problemSubmit: "/problemset/main/:problem/submit/",
                    problemSubmissions:
                        "/problemset/main/:problem/submissions/",
                    problemRanklist: "/problemset/main/:problem/ranklist/",
                    login: "/user/login/",
                    register: "/user/register/",
                },
                Pt = [
                    _t.main,
                    _t.contests,
                    _t.archive,
                    _t.submissions,
                    _t.problems,
                    _t.info,
                ],
                Ot = [
                    "F\u0151oldal",
                    "Versenyek",
                    "Arch\xedvum",
                    "Bek\xfcld\xe9sek",
                    "Feladatok",
                    "Tudnival\xf3k",
                ],
                zt = [_t.profile.replace(":user", "dbence"), _t.main],
                Mt = [_t.profile, _t.main],
                Rt = ["Profil", "Kil\xe9p\xe9s"];
            function Tt(e) {
                var t = e.label,
                    n = e.route,
                    r = e.selected,
                    a = e.horizontal,
                    l = e.onClick;
                return (0, Ze.jsx)("li", {
                    children: (0, Ze.jsx)(Ge, {
                        onClick: l,
                        className: "flex items-center h-full px-4 "
                            .concat(
                                a ? "border-b-3 pt-1" : "border-l-3 p-3",
                                " ",
                            )
                            .concat(
                                r
                                    ? "border-indigo-500 bg-grey-775"
                                    : "border-transparent hover:bg-grey-800",
                            ),
                        to: n,
                        children: t,
                    }),
                });
            }
            function Ft(e) {
                var t = e.isOpen,
                    n = e.onClick;
                return (0, Ze.jsxs)("button", {
                    className:
                        "border-1 border-grey-675 rounded-tl-md rounded-bl-md flex items-center justify-between px-3 py-2 min-w-32 w-full h-full ".concat(
                            t
                                ? "bg-grey-750 hover:bg-grey-700"
                                : "hover:bg-grey-800",
                        ),
                    onClick: n,
                    children: [
                        (0, Ze.jsx)("span", {
                            className: "flex items-center",
                            children: (0, Ze.jsx)("span", {
                                children: "Profil",
                            }),
                        }),
                        (0, Ze.jsx)(lt, { isOpen: t }),
                    ],
                });
            }
            function It() {
                return (0, Ze.jsxs)("div", {
                    className: "flex",
                    children: [
                        (0, Ze.jsx)(Et, {
                            button: Ft,
                            routes: zt,
                            routePatterns: Mt,
                            routeLabels: Rt,
                        }),
                        (0, Ze.jsxs)("div", {
                            className:
                                "px-4 flex items-center justify-center border-1 border-l-0 border-grey-675 rounded-tr-md rounded-br-md",
                            children: [
                                (0, Ze.jsx)("button", {
                                    className:
                                        "px-2 bg-grey-725 rounded-md mr-1",
                                    children: "hu",
                                }),
                                (0, Ze.jsx)("button", {
                                    className:
                                        "px-2 hover:bg-grey-800 rounded-md",
                                    children: "en",
                                }),
                            ],
                        }),
                    ],
                });
            }
            function Dt(e) {
                var t = e.selected,
                    n = e.isOpen,
                    r = e.onClose,
                    a = Pt.map(function (e, n) {
                        return (0, Ze.jsx)(
                            Tt,
                            {
                                label: Ot[n],
                                route: e,
                                selected: n === t,
                                horizontal: !1,
                                onClick: r,
                            },
                            n,
                        );
                    });
                return (0, Ze.jsxs)("aside", {
                    className:
                        "z-20 h-full overflow-hidden lg:hidden fixed right-0 bg-grey-825 border-l-1 border-default ".concat(
                            n ? "w-72 opacity-100" : "w-0 opacity-0",
                            " ease-in-out transition-all duration-200",
                        ),
                    children: [
                        (0, Ze.jsx)("div", {
                            className: "p-3",
                            children: (0, Ze.jsx)("button", {
                                className:
                                    "rounded-full p-3 hover:bg-grey-800 transition duration-200",
                                onClick: r,
                                children: (0, Ze.jsx)(ut, { size: "w-4 h-4" }),
                            }),
                        }),
                        (0, Ze.jsxs)("div", {
                            className: "flex flex-col justify-center",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mx-4 mb-4",
                                    children: (0, Ze.jsx)(It, {}),
                                }),
                                (0, Ze.jsx)("ol", {
                                    className:
                                        "divide-y divide-default border-t border-b border-grey-750",
                                    children: a,
                                }),
                            ],
                        }),
                    ],
                });
            }
            function Ut(e) {
                var t = e.selected,
                    n = e.onOpen,
                    r = Pt.map(function (e, n) {
                        return (0, Ze.jsx)(
                            Tt,
                            {
                                label: Ot[n],
                                route: e,
                                selected: n === t,
                                horizontal: !0,
                            },
                            n,
                        );
                    });
                return (0, Ze.jsx)("div", {
                    className:
                        "z-10 flex justify-center bg-grey-825 border-b-1 border-grey-725 fixed w-full top-0",
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "w-full max-w-7xl flex justify-between items-center",
                        children: [
                            (0, Ze.jsxs)("div", {
                                className: "flex w-full",
                                children: [
                                    (0, Ze.jsx)(Ge, {
                                        to: "/",
                                        className:
                                            "font-semibold text-lg mx-8 my-4",
                                        children: "nJudge",
                                    }),
                                    (0, Ze.jsx)("ol", {
                                        className: "hidden lg:flex",
                                        children: r,
                                    }),
                                    (0, Ze.jsx)("div", {
                                        className:
                                            "w-full hidden lg:flex justify-end mx-4 my-2",
                                        children: (0, Ze.jsx)(It, {}),
                                    }),
                                ],
                            }),
                            (0, Ze.jsx)("div", {
                                className: "lg:hidden mx-4",
                                children: (0, Ze.jsx)("button", {
                                    className:
                                        "rounded-full p-2 hover:bg-grey-800 transition duration-200",
                                    onClick: n,
                                    children: (0, Ze.jsx)(st, {}),
                                }),
                            }),
                        ],
                    }),
                });
            }
            var Bt = function () {
                var e = ye(),
                    t = kt(Pt, e.pathname),
                    n = c((0, r.useState)(!1), 2),
                    a = n[0],
                    l = n[1];
                return (0, Ze.jsxs)("div", {
                    children: [
                        (0, Ze.jsx)(Ut, {
                            selected: t,
                            onOpen: function () {
                                l(!0);
                            },
                        }),
                        (0, Ze.jsx)(Dt, {
                            selected: t,
                            isOpen: a,
                            onClose: function () {
                                l(!1);
                            },
                        }),
                    ],
                });
            };
            var At = function (e) {
                var t = e.children,
                    n = e.title,
                    r = e.titleComponent;
                return (0, Ze.jsx)("div", {
                    className:
                        "bg-grey-800 border-1 rounded-md flex flex-col border-default w-full",
                    children: (0, Ze.jsxs)("div", {
                        className: "flex flex-col",
                        children: [
                            n &&
                                (0, Ze.jsx)("span", {
                                    className:
                                        "font-medium px-6 py-4 text-center border-b-1 border-grey-700",
                                    children: n,
                                }),
                            r,
                            (0, Ze.jsx)("div", {
                                className: "w-full text-dropdown-list",
                                children: t,
                            }),
                        ],
                    }),
                });
            };
            var Vt = function (e) {
                var t = e.children,
                    n = e.title,
                    r = e.titleComponent;
                return (0, Ze.jsx)(At, {
                    title: n,
                    titleComponent: r,
                    children: (0, Ze.jsx)("div", {
                        className:
                            "flex flex-col w-full overflow-x-auto ".concat(
                                n || r
                                    ? "rounded-bl-md rounded-br-md"
                                    : "rounded-md",
                                " text-table",
                            ),
                        children: (0, Ze.jsx)("table", {
                            className:
                                "table-fixed divide-y divide-indigo-600 bg-grey-850 border-collapse",
                            children: t,
                        }),
                    }),
                });
            };
            var $t = function (e) {
                var t = e.data,
                    n = e.title,
                    r = e.titleComponent,
                    a = t.map(function (e, t) {
                        return (0, Ze.jsxs)(
                            "tr",
                            {
                                className: "divide-x divide-grey-700",
                                children: [
                                    (0, Ze.jsx)("td", {
                                        className:
                                            "padding-td-default bg-grey-800 font-medium align-top",
                                        children: e[0],
                                    }),
                                    (0, Ze.jsx)("td", {
                                        className:
                                            "padding-td-default bg-grey-825",
                                        children: e[1],
                                    }),
                                ],
                            },
                            t,
                        );
                    });
                return (0, Ze.jsx)(Vt, {
                    title: n,
                    titleComponent: r,
                    children: (0, Ze.jsx)("tbody", {
                        className: "divide-y divide-default",
                        children: a,
                    }),
                });
            };
            var Ht = function (e) {
                var t = e.title,
                    n = e.svg;
                return (0, Ze.jsxs)("div", {
                    className:
                        "py-3 px-4 border-b border-default font-medium flex items-center justify-center",
                    children: [n, (0, Ze.jsx)("span", { children: t })],
                });
            };
            function Kt(e) {
                var t = e.src;
                return (0, Ze.jsx)("img", {
                    alt: "avatar",
                    className: "object-contain border-1 border-default",
                    src: t,
                });
            }
            function Wt(e) {
                var t = e.src,
                    n = e.username,
                    r = e.rating;
                return (0, Ze.jsx)("div", {
                    className: "flex flex-col items-center",
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "flex flex-col items-center p-8 pb-4 border-1 border-default rounded-md bg-grey-825 w-full",
                        children: [
                            (0, Ze.jsx)(Kt, { src: t }),
                            (0, Ze.jsxs)("div", {
                                className:
                                    "flex justify-center items-center w-full",
                                children: [
                                    (0, Ze.jsx)("span", {
                                        className:
                                            "mt-2 text-md font-medium truncate",
                                        children: (0, Ze.jsx)("a", {
                                            href: "#",
                                            children: n,
                                        }),
                                    }),
                                    (0, Ze.jsx)("span", {
                                        className:
                                            "mt-2 text-2xl font-semibold text-indigo-500 mx-2",
                                        children: "\u2022",
                                    }),
                                    (0, Ze.jsx)("span", {
                                        className: "mt-2 text-md truncate",
                                        children: r,
                                    }),
                                ],
                            }),
                        ],
                    }),
                });
            }
            function Qt(e) {
                var t = e.rating,
                    n = e.score,
                    r = e.solved,
                    a = (0, Ze.jsx)(Ht, {
                        svg: (0, Ze.jsx)(ft, { cls: "w-6 h-6 mr-2" }),
                        title: "Statisztik\xe1k",
                    });
                return (0, Ze.jsx)($t, {
                    data: [
                        ["\xc9rt\xe9kel\xe9s", "".concat(t)],
                        ["Pontsz\xe1m", "".concat(n)],
                        ["Megoldott feladatok", "".concat(r)],
                    ],
                    titleComponent: a,
                });
            }
            function qt(e) {
                var t = e.titleComponent,
                    n = e.submissions.map(function (e, t) {
                        return (0, Ze.jsxs)(
                            "tr",
                            {
                                className: "divide-x divide-default",
                                children: [
                                    (0, Ze.jsx)("td", {
                                        className: "padding-td-default",
                                        children: (0, Ze.jsx)("a", {
                                            className: "link",
                                            href: e[2],
                                            children: e[0],
                                        }),
                                    }),
                                    (0, Ze.jsx)("td", {
                                        className: "padding-td-default",
                                        children: e[1],
                                    }),
                                ],
                            },
                            t,
                        );
                    });
                return (0, Ze.jsx)(Vt, {
                    titleComponent: t,
                    children: (0, Ze.jsx)("tbody", {
                        className: "divide-y divide-default",
                        children: n,
                    }),
                });
            }
            var Gt = function () {
                var e = (0, Ze.jsx)(Ht, {
                    svg: (0, Ze.jsx)(pt, {}),
                    title: "Utols\xf3 bek\xfcld\xe9sek",
                });
                return (0, Ze.jsx)("div", {
                    className: "w-full hidden lg:flex justify-center",
                    children: (0, Ze.jsxs)("div", {
                        className: "flex flex-col bg-grey-900 w-80",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "mb-3",
                                children: (0, Ze.jsx)(Wt, {
                                    src: "/assets/profile.webp",
                                    username: "dbence",
                                    rating: 2350,
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "mb-3",
                                children: (0, Ze.jsx)(Qt, {
                                    rating: 2350,
                                    score: 65.4,
                                    solved: 314,
                                }),
                            }),
                            (0, Ze.jsx)(qt, {
                                titleComponent: e,
                                submissions: [
                                    ["31415", "2023-09-06, 14:23:42", "#"],
                                    ["92653", "2023-09-06, 14:23:42", "#"],
                                    ["58979", "2023-09-06, 14:23:42", "#"],
                                ],
                            }),
                        ],
                    }),
                });
            };
            function Yt(e) {
                var t = e.post,
                    n = c((0, r.useState)(!0), 2),
                    a = n[0],
                    l = n[1],
                    i = t.title,
                    o = t.content,
                    s = t.date;
                return (0, Ze.jsx)(At, {
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "px-6 py-5 sm:px-10 sm:py-8 hover:bg-grey-775 transition duration-200 cursor-pointer rounded-md",
                        onClick: function () {
                            return l(!a);
                        },
                        children: [
                            (0, Ze.jsxs)("div", {
                                className: "flex justify-between items-start",
                                children: [
                                    (0, Ze.jsx)("span", {
                                        className: "text-lg font-semibold",
                                        children: i,
                                    }),
                                    (0, Ze.jsx)("span", {
                                        className: "ml-4 date-label",
                                        children: s,
                                    }),
                                ],
                            }),
                            (0, Ze.jsx)("div", {
                                className: "mt-2 ".concat(a ? "truncate" : ""),
                                children: o,
                            }),
                        ],
                    }),
                });
            }
            var Xt = function (e) {
                var t = e.posts.map(function (e, t) {
                    return (0, Ze.jsx)(
                        "div",
                        {
                            className: "mb-3",
                            children: (0, Ze.jsx)(Yt, { post: e }),
                        },
                        t,
                    );
                });
                return (0, Ze.jsx)("div", { children: t });
            };
            var Zt = function (e) {
                var t = e.data;
                return t && G(_t.main, t.route)
                    ? (0, Ze.jsx)("div", {
                          className: "relative w-full flex justify-center",
                          children: (0, Ze.jsxs)("div", {
                              className: "flex justify-center w-full max-w-7xl",
                              children: [
                                  (0, Ze.jsx)(vt, {
                                      cls: "hidden w-16 h-16 absolute top-1/2 left-1/2",
                                  }),
                                  (0, Ze.jsx)("div", {
                                      className: "ml-0 lg:ml-4",
                                      children: (0, Ze.jsx)(Gt, {
                                          src: "https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg",
                                          username: "dbence",
                                          score: "2550",
                                      }),
                                  }),
                                  (0, Ze.jsx)("div", {
                                      className: "w-full px-4 lg:pl-3",
                                      children: (0, Ze.jsx)(Xt, {
                                          posts: t.posts,
                                      }),
                                  }),
                              ],
                          }),
                      })
                    : (0, Ze.jsx)(Ze.Fragment, {});
            };
            function Jt(e) {
                var t = e.name,
                    n = e.date,
                    r = e.active,
                    a = [
                        (0, Ze.jsx)(
                            "button",
                            {
                                className: "btn-gray mr-1",
                                children: "Megtekint\xe9s",
                            },
                            0,
                        ),
                    ];
                return (
                    r &&
                        a.push(
                            (0, Ze.jsx)(
                                "button",
                                {
                                    className: "btn-indigo ml-1",
                                    children: "Regisztr\xe1ci\xf3",
                                },
                                a.length,
                            ),
                        ),
                    (0, Ze.jsx)(At, {
                        children: (0, Ze.jsxs)("div", {
                            className: "px-6 py-5 sm:px-10 sm:py-8",
                            children: [
                                (0, Ze.jsxs)("div", {
                                    className:
                                        "flex justify-between items-start",
                                    children: [
                                        (0, Ze.jsx)("span", {
                                            className: "text-lg font-semibold",
                                            children: t,
                                        }),
                                        (0, Ze.jsx)("span", {
                                            className: "ml-4 date-label",
                                            children: n,
                                        }),
                                    ],
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "mt-2 flex",
                                    children: a,
                                }),
                            ],
                        }),
                    })
                );
            }
            var en = function (e) {
                var t = e.contestData.map(function (e, t) {
                    return (0, Ze.jsx)(
                        "div",
                        {
                            className: "mb-3",
                            children: (0, Ze.jsx)(Jt, {
                                name: e[0],
                                date: e[1],
                                active: e[2],
                            }),
                        },
                        t,
                    );
                });
                return (0, Ze.jsx)("div", { children: t });
            };
            var tn = function () {
                return (0, Ze.jsx)("div", {
                    className: "w-full flex justify-center",
                    children: (0, Ze.jsxs)("div", {
                        className: "flex justify-center w-full max-w-7xl",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "ml-0 lg:ml-4",
                                children: (0, Ze.jsx)(Gt, {
                                    src: "https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg",
                                    username: "dbence",
                                    score: "2550",
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "w-full px-4 lg:pl-3",
                                children: (0, Ze.jsx)(en, {
                                    contestData: [
                                        [
                                            "Online oktat\xf3 programoz\xf3verseny #3",
                                            "2023-12-23, 14:00",
                                            !0,
                                        ],
                                        [
                                            "Online oktat\xf3 programoz\xf3verseny #2",
                                            "2023-08-23, 14:00",
                                            !1,
                                        ],
                                        [
                                            "Online oktat\xf3 programoz\xf3verseny #1",
                                            "2023-04-23, 14:00",
                                            !1,
                                        ],
                                    ],
                                }),
                            }),
                        ],
                    }),
                });
            };
            function nn(e) {
                var t = e.lang,
                    n = e.command;
                return (0, Ze.jsxs)("tr", {
                    className: "divide-x divide-default ",
                    children: [
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default whitespace-nowrap",
                            children: t,
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default text-white",
                            children: (0, Ze.jsxs)("div", {
                                className: "flex items-center",
                                children: [
                                    (0, Ze.jsx)("button", {
                                        className:
                                            "p-2 mr-2 rounded-md border-1 bg-grey-800 border-grey-725 hover:bg-grey-775 transition duration-200",
                                        onClick: function () {
                                            navigator.clipboard.writeText(n);
                                        },
                                        children: (0, Ze.jsx)(it, {}),
                                    }),
                                    (0, Ze.jsx)("div", {
                                        className:
                                            "flex items-center px-3 py-2 border-1 border-grey-725 rounded-md bg-grey-875",
                                        children: (0, Ze.jsx)("pre", {
                                            children: n,
                                        }),
                                    }),
                                ],
                            }),
                        }),
                    ],
                });
            }
            function rn() {
                var e = [
                        [
                            "C++ (11 / 14 / 17)",
                            "g++ -std=c++<verzi\xf3> -O2 -static -DONLINE_JUDGE main.cpp",
                        ],
                        ["C#", "/usr/bin/mcs -out:main.exe -optimize+ main.cs"],
                        ["Go", "/usr/bin/gccgo main.go"],
                        ["Java", "/usr/bin/javac main.java"],
                        ["Pascal", "/usr/bin/fpc -Mobjfpc -O2 -Xss main.pas"],
                        ["PyPy3", "/usr/bin/pypy3 main.py"],
                        ["Python3", "/usr/bin/python3 main.py"],
                    ].map(function (e, t) {
                        return (0, Ze.jsx)(
                            nn,
                            { lang: e[0], command: e[1] },
                            t,
                        );
                    }),
                    t = (0, Ze.jsx)(Ht, {
                        title: "Ford\xedt\xe1si, futtat\xe1si opci\xf3k",
                        svg: (0, Ze.jsx)(xt, { cls: "w-7 h-7 mr-2" }),
                    });
                return (0, Ze.jsx)(Vt, {
                    titleComponent: t,
                    children: (0, Ze.jsx)("tbody", {
                        className: "divide-y divide-default text-sm",
                        children: e,
                    }),
                });
            }
            var an = function () {
                return (0, Ze.jsx)("div", {
                    className: "w-full flex justify-center",
                    children: (0, Ze.jsxs)("div", {
                        className: "flex justify-center w-full max-w-7xl",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "ml-0 lg:ml-4",
                                children: (0, Ze.jsx)(Gt, {
                                    src: "https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg",
                                    username: "dbence",
                                    score: "2550",
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className:
                                    "w-full px-4 lg:pl-3 overflow-x-auto",
                                children: (0, Ze.jsx)(rn, {}),
                            }),
                        ],
                    }),
                });
            };
            var ln = function e(t) {
                var n = t.tree,
                    a = !n.title,
                    l = c((0, r.useState)(a), 2),
                    i = l[0],
                    o = l[1],
                    s = n.children
                        ? n.children.map(function (t, n) {
                              return (0, Ze.jsx)(
                                  "li",
                                  {
                                      className: "mt-2",
                                      children: (0, Ze.jsx)(e, { tree: t }),
                                  },
                                  n,
                              );
                          })
                        : [],
                    u = (0, Ze.jsx)(at, { isOpen: i }),
                    d = (0, Ze.jsxs)("span", {
                        className:
                            "w-fit flex items-center cursor-pointer hover:text-indigo-300 font-medium transition-all duration-100",
                        onClick: function () {
                            return o(!i);
                        },
                        children: [u, n.title],
                    }),
                    f = (0, Ze.jsx)(Ge, {
                        to: n.link,
                        className:
                            "w-fit flex items-center cursor-pointer hover:text-indigo-300 transition-all duration-100",
                        children: n.title,
                    }),
                    p = !n.children || 0 === n.children.length;
                return (0, Ze.jsxs)("div", {
                    children: [
                        !a && !p && d,
                        !a && p && f,
                        (0, Ze.jsx)("ul", {
                            className: ""
                                .concat(i ? "" : "hidden", " ")
                                .concat(a ? "" : "ml-8", " mb-4"),
                            children: s,
                        }),
                    ],
                });
            };
            var on = function (e) {
                var t = e.title,
                    n = e.tree;
                return (0, Ze.jsx)(At, {
                    title: t,
                    children: (0, Ze.jsx)("div", {
                        className: "rounded-md overflow-hidden",
                        children: (0, Ze.jsx)("div", {
                            className: "px-8 pt-4 pb-2 bg-grey-850",
                            children: (0, Ze.jsx)(ln, { tree: n }),
                        }),
                    }),
                });
            };
            var sn = function (e) {
                var t = e.data;
                if (!t || !G(_t.archive, t.route))
                    return (0, Ze.jsx)(Ze.Fragment, {});
                var n = t.categories.map(function (e, t) {
                    return (0, Ze.jsx)(
                        "div",
                        {
                            className: "mb-3",
                            children: (0, Ze.jsx)(on, {
                                title: e.title,
                                tree: { children: e.children },
                            }),
                        },
                        t,
                    );
                });
                return (0, Ze.jsx)("div", {
                    className: "relative w-full flex justify-center",
                    children: (0, Ze.jsxs)("div", {
                        className: "flex justify-center w-full max-w-7xl",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "ml-0 lg:ml-4",
                                children: (0, Ze.jsx)(Gt, {
                                    src: "https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg",
                                    username: "dbence",
                                    score: "2550",
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "w-full px-4 lg:pl-3",
                                children: n,
                            }),
                        ],
                    }),
                });
            };
            function un(e, t, n) {
                return (
                    (t = h(t)) in e
                        ? Object.defineProperty(e, t, {
                              value: n,
                              enumerable: !0,
                              configurable: !0,
                              writable: !0,
                          })
                        : (e[t] = n),
                    e
                );
            }
            function cn(e, t) {
                var n = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    t &&
                        (r = r.filter(function (t) {
                            return Object.getOwnPropertyDescriptor(e, t)
                                .enumerable;
                        })),
                        n.push.apply(n, r);
                }
                return n;
            }
            function dn(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var n = null != arguments[t] ? arguments[t] : {};
                    t % 2
                        ? cn(Object(n), !0).forEach(function (t) {
                              un(e, t, n[t]);
                          })
                        : Object.getOwnPropertyDescriptors
                        ? Object.defineProperties(
                              e,
                              Object.getOwnPropertyDescriptors(n),
                          )
                        : cn(Object(n)).forEach(function (t) {
                              Object.defineProperty(
                                  e,
                                  t,
                                  Object.getOwnPropertyDescriptor(n, t),
                              );
                          });
                }
                return e;
            }
            var fn = "%[a-f0-9]{2}",
                pn = new RegExp("(" + fn + ")|([^%]+?)", "gi"),
                mn = new RegExp("(" + fn + ")+", "gi");
            function hn(e, t) {
                try {
                    return [decodeURIComponent(e.join(""))];
                } catch (a) {}
                if (1 === e.length) return e;
                t = t || 1;
                var n = e.slice(0, t),
                    r = e.slice(t);
                return Array.prototype.concat.call([], hn(n), hn(r));
            }
            function vn(e) {
                try {
                    return decodeURIComponent(e);
                } catch (r) {
                    for (var t = e.match(pn) || [], n = 1; n < t.length; n++)
                        t = (e = hn(t, n).join("")).match(pn) || [];
                    return e;
                }
            }
            function gn(e) {
                if ("string" !== typeof e)
                    throw new TypeError(
                        "Expected `encodedURI` to be of type `string`, got `" +
                            typeof e +
                            "`",
                    );
                try {
                    return decodeURIComponent(e);
                } catch (t) {
                    return (function (e) {
                        for (
                            var t = {
                                    "%FE%FF": "\ufffd\ufffd",
                                    "%FF%FE": "\ufffd\ufffd",
                                },
                                n = mn.exec(e);
                            n;

                        ) {
                            try {
                                t[n[0]] = decodeURIComponent(n[0]);
                            } catch (o) {
                                var r = vn(n[0]);
                                r !== n[0] && (t[n[0]] = r);
                            }
                            n = mn.exec(e);
                        }
                        t["%C2"] = "\ufffd";
                        for (var a = 0, l = Object.keys(t); a < l.length; a++) {
                            var i = l[a];
                            e = e.replace(new RegExp(i, "g"), t[i]);
                        }
                        return e;
                    })(e);
                }
            }
            function yn(e, t) {
                if ("string" !== typeof e || "string" !== typeof t)
                    throw new TypeError(
                        "Expected the arguments to be of type `string`",
                    );
                if ("" === e || "" === t) return [];
                var n = e.indexOf(t);
                return -1 === n ? [] : [e.slice(0, n), e.slice(n + t.length)];
            }
            function bn(e, t) {
                var n = {};
                if (Array.isArray(t)) {
                    var r,
                        a = C(t);
                    try {
                        for (a.s(); !(r = a.n()).done; ) {
                            var l = r.value,
                                i = Object.getOwnPropertyDescriptor(e, l);
                            null !== i &&
                                void 0 !== i &&
                                i.enumerable &&
                                Object.defineProperty(n, l, i);
                        }
                    } catch (d) {
                        a.e(d);
                    } finally {
                        a.f();
                    }
                } else {
                    var o,
                        s = C(Reflect.ownKeys(e));
                    try {
                        for (s.s(); !(o = s.n()).done; ) {
                            var u = o.value,
                                c = Object.getOwnPropertyDescriptor(e, u);
                            if (c.enumerable)
                                t(u, e[u], e) && Object.defineProperty(n, u, c);
                        }
                    } catch (d) {
                        s.e(d);
                    } finally {
                        s.f();
                    }
                }
                return n;
            }
            var xn = function (e) {
                    return null === e || void 0 === e;
                },
                wn = function (e) {
                    return encodeURIComponent(e).replace(
                        /[!'()*]/g,
                        function (e) {
                            return "%".concat(
                                e.charCodeAt(0).toString(16).toUpperCase(),
                            );
                        },
                    );
                },
                jn = Symbol("encodeFragmentIdentifier");
            function kn(e) {
                if ("string" !== typeof e || 1 !== e.length)
                    throw new TypeError(
                        "arrayFormatSeparator must be single character string",
                    );
            }
            function Sn(e, t) {
                return t.encode
                    ? t.strict
                        ? wn(e)
                        : encodeURIComponent(e)
                    : e;
            }
            function Nn(e, t) {
                return t.decode ? gn(e) : e;
            }
            function Cn(e) {
                return Array.isArray(e)
                    ? e.sort()
                    : "object" === typeof e
                    ? Cn(Object.keys(e))
                          .sort(function (e, t) {
                              return Number(e) - Number(t);
                          })
                          .map(function (t) {
                              return e[t];
                          })
                    : e;
            }
            function En(e) {
                var t = e.indexOf("#");
                return -1 !== t && (e = e.slice(0, t)), e;
            }
            function Ln(e, t) {
                return (
                    t.parseNumbers &&
                    !Number.isNaN(Number(e)) &&
                    "string" === typeof e &&
                    "" !== e.trim()
                        ? (e = Number(e))
                        : !t.parseBooleans ||
                          null === e ||
                          ("true" !== e.toLowerCase() &&
                              "false" !== e.toLowerCase()) ||
                          (e = "true" === e.toLowerCase()),
                    e
                );
            }
            function _n(e) {
                var t = (e = En(e)).indexOf("?");
                return -1 === t ? "" : e.slice(t + 1);
            }
            function Pn(e, t) {
                kn(
                    (t = dn(
                        {
                            decode: !0,
                            sort: !0,
                            arrayFormat: "none",
                            arrayFormatSeparator: ",",
                            parseNumbers: !1,
                            parseBooleans: !1,
                        },
                        t,
                    )).arrayFormatSeparator,
                );
                var n = (function (e) {
                        var t;
                        switch (e.arrayFormat) {
                            case "index":
                                return function (e, n, r) {
                                    (t = /\[(\d*)]$/.exec(e)),
                                        (e = e.replace(/\[\d*]$/, "")),
                                        t
                                            ? (void 0 === r[e] && (r[e] = {}),
                                              (r[e][t[1]] = n))
                                            : (r[e] = n);
                                };
                            case "bracket":
                                return function (e, n, r) {
                                    (t = /(\[])$/.exec(e)),
                                        (e = e.replace(/\[]$/, "")),
                                        t
                                            ? void 0 !== r[e]
                                                ? (r[e] = [].concat(f(r[e]), [
                                                      n,
                                                  ]))
                                                : (r[e] = [n])
                                            : (r[e] = n);
                                };
                            case "colon-list-separator":
                                return function (e, n, r) {
                                    (t = /(:list)$/.exec(e)),
                                        (e = e.replace(/:list$/, "")),
                                        t
                                            ? void 0 !== r[e]
                                                ? (r[e] = [].concat(f(r[e]), [
                                                      n,
                                                  ]))
                                                : (r[e] = [n])
                                            : (r[e] = n);
                                };
                            case "comma":
                            case "separator":
                                return function (t, n, r) {
                                    var a =
                                            "string" === typeof n &&
                                            n.includes(e.arrayFormatSeparator),
                                        l =
                                            "string" === typeof n &&
                                            !a &&
                                            Nn(n, e).includes(
                                                e.arrayFormatSeparator,
                                            );
                                    n = l ? Nn(n, e) : n;
                                    var i =
                                        a || l
                                            ? n
                                                  .split(e.arrayFormatSeparator)
                                                  .map(function (t) {
                                                      return Nn(t, e);
                                                  })
                                            : null === n
                                            ? n
                                            : Nn(n, e);
                                    r[t] = i;
                                };
                            case "bracket-separator":
                                return function (t, n, r) {
                                    var a = /(\[])$/.test(t);
                                    if (((t = t.replace(/\[]$/, "")), a)) {
                                        var l =
                                            null === n
                                                ? []
                                                : n
                                                      .split(
                                                          e.arrayFormatSeparator,
                                                      )
                                                      .map(function (t) {
                                                          return Nn(t, e);
                                                      });
                                        void 0 !== r[t]
                                            ? (r[t] = [].concat(f(r[t]), f(l)))
                                            : (r[t] = l);
                                    } else r[t] = n ? Nn(n, e) : n;
                                };
                            default:
                                return function (e, t, n) {
                                    void 0 !== n[e]
                                        ? (n[e] = [].concat(f([n[e]].flat()), [
                                              t,
                                          ]))
                                        : (n[e] = t);
                                };
                        }
                    })(t),
                    r = Object.create(null);
                if ("string" !== typeof e) return r;
                if (!(e = e.trim().replace(/^[?#&]/, ""))) return r;
                var a,
                    l = C(e.split("&"));
                try {
                    for (l.s(); !(a = l.n()).done; ) {
                        var i = a.value;
                        if ("" !== i) {
                            var o = t.decode ? i.replace(/\+/g, " ") : i,
                                s = c(yn(o, "="), 2),
                                u = s[0],
                                d = s[1];
                            void 0 === u && (u = o),
                                (d =
                                    void 0 === d
                                        ? null
                                        : [
                                              "comma",
                                              "separator",
                                              "bracket-separator",
                                          ].includes(t.arrayFormat)
                                        ? d
                                        : Nn(d, t)),
                                n(Nn(u, t), d, r);
                        }
                    }
                } catch (k) {
                    l.e(k);
                } finally {
                    l.f();
                }
                for (var p = 0, m = Object.entries(r); p < m.length; p++) {
                    var h = c(m[p], 2),
                        v = h[0],
                        g = h[1];
                    if ("object" === typeof g && null !== g)
                        for (
                            var y = 0, b = Object.entries(g);
                            y < b.length;
                            y++
                        ) {
                            var x = c(b[y], 2),
                                w = x[0],
                                j = x[1];
                            g[w] = Ln(j, t);
                        }
                    else r[v] = Ln(g, t);
                }
                return !1 === t.sort
                    ? r
                    : (!0 === t.sort
                          ? Object.keys(r).sort()
                          : Object.keys(r).sort(t.sort)
                      ).reduce(function (e, t) {
                          var n = r[t];
                          return (
                              Boolean(n) &&
                              "object" === typeof n &&
                              !Array.isArray(n)
                                  ? (e[t] = Cn(n))
                                  : (e[t] = n),
                              e
                          );
                      }, Object.create(null));
            }
            function On(e, t) {
                if (!e) return "";
                kn(
                    (t = dn(
                        {
                            encode: !0,
                            strict: !0,
                            arrayFormat: "none",
                            arrayFormatSeparator: ",",
                        },
                        t,
                    )).arrayFormatSeparator,
                );
                for (
                    var n = function (n) {
                            return (
                                (t.skipNull && xn(e[n])) ||
                                (t.skipEmptyString && "" === e[n])
                            );
                        },
                        r = (function (e) {
                            switch (e.arrayFormat) {
                                case "index":
                                    return function (t) {
                                        return function (n, r) {
                                            var a = n.length;
                                            return void 0 === r ||
                                                (e.skipNull && null === r) ||
                                                (e.skipEmptyString && "" === r)
                                                ? n
                                                : [].concat(
                                                      f(n),
                                                      null === r
                                                          ? [
                                                                [
                                                                    Sn(t, e),
                                                                    "[",
                                                                    a,
                                                                    "]",
                                                                ].join(""),
                                                            ]
                                                          : [
                                                                [
                                                                    Sn(t, e),
                                                                    "[",
                                                                    Sn(a, e),
                                                                    "]=",
                                                                    Sn(r, e),
                                                                ].join(""),
                                                            ],
                                                  );
                                        };
                                    };
                                case "bracket":
                                    return function (t) {
                                        return function (n, r) {
                                            return void 0 === r ||
                                                (e.skipNull && null === r) ||
                                                (e.skipEmptyString && "" === r)
                                                ? n
                                                : [].concat(
                                                      f(n),
                                                      null === r
                                                          ? [
                                                                [
                                                                    Sn(t, e),
                                                                    "[]",
                                                                ].join(""),
                                                            ]
                                                          : [
                                                                [
                                                                    Sn(t, e),
                                                                    "[]=",
                                                                    Sn(r, e),
                                                                ].join(""),
                                                            ],
                                                  );
                                        };
                                    };
                                case "colon-list-separator":
                                    return function (t) {
                                        return function (n, r) {
                                            return void 0 === r ||
                                                (e.skipNull && null === r) ||
                                                (e.skipEmptyString && "" === r)
                                                ? n
                                                : [].concat(
                                                      f(n),
                                                      null === r
                                                          ? [
                                                                [
                                                                    Sn(t, e),
                                                                    ":list=",
                                                                ].join(""),
                                                            ]
                                                          : [
                                                                [
                                                                    Sn(t, e),
                                                                    ":list=",
                                                                    Sn(r, e),
                                                                ].join(""),
                                                            ],
                                                  );
                                        };
                                    };
                                case "comma":
                                case "separator":
                                case "bracket-separator":
                                    var t =
                                        "bracket-separator" === e.arrayFormat
                                            ? "[]="
                                            : "=";
                                    return function (n) {
                                        return function (r, a) {
                                            return void 0 === a ||
                                                (e.skipNull && null === a) ||
                                                (e.skipEmptyString && "" === a)
                                                ? r
                                                : ((a = null === a ? "" : a),
                                                  0 === r.length
                                                      ? [
                                                            [
                                                                Sn(n, e),
                                                                t,
                                                                Sn(a, e),
                                                            ].join(""),
                                                        ]
                                                      : [
                                                            [r, Sn(a, e)].join(
                                                                e.arrayFormatSeparator,
                                                            ),
                                                        ]);
                                        };
                                    };
                                default:
                                    return function (t) {
                                        return function (n, r) {
                                            return void 0 === r ||
                                                (e.skipNull && null === r) ||
                                                (e.skipEmptyString && "" === r)
                                                ? n
                                                : [].concat(
                                                      f(n),
                                                      null === r
                                                          ? [Sn(t, e)]
                                                          : [
                                                                [
                                                                    Sn(t, e),
                                                                    "=",
                                                                    Sn(r, e),
                                                                ].join(""),
                                                            ],
                                                  );
                                        };
                                    };
                            }
                        })(t),
                        a = {},
                        l = 0,
                        i = Object.entries(e);
                    l < i.length;
                    l++
                ) {
                    var o = c(i[l], 2),
                        s = o[0],
                        u = o[1];
                    n(s) || (a[s] = u);
                }
                var d = Object.keys(a);
                return (
                    !1 !== t.sort && d.sort(t.sort),
                    d
                        .map(function (n) {
                            var a = e[n];
                            return void 0 === a
                                ? ""
                                : null === a
                                ? Sn(n, t)
                                : Array.isArray(a)
                                ? 0 === a.length &&
                                  "bracket-separator" === t.arrayFormat
                                    ? Sn(n, t) + "[]"
                                    : a.reduce(r(n), []).join("&")
                                : Sn(n, t) + "=" + Sn(a, t);
                        })
                        .filter(function (e) {
                            return e.length > 0;
                        })
                        .join("&")
                );
            }
            function zn(e, t) {
                var n, r;
                t = dn({ decode: !0 }, t);
                var a = c(yn(e, "#"), 2),
                    l = a[0],
                    i = a[1];
                return (
                    void 0 === l && (l = e),
                    dn(
                        {
                            url:
                                null !==
                                    (n =
                                        null === (r = l) ||
                                        void 0 === r ||
                                        null === (r = r.split("?")) ||
                                        void 0 === r
                                            ? void 0
                                            : r[0]) && void 0 !== n
                                    ? n
                                    : "",
                            query: Pn(_n(e), t),
                        },
                        t && t.parseFragmentIdentifier && i
                            ? { fragmentIdentifier: Nn(i, t) }
                            : {},
                    )
                );
            }
            function Mn(e, t) {
                t = dn(un({ encode: !0, strict: !0 }, jn, !0), t);
                var n = En(e.url).split("?")[0] || "",
                    r = On(dn(dn({}, Pn(_n(e.url), { sort: !1 })), e.query), t);
                r && (r = "?".concat(r));
                var a = (function (e) {
                    var t = "",
                        n = e.indexOf("#");
                    return -1 !== n && (t = e.slice(n)), t;
                })(e.url);
                if (e.fragmentIdentifier) {
                    var l = new URL(n);
                    (l.hash = e.fragmentIdentifier),
                        (a = t[jn] ? l.hash : "#".concat(e.fragmentIdentifier));
                }
                return "".concat(n).concat(r).concat(a);
            }
            function Rn(e, t, n) {
                var r = zn(
                        e,
                        (n = dn(
                            un({ parseFragmentIdentifier: !0 }, jn, !1),
                            n,
                        )),
                    ),
                    a = r.url,
                    l = r.query,
                    i = r.fragmentIdentifier;
                return Mn(
                    { url: a, query: bn(l, t), fragmentIdentifier: i },
                    n,
                );
            }
            function Tn(e, t, n) {
                return Rn(
                    e,
                    Array.isArray(t)
                        ? function (e) {
                              return !t.includes(e);
                          }
                        : function (e, n) {
                              return !t(e, n);
                          },
                    n,
                );
            }
            var Fn = e;
            var In = function (e) {
                var t = e.paginationData,
                    n = t.currentPage,
                    r = t.lastPage,
                    a = ye(),
                    l = xe(),
                    i = function (e) {
                        var t = a.search,
                            n = Fn.parse(t);
                        n.page = e;
                        var r = Fn.stringify(n);
                        l("".concat(a.pathname, "?").concat(r));
                    },
                    o =
                        "px-3 py-1.5 text-sm transition duration-200 border-default border hover:bg-grey-750 text-center";
                return (0, Ze.jsx)(At, {
                    children: (0, Ze.jsxs)("div", {
                        className: "flex justify-center p-3 overflow-x-auto",
                        children: [
                            (0, Ze.jsx)("button", {
                                className: "".concat(
                                    o,
                                    " border-r-0 rounded-l-md",
                                ),
                                onClick: function () {
                                    return i(1);
                                },
                                children: (0, Ze.jsx)(gt, { cls: "w-4 h-4" }),
                            }),
                            n >= 3 &&
                                (0, Ze.jsx)("button", {
                                    className: "".concat(
                                        o,
                                        " hidden lg:block border-r-0",
                                    ),
                                    onClick: function () {
                                        return i(n - 2);
                                    },
                                    children: n - 2,
                                }),
                            n >= 2 &&
                                (0, Ze.jsx)("button", {
                                    className: "".concat(o, " border-r-0"),
                                    onClick: function () {
                                        return i(n - 1);
                                    },
                                    children: n - 1,
                                }),
                            (0, Ze.jsx)("button", {
                                className:
                                    "px-3 py-1.5 text-sm font-medium bg-indigo-600 border-indigo-600 hover:bg-indigo-500 hover:border-indigo-500 transition duration-200 text-center",
                                children: n,
                            }),
                            n <= r - 1 &&
                                (0, Ze.jsx)("button", {
                                    className: "".concat(o, " border-l-0"),
                                    onClick: function () {
                                        return i(n + 1);
                                    },
                                    children: n + 1,
                                }),
                            n <= r - 2 &&
                                (0, Ze.jsx)("button", {
                                    className: "".concat(
                                        o,
                                        " hidden lg:block border-l-0",
                                    ),
                                    onClick: function () {
                                        return i(n + 2);
                                    },
                                    children: n + 2,
                                }),
                            (0, Ze.jsx)("button", {
                                className: "".concat(
                                    o,
                                    " border-l-0 rounded-r-md",
                                ),
                                onClick: function () {
                                    return i(r);
                                },
                                children: (0, Ze.jsx)(gt, {
                                    cls: "w-4 h-4 rotate-180",
                                }),
                            }),
                        ],
                    }),
                });
            };
            function Dn(e) {
                var t = e.submission,
                    n = t.id,
                    r = t.date,
                    a = t.user,
                    l = t.problem,
                    i = t.lang,
                    o = t.verdict,
                    s = t.verdictType,
                    u = t.time,
                    c = t.memory;
                return (0, Ze.jsxs)("tr", {
                    className: "divide-x divide-default",
                    children: [
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsx)(Ge, {
                                className: "link",
                                to: _t.submission.replace(":id", t.id),
                                children: n,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: r,
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsx)(Ge, {
                                className: "link",
                                to: _t.submission.replace(":user", t.user),
                                children: a,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsx)(Ge, {
                                className: "link",
                                to: l.link,
                                children: l.label,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: i,
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsxs)("div", {
                                className: "flex items-center",
                                children: [
                                    0 === s &&
                                        (0, Ze.jsx)(vt, {
                                            cls: "w-5 h-5 mr-2 shrink-0",
                                        }),
                                    1 === s &&
                                        (0, Ze.jsx)(wt, {
                                            cls: "w-5 h-5 text-red-500 mr-2 shrink-0",
                                        }),
                                    2 === s &&
                                        (0, Ze.jsx)(jt, {
                                            cls: "w-5 h-5 text-green-500 mr-2 shrink-0",
                                        }),
                                    (0, Ze.jsx)("span", {
                                        className: "whitespace-nowrap",
                                        children: o,
                                    }),
                                ],
                            }),
                        }),
                        (0, Ze.jsxs)("td", {
                            className: "padding-td-default",
                            children: [u, " ms"],
                        }),
                        (0, Ze.jsxs)("td", {
                            className: "padding-td-default",
                            children: [c, " KiB"],
                        }),
                    ],
                });
            }
            var Un = function (e) {
                var t = e.submissions.map(function (e, t) {
                    return (0, Ze.jsx)(Dn, { submission: e }, t);
                });
                return (0, Ze.jsxs)(Vt, {
                    children: [
                        (0, Ze.jsx)("thead", {
                            className: "bg-grey-800",
                            children: (0, Ze.jsxs)("tr", {
                                className: "divide-x divide-default",
                                children: [
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "ID",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "D\xe1tum",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Felhaszn\xe1l\xf3",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Feladat",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Nyelv",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Verdikt",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Id\u0151",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Mem\xf3ria",
                                    }),
                                ],
                            }),
                        }),
                        (0, Ze.jsx)("tbody", {
                            className: "divide-y divide-default",
                            children: t,
                        }),
                    ],
                });
            };
            var Bn = function (e) {
                var t = e.data;
                return t && G(_t.submissions, t.route)
                    ? (0, Ze.jsx)("div", {
                          className: "relative w-full flex justify-center",
                          children: (0, Ze.jsxs)("div", {
                              className: "flex justify-center w-full max-w-7xl",
                              children: [
                                  (0, Ze.jsx)("div", {
                                      className: "ml-0 lg:ml-4",
                                      children: (0, Ze.jsx)(Gt, {
                                          src: "https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg",
                                          username: "dbence",
                                          score: "2550",
                                      }),
                                  }),
                                  (0, Ze.jsxs)("div", {
                                      className:
                                          "w-full px-4 lg:pl-3 overflow-x-auto",
                                      children: [
                                          (0, Ze.jsx)("div", {
                                              className: "mb-2",
                                              children: (0, Ze.jsx)(Un, {
                                                  submissions: t.submissions,
                                              }),
                                          }),
                                          (0, Ze.jsx)(In, {
                                              paginationData: t.paginationData,
                                          }),
                                      ],
                                  }),
                              ],
                          }),
                      })
                    : (0, Ze.jsx)(Ze.Fragment, {});
            };
            var An = function (e) {
                var t = e.id,
                    n = e.label,
                    a = e.type,
                    l = e.initText,
                    i = e.onChange,
                    o = e.onFocus,
                    s = e.onBlur;
                l || (l = ""), a || (a = "text");
                var u = c((0, r.useState)(!1), 2),
                    d = u[0],
                    f = u[1];
                return (0, Ze.jsxs)("div", {
                    children: [
                        (0, Ze.jsx)("label", {
                            htmlFor: t,
                            className: "text-label",
                            children: n,
                        }),
                        (0, Ze.jsx)("div", {
                            className: "border-b-1 ".concat(
                                d ? "border-indigo-600" : "border-transparent",
                                " w-full mt-1",
                            ),
                            children: (0, Ze.jsx)("div", {
                                className: "border-b-1 ".concat(
                                    d ? "border-indigo-600" : "border-grey-650",
                                    " w-full",
                                ),
                                children: (0, Ze.jsx)("input", {
                                    autoComplete: "off",
                                    id: t,
                                    value: l,
                                    type: a,
                                    onChange: function (e) {
                                        i && i(e.target.value);
                                    },
                                    onFocus: function () {
                                        f(!0), o && o();
                                    },
                                    onBlur: function () {
                                        f(!1), s && s();
                                    },
                                    className:
                                        "py-1.5 px-2 bg-grey-850 border border-b-0 ".concat(
                                            d
                                                ? "border-grey-575"
                                                : "border-grey-650",
                                            " w-full outline-none transition-all duration-200",
                                        ),
                                }),
                            }),
                        }),
                    ],
                });
            };
            function Vn(e) {
                e.index;
                var t = e.itemName,
                    n = e.onClick;
                return (0, Ze.jsx)("li", {
                    className:
                        "cursor-pointer px-4 py-2 flex items-center hover:bg-grey-800 border-grey-750",
                    onMouseDown: n,
                    children: t,
                });
            }
            var $n = function (e) {
                var t = e.id,
                    n = e.label,
                    a = e.itemNames,
                    l = e.fillSelected,
                    i = e.initText,
                    o = e.initSelected,
                    s = e.onChange,
                    u = e.onClick,
                    d = c((0, r.useState)(!1), 2),
                    f = d[0],
                    p = d[1],
                    m = c((0, r.useState)(o || -1), 2),
                    h = m[0],
                    v = m[1],
                    g = c((0, r.useState)(i || ""), 2),
                    y = g[0],
                    b = g[1];
                (0, r.useEffect)(
                    function () {
                        s && s(h, y);
                    },
                    [h, y],
                );
                var x = a
                    .filter(function (e) {
                        return e.toLowerCase().includes(y.toLowerCase());
                    })
                    .map(function (e, t) {
                        return (0, Ze.jsx)(
                            Vn,
                            {
                                index: t,
                                itemName: e,
                                onClick: function () {
                                    l && b(e), v(t), u && u(t, e);
                                },
                            },
                            t,
                        );
                    });
                return (0, Ze.jsxs)("div", {
                    className: "relative",
                    children: [
                        (0, Ze.jsx)(An, {
                            id: t,
                            label: n,
                            initText: y,
                            onChange: function (e) {
                                v(
                                    a
                                        .map(function (e) {
                                            return e.toLowerCase();
                                        })
                                        .indexOf(e.toLowerCase()),
                                ),
                                    b(e);
                            },
                            onFocus: function () {
                                p(!0);
                            },
                            onBlur: function () {
                                p(!1);
                            },
                        }),
                        (0, Ze.jsx)("div", {
                            className:
                                "z-10 absolute overflow-hidden inset-x-0 ".concat(
                                    f ? "max-h-60" : "max-h-0",
                                ),
                            children: (0, Ze.jsx)("div", {
                                className:
                                    "rounded-sm max-h-60 overflow-y-auto border-default ".concat(
                                        x.length > 0 ? "border-1" : "",
                                    ),
                                children: (0, Ze.jsx)("ul", {
                                    className:
                                        "divide-y divide-default bg-grey-875",
                                    children: x,
                                }),
                            }),
                        }),
                    ],
                });
            };
            function Hn(e) {
                var t = e.title,
                    n = e.onClick,
                    a = c((0, r.useState)(!1), 2),
                    l = a[0],
                    i = a[1];
                return (0, Ze.jsxs)("span", {
                    className:
                        "whitespace-nowrap flex items-center cursor-pointer text-sm px-2 py-1 border-1 rounded m-1 hover:bg-grey-700 ".concat(
                            l
                                ? "hover:bg-red-600 hover:border-red-500"
                                : "bg-grey-725 border-grey-650",
                            " transition-all duration-200",
                        ),
                    children: [
                        t,
                        (0, Ze.jsx)("span", {
                            className:
                                "ml-3 rounded-full p-1 hover:bg-red-800 transition-all duration-200",
                            onMouseOver: function () {
                                return i(!0);
                            },
                            onMouseLeave: function () {
                                return i(!1);
                            },
                            onClick: function (e) {
                                e.stopPropagation(), n();
                            },
                            children: (0, Ze.jsx)(ut, { size: "h-2 w-2" }),
                        }),
                    ],
                });
            }
            var Kn = function (e) {
                var t = e.id,
                    n = e.label,
                    a = e.itemNames,
                    l = e.initTags,
                    i = e.onChange,
                    o = c((0, r.useState)(l || []), 2),
                    s = o[0],
                    u = o[1];
                (0, r.useEffect)(
                    function () {
                        return i(s);
                    },
                    [i, s],
                );
                var d = s.map(function (e, t) {
                    return (0, Ze.jsx)(
                        Hn,
                        {
                            title: e,
                            onClick: function () {
                                u(function (t) {
                                    return t.filter(function (t) {
                                        return t !== e;
                                    });
                                });
                            },
                        },
                        t,
                    );
                });
                return (0, Ze.jsxs)("div", {
                    children: [
                        (0, Ze.jsx)($n, {
                            id: t,
                            label: n,
                            itemNames: a.filter(function (e) {
                                return !s.includes(e);
                            }),
                            onClick: function (e, t) {
                                u(function (e) {
                                    return e.includes(t) ? e : e.concat([t]);
                                });
                            },
                        }),
                        (0, Ze.jsx)("div", {
                            className: "".concat(
                                d.length > 0 ? "mt-2" : "",
                                " flex flex-wrap",
                            ),
                            children: d,
                        }),
                    ],
                });
            };
            var Wn = function (e) {
                var t = e.children,
                    n = c((0, r.useState)(!1), 2),
                    a = n[0],
                    l = n[1];
                return (0, Ze.jsxs)(At, {
                    children: [
                        (0, Ze.jsx)("button", {
                            onClick: function () {
                                return l(!a);
                            },
                            className: "w-full ".concat(
                                a
                                    ? "bg-grey-750 hover:bg-grey-725 rounded-tl-md rounded-tr-md"
                                    : "bg-grey-800 hover:bg-grey-775 rounded-md",
                                " transiton-all duration-200 border-default flex items-center justify-center",
                            ),
                            children: (0, Ze.jsx)(ct, { isOpen: a }),
                        }),
                        (0, Ze.jsx)("div", {
                            className: "".concat(
                                a ? "" : "h-0 overflow-hidden",
                            ),
                            children: t,
                        }),
                    ],
                });
            };
            function Qn() {
                var e = c((0, r.useState)(""), 2),
                    t = e[0],
                    n = e[1],
                    a = c((0, r.useState)([]), 2),
                    l = a[0],
                    i = a[1],
                    o = c((0, r.useState)([-1, ""]), 2),
                    s = o[0],
                    u = o[1],
                    d = ye(),
                    f = xe();
                return (0, Ze.jsxs)("div", {
                    className: "w-full",
                    children: [
                        (0, Ze.jsx)("div", {
                            className: "mb-4",
                            children: (0, Ze.jsx)(An, {
                                id: "filterTitle",
                                label: "Feladatc\xedm",
                                initText: t,
                                onChange: function (e) {
                                    n(e);
                                },
                            }),
                        }),
                        (0, Ze.jsx)("div", {
                            className: "mb-4",
                            children: (0, Ze.jsx)(Kn, {
                                id: "filterTags",
                                label: "C\xedmk\xe9k",
                                fillSelected: !1,
                                itemNames: [
                                    "matematika",
                                    "moh\xf3",
                                    "dinamikus programoz\xe1s",
                                    "adatszerkezetek",
                                ],
                                initTags: l,
                                onChange: function (e) {
                                    i(e);
                                },
                            }),
                        }),
                        (0, Ze.jsx)("div", {
                            className: "mb-5",
                            children: (0, Ze.jsx)($n, {
                                id: "filterCategory",
                                label: "Kateg\xf3ria",
                                initText: s[1],
                                initSelected: s[0],
                                fillSelected: !0,
                                itemNames: [
                                    "IOI-CEOI V\xe1logat\xf3 2023",
                                    "IOI-CEOI V\xe1logat\xf3 2023 \u2212 1. fordul\xf3",
                                    "IOI-CEOI V\xe1logat\xf3 2023 \u2212 2. fordul\xf3",
                                    "IOI-CEOI V\xe1logat\xf3 2023 \u2212 3. fordul\xf3",
                                ],
                                onChange: function (e, t) {
                                    u([e, t]);
                                },
                            }),
                        }),
                        (0, Ze.jsxs)("div", {
                            className: "flex justify-center",
                            children: [
                                (0, Ze.jsx)("button", {
                                    className: "mr-1 btn-indigo w-32",
                                    onClick: function () {
                                        var e = Fn.stringify({
                                            title: t,
                                            tags: l.join(","),
                                            category: s,
                                        });
                                        f("".concat(d.pathname, "?").concat(e));
                                    },
                                    children: "Keres",
                                }),
                                (0, Ze.jsx)("button", {
                                    className: "ml-1 btn-gray w-32",
                                    children: "Vissza\xe1ll\xedt",
                                }),
                            ],
                        }),
                    ],
                });
            }
            function qn() {
                return (0, Ze.jsx)(Wn, {
                    children: (0, Ze.jsx)("div", {
                        className: "px-8 py-6 border-t border-default",
                        children: (0, Ze.jsx)(Qn, {}),
                    }),
                });
            }
            function Gn(e) {
                var t = e.problem,
                    n = t.id,
                    r = t.title,
                    a = t.category,
                    l = t.tags,
                    i = t.numSolved,
                    o = l.map(function (e, t) {
                        return (0, Ze.jsx)(
                            "span",
                            { className: "tag", children: e },
                            t,
                        );
                    });
                return (0, Ze.jsxs)("tr", {
                    className: "divide-x divide-default",
                    children: [
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: n,
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsx)(Ge, {
                                className: "link",
                                to: _t.problem.replace(":problem", n),
                                children: r,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsx)(Ge, {
                                className: "link",
                                to: a.link,
                                children: a.label,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsx)("div", {
                                className: "flex flex-wrap",
                                children: o,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default",
                            children: (0, Ze.jsxs)(Ge, {
                                className: "link flex items-center",
                                to: "".concat(
                                    _t.problemSubmissions.replace(
                                        ":problem",
                                        n,
                                    ),
                                    "?ac=1",
                                ),
                                children: [
                                    (0, Ze.jsx)(ot, { cls: "w-4 h-4 mr-1" }),
                                    (0, Ze.jsx)("span", { children: i }),
                                ],
                            }),
                        }),
                    ],
                });
            }
            var Yn = function (e) {
                var t = e.problems.map(function (e, t) {
                    return (0, Ze.jsx)(Gn, { problem: e }, t);
                });
                return (0, Ze.jsxs)(Vt, {
                    children: [
                        (0, Ze.jsx)("thead", {
                            className: "bg-grey-800",
                            children: (0, Ze.jsxs)("tr", {
                                className: "divide-x divide-default",
                                children: [
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Azonos\xedt\xf3",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Feladatc\xedm",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Kateg\xf3ria",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "C\xedmk\xe9k",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Megold\xf3k",
                                    }),
                                ],
                            }),
                        }),
                        (0, Ze.jsx)("tbody", {
                            className: "divide-y divide-default",
                            children: t,
                        }),
                    ],
                });
            };
            var Xn = function (e) {
                var t = e.data;
                return t && G(_t.problems, t.route)
                    ? (0, Ze.jsx)("div", {
                          className: "relative w-full flex justify-center",
                          children: (0, Ze.jsxs)("div", {
                              className: "flex justify-center w-full max-w-7xl",
                              children: [
                                  (0, Ze.jsx)("div", {
                                      className: "ml-0 lg:ml-4",
                                      children: (0, Ze.jsx)(Gt, {
                                          src: "https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg",
                                          username: "dbence",
                                          score: "2550",
                                      }),
                                  }),
                                  (0, Ze.jsx)("div", {
                                      className:
                                          "w-full flex flex-col overflow-x-auto",
                                      children: (0, Ze.jsxs)("div", {
                                          className: "w-full px-4 lg:pl-3",
                                          children: [
                                              (0, Ze.jsx)("div", {
                                                  className: "mb-2",
                                                  children: (0, Ze.jsx)(qn, {}),
                                              }),
                                              (0, Ze.jsx)("div", {
                                                  className: "mb-2",
                                                  children: (0, Ze.jsx)(Yn, {
                                                      problems: t.problems,
                                                  }),
                                              }),
                                              (0, Ze.jsx)(In, {
                                                  paginationData:
                                                      t.paginationData,
                                              }),
                                          ],
                                      }),
                                  }),
                              ],
                          }),
                      })
                    : (0, Ze.jsx)(Ze.Fragment, {});
            };
            function Zn(e) {
                var t = e.index,
                    n = e.numCases,
                    r = e.testCase,
                    a = e.group,
                    l = e.isLastGroup,
                    i = e.isLastCase,
                    o = l && i ? "border-b-0" : "",
                    s = l ? "border-b-0" : "";
                return (0, Ze.jsxs)("tr", {
                    children: [
                        0 === t &&
                            (0, Ze.jsxs)(Ze.Fragment, {
                                children: [
                                    (0, Ze.jsx)("td", {
                                        className:
                                            "padding-td-default border border-t-0 border-divide-col text-center ".concat(
                                                s,
                                            ),
                                        rowSpan: n,
                                        children: (0, Ze.jsxs)("div", {
                                            className:
                                                "flex flex-col justify-center",
                                            children: [
                                                (0, Ze.jsxs)("div", {
                                                    className:
                                                        "flex items-center justify-center mb-2",
                                                    children: [
                                                        a.failed &&
                                                            (0, Ze.jsx)(wt, {
                                                                cls: "w-7 h-7 text-red-500",
                                                            }),
                                                        !a.failed &&
                                                            a.completed &&
                                                            (0, Ze.jsx)(jt, {
                                                                cls: "w-7 h-7 text-indigo-500",
                                                            }),
                                                        !a.failed &&
                                                            !a.completed &&
                                                            (0, Ze.jsx)(vt, {
                                                                cls: "w-7 h-7",
                                                            }),
                                                    ],
                                                }),
                                                a.name,
                                            ],
                                        }),
                                    }),
                                    (0, Ze.jsx)("td", {
                                        className:
                                            "padding-td-default border border-t-0 border-divide-col text-center ".concat(
                                                s,
                                            ),
                                        rowSpan: n,
                                        children: ""
                                            .concat(a.score, " / ")
                                            .concat(a.maxScore),
                                    }),
                                ],
                            }),
                        (0, Ze.jsx)("td", {
                            className:
                                "padding-td-default border border-t-0 border-divide-col ".concat(
                                    o,
                                ),
                            children: r.index,
                        }),
                        1 !== a.scoring &&
                            (0, Ze.jsx)("td", {
                                className:
                                    "padding-td-default border border-t-0 border-divide-col ".concat(
                                        o,
                                    ),
                                colSpan: 2,
                                children: (0, Ze.jsxs)("div", {
                                    className: "flex",
                                    children: [
                                        (0, Ze.jsx)(vt, {
                                            cls: "mr-2 w-5 h-5",
                                        }),
                                        (0, Ze.jsx)("span", {
                                            className: "whitespace-nowrap",
                                            children: r.verdictName,
                                        }),
                                    ],
                                }),
                            }),
                        1 === a.scoring &&
                            (0, Ze.jsxs)(Ze.Fragment, {
                                children: [
                                    (0, Ze.jsx)("td", {
                                        className:
                                            "padding-td-default border border-t-0 border-divide-col ".concat(
                                                o,
                                            ),
                                        children: (0, Ze.jsxs)("div", {
                                            className: "flex items-center",
                                            children: [
                                                (0, Ze.jsx)(wt, {
                                                    cls: "mr-2 w-5 h-5 text-red-500",
                                                }),
                                                (0, Ze.jsx)("span", {
                                                    className:
                                                        "whitespace-nowrap",
                                                    children: r.verdictName,
                                                }),
                                            ],
                                        }),
                                    }),
                                    (0, Ze.jsxs)("td", {
                                        className:
                                            "padding-td-default border border-t-0 border-divide-col whitespace-nowrap ".concat(
                                                o,
                                            ),
                                        children: [r.score, " / ", r.maxScore],
                                    }),
                                ],
                            }),
                        (0, Ze.jsx)("td", {
                            className:
                                "padding-td-default border border-t-0 border-divide-col ".concat(
                                    o,
                                ),
                            children: r.timeSpent,
                        }),
                        (0, Ze.jsx)("td", {
                            className:
                                "padding-td-default border border-t-0 border-r-0 border-divide-col ".concat(
                                    o,
                                ),
                            children: r.memoryUsed,
                        }),
                    ],
                });
            }
            function Jn(e) {
                var t = e.group,
                    n = e.isLast,
                    r = t.testCases,
                    a = r.map(function (e, a) {
                        return (0, Ze.jsx)(
                            Zn,
                            {
                                index: a,
                                numCases: r.length,
                                testCase: e,
                                group: t,
                                isLastGroup: n,
                                isLastCase: a === r.length - 1,
                            },
                            a,
                        );
                    });
                return (0, Ze.jsx)(Ze.Fragment, { children: a });
            }
            var er = function (e) {
                var t = e.submission.testSets[0].groups,
                    n = t.map(function (e, n) {
                        return (0, Ze.jsx)(
                            Jn,
                            { group: e, isLast: n === t.length - 1 },
                            n,
                        );
                    });
                return (0, Ze.jsxs)(Vt, {
                    children: [
                        (0, Ze.jsx)("thead", {
                            className: "bg-grey-800",
                            children: (0, Ze.jsxs)("tr", {
                                className: "divide-x divide-default",
                                children: [
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "R\xe9szfeladat",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "\xd6sszpont",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Teszt",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        colSpan: 2,
                                        children: "Verdikt",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Id\u0151",
                                    }),
                                    (0, Ze.jsx)("th", {
                                        className: "padding-td-default",
                                        children: "Mem\xf3ria",
                                    }),
                                ],
                            }),
                        }),
                        (0, Ze.jsx)("tbody", { children: n }),
                    ],
                });
            };
            function tr(e, t, n) {
                return (
                    t in e
                        ? Object.defineProperty(e, t, {
                              value: n,
                              enumerable: !0,
                              configurable: !0,
                              writable: !0,
                          })
                        : (e[t] = n),
                    e
                );
            }
            function nr(e, t) {
                var n = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    t &&
                        (r = r.filter(function (t) {
                            return Object.getOwnPropertyDescriptor(e, t)
                                .enumerable;
                        })),
                        n.push.apply(n, r);
                }
                return n;
            }
            function rr(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var n = null != arguments[t] ? arguments[t] : {};
                    t % 2
                        ? nr(Object(n), !0).forEach(function (t) {
                              tr(e, t, n[t]);
                          })
                        : Object.getOwnPropertyDescriptors
                        ? Object.defineProperties(
                              e,
                              Object.getOwnPropertyDescriptors(n),
                          )
                        : nr(Object(n)).forEach(function (t) {
                              Object.defineProperty(
                                  e,
                                  t,
                                  Object.getOwnPropertyDescriptor(n, t),
                              );
                          });
                }
                return e;
            }
            function ar(e, t) {
                if (null == e) return {};
                var n,
                    r,
                    a = (function (e, t) {
                        if (null == e) return {};
                        var n,
                            r,
                            a = {},
                            l = Object.keys(e);
                        for (r = 0; r < l.length; r++)
                            (n = l[r]), t.indexOf(n) >= 0 || (a[n] = e[n]);
                        return a;
                    })(e, t);
                if (Object.getOwnPropertySymbols) {
                    var l = Object.getOwnPropertySymbols(e);
                    for (r = 0; r < l.length; r++)
                        (n = l[r]),
                            t.indexOf(n) >= 0 ||
                                (Object.prototype.propertyIsEnumerable.call(
                                    e,
                                    n,
                                ) &&
                                    (a[n] = e[n]));
                }
                return a;
            }
            function lr(e, t) {
                (null == t || t > e.length) && (t = e.length);
                for (var n = 0, r = new Array(t); n < t; n++) r[n] = e[n];
                return r;
            }
            function ir(e, t, n) {
                return (
                    t in e
                        ? Object.defineProperty(e, t, {
                              value: n,
                              enumerable: !0,
                              configurable: !0,
                              writable: !0,
                          })
                        : (e[t] = n),
                    e
                );
            }
            function or(e, t) {
                var n = Object.keys(e);
                if (Object.getOwnPropertySymbols) {
                    var r = Object.getOwnPropertySymbols(e);
                    t &&
                        (r = r.filter(function (t) {
                            return Object.getOwnPropertyDescriptor(e, t)
                                .enumerable;
                        })),
                        n.push.apply(n, r);
                }
                return n;
            }
            function sr(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var n = null != arguments[t] ? arguments[t] : {};
                    t % 2
                        ? or(Object(n), !0).forEach(function (t) {
                              ir(e, t, n[t]);
                          })
                        : Object.getOwnPropertyDescriptors
                        ? Object.defineProperties(
                              e,
                              Object.getOwnPropertyDescriptors(n),
                          )
                        : or(Object(n)).forEach(function (t) {
                              Object.defineProperty(
                                  e,
                                  t,
                                  Object.getOwnPropertyDescriptor(n, t),
                              );
                          });
                }
                return e;
            }
            function ur(e) {
                return function t() {
                    for (
                        var n = this,
                            r = arguments.length,
                            a = new Array(r),
                            l = 0;
                        l < r;
                        l++
                    )
                        a[l] = arguments[l];
                    return a.length >= e.length
                        ? e.apply(this, a)
                        : function () {
                              for (
                                  var e = arguments.length,
                                      r = new Array(e),
                                      l = 0;
                                  l < e;
                                  l++
                              )
                                  r[l] = arguments[l];
                              return t.apply(n, [].concat(a, r));
                          };
                };
            }
            function cr(e) {
                return {}.toString.call(e).includes("Object");
            }
            function dr(e) {
                return "function" === typeof e;
            }
            var fr = ur(function (e, t) {
                    throw new Error(e[t] || e.default);
                })({
                    initialIsRequired: "initial state is required",
                    initialType: "initial state should be an object",
                    initialContent:
                        "initial state shouldn't be an empty object",
                    handlerType: "handler should be an object or a function",
                    handlersType: "all handlers should be a functions",
                    selectorType: "selector should be a function",
                    changeType: "provided value of changes should be an object",
                    changeField:
                        'it seams you want to change a field in the state which is not specified in the "initial" state',
                    default:
                        "an unknown error accured in `state-local` package",
                }),
                pr = {
                    changes: function (e, t) {
                        return (
                            cr(t) || fr("changeType"),
                            Object.keys(t).some(function (t) {
                                return (
                                    (n = e),
                                    (r = t),
                                    !Object.prototype.hasOwnProperty.call(n, r)
                                );
                                var n, r;
                            }) && fr("changeField"),
                            t
                        );
                    },
                    selector: function (e) {
                        dr(e) || fr("selectorType");
                    },
                    handler: function (e) {
                        dr(e) || cr(e) || fr("handlerType"),
                            cr(e) &&
                                Object.values(e).some(function (e) {
                                    return !dr(e);
                                }) &&
                                fr("handlersType");
                    },
                    initial: function (e) {
                        var t;
                        e || fr("initialIsRequired"),
                            cr(e) || fr("initialType"),
                            (t = e),
                            Object.keys(t).length || fr("initialContent");
                    },
                };
            function mr(e, t) {
                return dr(t) ? t(e.current) : t;
            }
            function hr(e, t) {
                return (e.current = sr(sr({}, e.current), t)), t;
            }
            function vr(e, t, n) {
                return (
                    dr(t)
                        ? t(e.current)
                        : Object.keys(n).forEach(function (n) {
                              var r;
                              return null === (r = t[n]) || void 0 === r
                                  ? void 0
                                  : r.call(t, e.current[n]);
                          }),
                    n
                );
            }
            var gr = {
                    create: function (e) {
                        var t =
                            arguments.length > 1 && void 0 !== arguments[1]
                                ? arguments[1]
                                : {};
                        pr.initial(e), pr.handler(t);
                        var n = { current: e },
                            r = ur(vr)(n, t),
                            a = ur(hr)(n),
                            l = ur(pr.changes)(e),
                            i = ur(mr)(n);
                        return [
                            function () {
                                var e =
                                    arguments.length > 0 &&
                                    void 0 !== arguments[0]
                                        ? arguments[0]
                                        : function (e) {
                                              return e;
                                          };
                                return pr.selector(e), e(n.current);
                            },
                            function (e) {
                                !(function () {
                                    for (
                                        var e = arguments.length,
                                            t = new Array(e),
                                            n = 0;
                                        n < e;
                                        n++
                                    )
                                        t[n] = arguments[n];
                                    return function (e) {
                                        return t.reduceRight(function (e, t) {
                                            return t(e);
                                        }, e);
                                    };
                                })(
                                    r,
                                    a,
                                    l,
                                    i,
                                )(e);
                            },
                        ];
                    },
                },
                yr = gr,
                br = {
                    paths: {
                        vs: "https://cdn.jsdelivr.net/npm/monaco-editor@0.36.1/min/vs",
                    },
                };
            var xr = function (e) {
                return function t() {
                    for (
                        var n = this,
                            r = arguments.length,
                            a = new Array(r),
                            l = 0;
                        l < r;
                        l++
                    )
                        a[l] = arguments[l];
                    return a.length >= e.length
                        ? e.apply(this, a)
                        : function () {
                              for (
                                  var e = arguments.length,
                                      r = new Array(e),
                                      l = 0;
                                  l < e;
                                  l++
                              )
                                  r[l] = arguments[l];
                              return t.apply(n, [].concat(a, r));
                          };
                };
            };
            var wr = function (e) {
                return {}.toString.call(e).includes("Object");
            };
            var jr = {
                    configIsRequired: "the configuration object is required",
                    configType: "the configuration object should be an object",
                    default:
                        "an unknown error accured in `@monaco-editor/loader` package",
                    deprecation:
                        "Deprecation warning!\n    You are using deprecated way of configuration.\n\n    Instead of using\n      monaco.config({ urls: { monacoBase: '...' } })\n    use\n      monaco.config({ paths: { vs: '...' } })\n\n    For more please check the link https://github.com/suren-atoyan/monaco-loader#config\n  ",
                },
                kr = xr(function (e, t) {
                    throw new Error(e[t] || e.default);
                })(jr),
                Sr = {
                    config: function (e) {
                        return (
                            e || kr("configIsRequired"),
                            wr(e) || kr("configType"),
                            e.urls
                                ? (console.warn(jr.deprecation),
                                  { paths: { vs: e.urls.monacoBase } })
                                : e
                        );
                    },
                },
                Nr = Sr,
                Cr = function () {
                    for (
                        var e = arguments.length, t = new Array(e), n = 0;
                        n < e;
                        n++
                    )
                        t[n] = arguments[n];
                    return function (e) {
                        return t.reduceRight(function (e, t) {
                            return t(e);
                        }, e);
                    };
                };
            var Er = function e(t, n) {
                    return (
                        Object.keys(n).forEach(function (r) {
                            n[r] instanceof Object &&
                                t[r] &&
                                Object.assign(n[r], e(t[r], n[r]));
                        }),
                        rr(rr({}, t), n)
                    );
                },
                Lr = {
                    type: "cancelation",
                    msg: "operation is manually canceled",
                };
            var _r,
                Pr,
                Or = function (e) {
                    var t = !1,
                        n = new Promise(function (n, r) {
                            e.then(function (e) {
                                return t ? r(Lr) : n(e);
                            }),
                                e.catch(r);
                        });
                    return (
                        (n.cancel = function () {
                            return (t = !0);
                        }),
                        n
                    );
                },
                zr = yr.create({
                    config: br,
                    isInitialized: !1,
                    resolve: null,
                    reject: null,
                    monaco: null,
                }),
                Mr =
                    ((Pr = 2),
                    (function (e) {
                        if (Array.isArray(e)) return e;
                    })((_r = zr)) ||
                        (function (e, t) {
                            if (
                                "undefined" !== typeof Symbol &&
                                Symbol.iterator in Object(e)
                            ) {
                                var n = [],
                                    r = !0,
                                    a = !1,
                                    l = void 0;
                                try {
                                    for (
                                        var i, o = e[Symbol.iterator]();
                                        !(r = (i = o.next()).done) &&
                                        (n.push(i.value), !t || n.length !== t);
                                        r = !0
                                    );
                                } catch (s) {
                                    (a = !0), (l = s);
                                } finally {
                                    try {
                                        r || null == o.return || o.return();
                                    } finally {
                                        if (a) throw l;
                                    }
                                }
                                return n;
                            }
                        })(_r, Pr) ||
                        (function (e, t) {
                            if (e) {
                                if ("string" === typeof e) return lr(e, t);
                                var n = Object.prototype.toString
                                    .call(e)
                                    .slice(8, -1);
                                return (
                                    "Object" === n &&
                                        e.constructor &&
                                        (n = e.constructor.name),
                                    "Map" === n || "Set" === n
                                        ? Array.from(e)
                                        : "Arguments" === n ||
                                          /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(
                                              n,
                                          )
                                        ? lr(e, t)
                                        : void 0
                                );
                            }
                        })(_r, Pr) ||
                        (function () {
                            throw new TypeError(
                                "Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.",
                            );
                        })()),
                Rr = Mr[0],
                Tr = Mr[1];
            function Fr(e) {
                return document.body.appendChild(e);
            }
            function Ir(e) {
                var t = Rr(function (e) {
                        return { config: e.config, reject: e.reject };
                    }),
                    n = (function (e) {
                        var t = document.createElement("script");
                        return e && (t.src = e), t;
                    })("".concat(t.config.paths.vs, "/loader.js"));
                return (
                    (n.onload = function () {
                        return e();
                    }),
                    (n.onerror = t.reject),
                    n
                );
            }
            function Dr() {
                var e = Rr(function (e) {
                        return {
                            config: e.config,
                            resolve: e.resolve,
                            reject: e.reject,
                        };
                    }),
                    t = window.require;
                t.config(e.config),
                    t(
                        ["vs/editor/editor.main"],
                        function (t) {
                            Ur(t), e.resolve(t);
                        },
                        function (t) {
                            e.reject(t);
                        },
                    );
            }
            function Ur(e) {
                Rr().monaco || Tr({ monaco: e });
            }
            var Br = new Promise(function (e, t) {
                    return Tr({ resolve: e, reject: t });
                }),
                Ar = {
                    config: function (e) {
                        var t = Nr.config(e),
                            n = t.monaco,
                            r = ar(t, ["monaco"]);
                        Tr(function (e) {
                            return { config: Er(e.config, r), monaco: n };
                        });
                    },
                    init: function () {
                        var e = Rr(function (e) {
                            return {
                                monaco: e.monaco,
                                isInitialized: e.isInitialized,
                                resolve: e.resolve,
                            };
                        });
                        if (!e.isInitialized) {
                            if ((Tr({ isInitialized: !0 }), e.monaco))
                                return e.resolve(e.monaco), Or(Br);
                            if (window.monaco && window.monaco.editor)
                                return (
                                    Ur(window.monaco),
                                    e.resolve(window.monaco),
                                    Or(Br)
                                );
                            Cr(Fr, Ir)(Dr);
                        }
                        return Or(Br);
                    },
                    __getMonacoInstance: function () {
                        return Rr(function (e) {
                            return e.monaco;
                        });
                    },
                },
                Vr = Ar,
                $r = {
                    wrapper: {
                        display: "flex",
                        position: "relative",
                        textAlign: "initial",
                    },
                    fullWidth: { width: "100%" },
                    hide: { display: "none" },
                },
                Hr = {
                    container: {
                        display: "flex",
                        height: "100%",
                        width: "100%",
                        justifyContent: "center",
                        alignItems: "center",
                    },
                };
            var Kr = function (e) {
                var t = e.children;
                return r.createElement("div", { style: Hr.container }, t);
            };
            var Wr = function (e) {
                    var t = e.width,
                        n = e.height,
                        a = e.isEditorReady,
                        l = e.loading,
                        i = e._ref,
                        o = e.className,
                        s = e.wrapperProps;
                    return r.createElement(
                        "section",
                        dn(
                            {
                                style: dn(
                                    dn({}, $r.wrapper),
                                    {},
                                    { width: t, height: n },
                                ),
                            },
                            s,
                        ),
                        !a && r.createElement(Kr, null, l),
                        r.createElement("div", {
                            ref: i,
                            style: dn(dn({}, $r.fullWidth), !a && $r.hide),
                            className: o,
                        }),
                    );
                },
                Qr = (0, r.memo)(Wr);
            var qr = function (e) {
                (0, r.useEffect)(e, []);
            };
            var Gr = function (e, t) {
                var n =
                        !(arguments.length > 2 && void 0 !== arguments[2]) ||
                        arguments[2],
                    a = (0, r.useRef)(!0);
                (0, r.useEffect)(
                    a.current || !n
                        ? function () {
                              a.current = !1;
                          }
                        : e,
                    t,
                );
            };
            function Yr() {}
            function Xr(e, t, n, r) {
                return (
                    (function (e, t) {
                        return e.editor.getModel(Zr(e, t));
                    })(e, r) ||
                    (function (e, t, n, r) {
                        return e.editor.createModel(
                            t,
                            n,
                            r ? Zr(e, r) : void 0,
                        );
                    })(e, t, n, r)
                );
            }
            function Zr(e, t) {
                return e.Uri.parse(t);
            }
            var Jr = function (e) {
                var t = e.original,
                    n = e.modified,
                    a = e.language,
                    l = e.originalLanguage,
                    i = e.modifiedLanguage,
                    o = e.originalModelPath,
                    s = e.modifiedModelPath,
                    u = e.keepCurrentOriginalModel,
                    d = void 0 !== u && u,
                    f = e.keepCurrentModifiedModel,
                    p = void 0 !== f && f,
                    m = e.theme,
                    h = void 0 === m ? "light" : m,
                    v = e.loading,
                    g = void 0 === v ? "Loading..." : v,
                    y = e.options,
                    b = void 0 === y ? {} : y,
                    x = e.height,
                    w = void 0 === x ? "100%" : x,
                    j = e.width,
                    k = void 0 === j ? "100%" : j,
                    S = e.className,
                    N = e.wrapperProps,
                    C = void 0 === N ? {} : N,
                    E = e.beforeMount,
                    L = void 0 === E ? Yr : E,
                    _ = e.onMount,
                    P = void 0 === _ ? Yr : _,
                    O = c((0, r.useState)(!1), 2),
                    z = O[0],
                    M = O[1],
                    R = c((0, r.useState)(!0), 2),
                    T = R[0],
                    F = R[1],
                    I = (0, r.useRef)(null),
                    D = (0, r.useRef)(null),
                    U = (0, r.useRef)(null),
                    B = (0, r.useRef)(P),
                    A = (0, r.useRef)(L),
                    V = (0, r.useRef)(!1);
                qr(function () {
                    var e = Vr.init();
                    return (
                        e
                            .then(function (e) {
                                return (D.current = e) && F(!1);
                            })
                            .catch(function (e) {
                                return (
                                    "cancelation" !==
                                        (null === e || void 0 === e
                                            ? void 0
                                            : e.type) &&
                                    console.error(
                                        "Monaco initialization: error:",
                                        e,
                                    )
                                );
                            }),
                        function () {
                            return I.current
                                ? (function () {
                                      var e,
                                          t,
                                          n,
                                          r,
                                          a =
                                              null === (e = I.current) ||
                                              void 0 === e
                                                  ? void 0
                                                  : e.getModel();
                                      d ||
                                          (null !== a &&
                                              void 0 !== a &&
                                              null !== (t = a.original) &&
                                              void 0 !== t &&
                                              t.dispose()),
                                          p ||
                                              (null !== a &&
                                                  void 0 !== a &&
                                                  null !== (n = a.modified) &&
                                                  void 0 !== n &&
                                                  n.dispose()),
                                          null === (r = I.current) ||
                                              void 0 === r ||
                                              r.dispose();
                                  })()
                                : e.cancel();
                        }
                    );
                }),
                    Gr(
                        function () {
                            if (I.current && D.current) {
                                var e = I.current.getOriginalEditor(),
                                    n = Xr(
                                        D.current,
                                        t || "",
                                        l || a || "text",
                                        o || "",
                                    );
                                n !== e.getModel() && e.setModel(n);
                            }
                        },
                        [o],
                        z,
                    ),
                    Gr(
                        function () {
                            if (I.current && D.current) {
                                var e = I.current.getModifiedEditor(),
                                    t = Xr(
                                        D.current,
                                        n || "",
                                        i || a || "text",
                                        s || "",
                                    );
                                t !== e.getModel() && e.setModel(t);
                            }
                        },
                        [s],
                        z,
                    ),
                    Gr(
                        function () {
                            var e = I.current.getModifiedEditor();
                            e.getOption(D.current.editor.EditorOption.readOnly)
                                ? e.setValue(n || "")
                                : n !== e.getValue() &&
                                  (e.executeEdits("", [
                                      {
                                          range: e
                                              .getModel()
                                              .getFullModelRange(),
                                          text: n || "",
                                          forceMoveMarkers: !0,
                                      },
                                  ]),
                                  e.pushUndoStop());
                        },
                        [n],
                        z,
                    ),
                    Gr(
                        function () {
                            var e;
                            null === (e = I.current) ||
                                void 0 === e ||
                                null === (e = e.getModel()) ||
                                void 0 === e ||
                                e.original.setValue(t || "");
                        },
                        [t],
                        z,
                    ),
                    Gr(
                        function () {
                            var e = I.current.getModel(),
                                t = e.original,
                                n = e.modified;
                            D.current.editor.setModelLanguage(
                                t,
                                l || a || "text",
                            ),
                                D.current.editor.setModelLanguage(
                                    n,
                                    i || a || "text",
                                );
                        },
                        [a, l, i],
                        z,
                    ),
                    Gr(
                        function () {
                            var e;
                            null === (e = D.current) ||
                                void 0 === e ||
                                e.editor.setTheme(h);
                        },
                        [h],
                        z,
                    ),
                    Gr(
                        function () {
                            var e;
                            null === (e = I.current) ||
                                void 0 === e ||
                                e.updateOptions(b);
                        },
                        [b],
                        z,
                    );
                var $ = (0, r.useCallback)(
                        function () {
                            var e;
                            if (D.current) {
                                A.current(D.current);
                                var r = Xr(
                                        D.current,
                                        t || "",
                                        l || a || "text",
                                        o || "",
                                    ),
                                    u = Xr(
                                        D.current,
                                        n || "",
                                        i || a || "text",
                                        s || "",
                                    );
                                null === (e = I.current) ||
                                    void 0 === e ||
                                    e.setModel({ original: r, modified: u });
                            }
                        },
                        [a, n, i, t, l, o, s],
                    ),
                    H = (0, r.useCallback)(
                        function () {
                            var e;
                            !V.current &&
                                U.current &&
                                ((I.current = D.current.editor.createDiffEditor(
                                    U.current,
                                    dn({ automaticLayout: !0 }, b),
                                )),
                                $(),
                                null !== (e = D.current) &&
                                    void 0 !== e &&
                                    e.editor.setTheme(h),
                                M(!0),
                                (V.current = !0));
                        },
                        [b, h, $],
                    );
                return (
                    (0, r.useEffect)(
                        function () {
                            z && B.current(I.current, D.current);
                        },
                        [z],
                    ),
                    (0, r.useEffect)(
                        function () {
                            !T && !z && H();
                        },
                        [T, z, H],
                    ),
                    r.createElement(Qr, {
                        width: k,
                        height: w,
                        isEditorReady: z,
                        loading: g,
                        _ref: U,
                        className: S,
                        wrapperProps: C,
                    })
                );
            };
            (0, r.memo)(Jr);
            var ea = function (e) {
                    var t = (0, r.useRef)();
                    return (
                        (0, r.useEffect)(
                            function () {
                                t.current = e;
                            },
                            [e],
                        ),
                        t.current
                    );
                },
                ta = new Map();
            var na = function (e) {
                    var t = e.defaultValue,
                        n = e.defaultLanguage,
                        a = e.defaultPath,
                        l = e.value,
                        i = e.language,
                        o = e.path,
                        s = e.theme,
                        u = void 0 === s ? "light" : s,
                        d = e.line,
                        f = e.loading,
                        p = void 0 === f ? "Loading..." : f,
                        m = e.options,
                        h = void 0 === m ? {} : m,
                        v = e.overrideServices,
                        g = void 0 === v ? {} : v,
                        y = e.saveViewState,
                        b = void 0 === y || y,
                        x = e.keepCurrentModel,
                        w = void 0 !== x && x,
                        j = e.width,
                        k = void 0 === j ? "100%" : j,
                        S = e.height,
                        N = void 0 === S ? "100%" : S,
                        C = e.className,
                        E = e.wrapperProps,
                        L = void 0 === E ? {} : E,
                        _ = e.beforeMount,
                        P = void 0 === _ ? Yr : _,
                        O = e.onMount,
                        z = void 0 === O ? Yr : O,
                        M = e.onChange,
                        R = e.onValidate,
                        T = void 0 === R ? Yr : R,
                        F = c((0, r.useState)(!1), 2),
                        I = F[0],
                        D = F[1],
                        U = c((0, r.useState)(!0), 2),
                        B = U[0],
                        A = U[1],
                        V = (0, r.useRef)(null),
                        $ = (0, r.useRef)(null),
                        H = (0, r.useRef)(null),
                        K = (0, r.useRef)(z),
                        W = (0, r.useRef)(P),
                        Q = (0, r.useRef)(),
                        q = (0, r.useRef)(l),
                        G = ea(o),
                        Y = (0, r.useRef)(!1),
                        X = (0, r.useRef)(!1);
                    qr(function () {
                        var e = Vr.init();
                        return (
                            e
                                .then(function (e) {
                                    return (V.current = e) && A(!1);
                                })
                                .catch(function (e) {
                                    return (
                                        "cancelation" !==
                                            (null === e || void 0 === e
                                                ? void 0
                                                : e.type) &&
                                        console.error(
                                            "Monaco initialization: error:",
                                            e,
                                        )
                                    );
                                }),
                            function () {
                                return $.current
                                    ? (function () {
                                          var e, t;
                                          null !== (e = Q.current) &&
                                              void 0 !== e &&
                                              e.dispose(),
                                              w
                                                  ? b &&
                                                    ta.set(
                                                        o,
                                                        $.current.saveViewState(),
                                                    )
                                                  : null ===
                                                        (t =
                                                            $.current.getModel()) ||
                                                    void 0 === t ||
                                                    t.dispose(),
                                              $.current.dispose();
                                      })()
                                    : e.cancel();
                            }
                        );
                    }),
                        Gr(
                            function () {
                                var e,
                                    r,
                                    s,
                                    u,
                                    c = Xr(
                                        V.current,
                                        t || l || "",
                                        n || i || "",
                                        o || a || "",
                                    );
                                c !==
                                    (null === (e = $.current) || void 0 === e
                                        ? void 0
                                        : e.getModel()) &&
                                    (b &&
                                        ta.set(
                                            G,
                                            null === (r = $.current) ||
                                                void 0 === r
                                                ? void 0
                                                : r.saveViewState(),
                                        ),
                                    null !== (s = $.current) &&
                                        void 0 !== s &&
                                        s.setModel(c),
                                    b &&
                                        (null === (u = $.current) ||
                                            void 0 === u ||
                                            u.restoreViewState(ta.get(o))));
                            },
                            [o],
                            I,
                        ),
                        Gr(
                            function () {
                                var e;
                                null === (e = $.current) ||
                                    void 0 === e ||
                                    e.updateOptions(h);
                            },
                            [h],
                            I,
                        ),
                        Gr(
                            function () {
                                !$.current ||
                                    void 0 === l ||
                                    ($.current.getOption(
                                        V.current.editor.EditorOption.readOnly,
                                    )
                                        ? $.current.setValue(l)
                                        : l !== $.current.getValue() &&
                                          ((X.current = !0),
                                          $.current.executeEdits("", [
                                              {
                                                  range: $.current
                                                      .getModel()
                                                      .getFullModelRange(),
                                                  text: l,
                                                  forceMoveMarkers: !0,
                                              },
                                          ]),
                                          $.current.pushUndoStop(),
                                          (X.current = !1)));
                            },
                            [l],
                            I,
                        ),
                        Gr(
                            function () {
                                var e,
                                    t,
                                    n =
                                        null === (e = $.current) || void 0 === e
                                            ? void 0
                                            : e.getModel();
                                n &&
                                    i &&
                                    (null === (t = V.current) ||
                                        void 0 === t ||
                                        t.editor.setModelLanguage(n, i));
                            },
                            [i],
                            I,
                        ),
                        Gr(
                            function () {
                                var e;
                                void 0 !== d &&
                                    (null === (e = $.current) ||
                                        void 0 === e ||
                                        e.revealLine(d));
                            },
                            [d],
                            I,
                        ),
                        Gr(
                            function () {
                                var e;
                                null === (e = V.current) ||
                                    void 0 === e ||
                                    e.editor.setTheme(u);
                            },
                            [u],
                            I,
                        );
                    var Z = (0, r.useCallback)(
                        function () {
                            if (H.current && V.current && !Y.current) {
                                var e;
                                W.current(V.current);
                                var r = o || a,
                                    s = Xr(
                                        V.current,
                                        l || t || "",
                                        n || i || "",
                                        r || "",
                                    );
                                ($.current =
                                    null === (e = V.current) || void 0 === e
                                        ? void 0
                                        : e.editor.create(
                                              H.current,
                                              dn(
                                                  {
                                                      model: s,
                                                      automaticLayout: !0,
                                                  },
                                                  h,
                                              ),
                                              g,
                                          )),
                                    b && $.current.restoreViewState(ta.get(r)),
                                    V.current.editor.setTheme(u),
                                    D(!0),
                                    (Y.current = !0);
                            }
                        },
                        [t, n, a, l, i, o, h, g, b, u],
                    );
                    return (
                        (0, r.useEffect)(
                            function () {
                                I && K.current($.current, V.current);
                            },
                            [I],
                        ),
                        (0, r.useEffect)(
                            function () {
                                !B && !I && Z();
                            },
                            [B, I, Z],
                        ),
                        (q.current = l),
                        (0, r.useEffect)(
                            function () {
                                var e, t;
                                I &&
                                    M &&
                                    (null !== (e = Q.current) &&
                                        void 0 !== e &&
                                        e.dispose(),
                                    (Q.current =
                                        null === (t = $.current) || void 0 === t
                                            ? void 0
                                            : t.onDidChangeModelContent(
                                                  function (e) {
                                                      X.current ||
                                                          M(
                                                              $.current.getValue(),
                                                              e,
                                                          );
                                                  },
                                              )));
                            },
                            [I, M],
                        ),
                        (0, r.useEffect)(
                            function () {
                                if (I) {
                                    var e = V.current.editor.onDidChangeMarkers(
                                        function (e) {
                                            var t,
                                                n =
                                                    null ===
                                                        (t =
                                                            $.current.getModel()) ||
                                                    void 0 === t
                                                        ? void 0
                                                        : t.uri;
                                            if (
                                                n &&
                                                e.find(function (e) {
                                                    return e.path === n.path;
                                                })
                                            ) {
                                                var r =
                                                    V.current.editor.getModelMarkers(
                                                        { resource: n },
                                                    );
                                                null === T ||
                                                    void 0 === T ||
                                                    T(r);
                                            }
                                        },
                                    );
                                    return function () {
                                        null === e ||
                                            void 0 === e ||
                                            e.dispose();
                                    };
                                }
                                return function () {};
                            },
                            [I, T],
                        ),
                        r.createElement(Qr, {
                            width: k,
                            height: N,
                            isEditorReady: I,
                            loading: p,
                            _ref: H,
                            className: C,
                            wrapperProps: L,
                        })
                    );
                },
                ra = (0, r.memo)(na);
            var aa = function () {
                return (0, Ze.jsx)("div", {
                    className: "text-white",
                    children: (0, Ze.jsx)("div", {
                        className: "w-full flex justify-center",
                        children: (0, Ze.jsx)("div", {
                            className: "flex justify-center w-full max-w-7xl",
                            children: (0, Ze.jsxs)("div", {
                                className: "w-full px-4",
                                children: [
                                    (0, Ze.jsx)("div", {
                                        className: "mb-3",
                                        children: (0, Ze.jsx)(Un, {
                                            submissions: [],
                                        }),
                                    }),
                                    (0, Ze.jsx)("div", {
                                        className: "mb-3",
                                        children: (0, Ze.jsx)(ra, {
                                            className:
                                                "border-1 border-default",
                                            height: "60vh",
                                            theme: "vs-dark",
                                            defaultLanguage: "cpp",
                                            options: {
                                                domReadOnly: !0,
                                                readOnly: !0,
                                                fontFamily: "JetBrains Mono",
                                            },
                                            value: '#include <iostream>\nusing namespace std;\n\nint main() {\n    cout << "Hello world" << endl;\n}',
                                        }),
                                    }),
                                    (0, Ze.jsx)(er, {
                                        submission: {
                                            compiled: !0,
                                            feedbackType: 1,
                                            testSets: [
                                                {
                                                    groups: [
                                                        {
                                                            name: "subtask1",
                                                            completed: !0,
                                                            failed: !0,
                                                            score: 0,
                                                            maxScore: 9,
                                                            scoring: 1,
                                                            testCases: [
                                                                {
                                                                    index: 1,
                                                                    verdictName:
                                                                        "Id\u0151limit t\xfall\xe9p\xe9s",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 2,
                                                                },
                                                                {
                                                                    index: 2,
                                                                    verdictName:
                                                                        "Id\u0151limit t\xfall\xe9p\xe9s",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 3,
                                                                },
                                                                {
                                                                    index: 3,
                                                                    verdictName:
                                                                        "Id\u0151limit t\xfall\xe9p\xe9s",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 4,
                                                                },
                                                                {
                                                                    index: 4,
                                                                    verdictName:
                                                                        "Id\u0151limit t\xfall\xe9p\xe9s",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 2,
                                                                },
                                                                {
                                                                    index: 5,
                                                                    verdictName:
                                                                        "Id\u0151limit t\xfall\xe9p\xe9s",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 3,
                                                                },
                                                                {
                                                                    index: 6,
                                                                    verdictName:
                                                                        "Id\u0151limit t\xfall\xe9p\xe9s",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 4,
                                                                },
                                                            ],
                                                        },
                                                        {
                                                            name: "subtask2",
                                                            completed: !1,
                                                            failed: !1,
                                                            score: 0,
                                                            maxScore: 9,
                                                            scoring: 2,
                                                            testCases: [
                                                                {
                                                                    index: 7,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 2,
                                                                    maxScore: 2,
                                                                },
                                                                {
                                                                    index: 8,
                                                                    verdictName:
                                                                        "Futtat\xe1s...",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 3,
                                                                },
                                                                {
                                                                    index: 9,
                                                                    verdictName:
                                                                        "Futtat\xe1s...",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 4,
                                                                },
                                                                {
                                                                    index: 10,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 2,
                                                                    maxScore: 2,
                                                                },
                                                                {
                                                                    index: 11,
                                                                    verdictName:
                                                                        "Futtat\xe1s...",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 3,
                                                                },
                                                                {
                                                                    index: 12,
                                                                    verdictName:
                                                                        "Futtat\xe1s...",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 0,
                                                                    maxScore: 4,
                                                                },
                                                            ],
                                                        },
                                                        {
                                                            name: "subtask3",
                                                            completed: !0,
                                                            failed: !1,
                                                            score: 9,
                                                            maxScore: 9,
                                                            scoring: 2,
                                                            testCases: [
                                                                {
                                                                    index: 13,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 2,
                                                                    maxScore: 2,
                                                                },
                                                                {
                                                                    index: 14,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 3,
                                                                    maxScore: 3,
                                                                },
                                                                {
                                                                    index: 15,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 4,
                                                                    maxScore: 4,
                                                                },
                                                                {
                                                                    index: 16,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 2,
                                                                    maxScore: 2,
                                                                },
                                                                {
                                                                    index: 17,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 3,
                                                                    maxScore: 3,
                                                                },
                                                                {
                                                                    index: 18,
                                                                    verdictName:
                                                                        "Elfogadva",
                                                                    timeSpent:
                                                                        "150 ms",
                                                                    memoryUsed:
                                                                        "15127 KiB",
                                                                    score: 4,
                                                                    maxScore: 4,
                                                                },
                                                            ],
                                                        },
                                                    ],
                                                },
                                            ],
                                        },
                                    }),
                                ],
                            }),
                        }),
                    }),
                });
            };
            function la(e) {
                var t = e.isSelected,
                    n = e.label,
                    r = e.route;
                return (0, Ze.jsx)(Ge, {
                    className: "block rounded-md px-4 py-2 ".concat(
                        t ? "bg-grey-800" : "hover:bg-grey-850",
                    ),
                    to: r,
                    children: n,
                });
            }
            var ia = function (e) {
                var t = e.routes,
                    n = e.routeLabels,
                    r = e.routePatterns,
                    a = e.children,
                    l = kt(r, ye().pathname),
                    i = t.map(function (e, t) {
                        return (0, Ze.jsx)(
                            "div",
                            {
                                className: "mr-1.5",
                                children: (0, Ze.jsx)(
                                    la,
                                    {
                                        isSelected: t === l,
                                        label: n[t],
                                        route: e,
                                    },
                                    t,
                                ),
                            },
                            t,
                        );
                    });
                return (0, Ze.jsxs)("div", {
                    className: "w-full",
                    children: [
                        (0, Ze.jsx)("ul", {
                            className: "hidden sm:flex mb-2",
                            children: i,
                        }),
                        (0, Ze.jsx)("div", {
                            className: "block sm:hidden mb-2",
                            children: (0, Ze.jsx)(Et, {
                                label: "Profil",
                                routes: t,
                                routePatterns: r,
                                routeLabels: n,
                            }),
                        }),
                        (0, Ze.jsx)("div", { children: a }),
                    ],
                });
            };
            function oa() {
                oa = function () {
                    return t;
                };
                var e,
                    t = {},
                    n = Object.prototype,
                    r = n.hasOwnProperty,
                    a =
                        Object.defineProperty ||
                        function (e, t, n) {
                            e[t] = n.value;
                        },
                    l = "function" == typeof Symbol ? Symbol : {},
                    i = l.iterator || "@@iterator",
                    o = l.asyncIterator || "@@asyncIterator",
                    s = l.toStringTag || "@@toStringTag";
                function u(e, t, n) {
                    return (
                        Object.defineProperty(e, t, {
                            value: n,
                            enumerable: !0,
                            configurable: !0,
                            writable: !0,
                        }),
                        e[t]
                    );
                }
                try {
                    u({}, "");
                } catch (e) {
                    u = function (e, t, n) {
                        return (e[t] = n);
                    };
                }
                function c(e, t, n, r) {
                    var l = t && t.prototype instanceof y ? t : y,
                        i = Object.create(l.prototype),
                        o = new O(r || []);
                    return a(i, "_invoke", { value: E(e, n, o) }), i;
                }
                function d(e, t, n) {
                    try {
                        return { type: "normal", arg: e.call(t, n) };
                    } catch (e) {
                        return { type: "throw", arg: e };
                    }
                }
                t.wrap = c;
                var f = "suspendedStart",
                    p = "suspendedYield",
                    h = "executing",
                    v = "completed",
                    g = {};
                function y() {}
                function b() {}
                function x() {}
                var w = {};
                u(w, i, function () {
                    return this;
                });
                var j = Object.getPrototypeOf,
                    k = j && j(j(z([])));
                k && k !== n && r.call(k, i) && (w = k);
                var S = (x.prototype = y.prototype = Object.create(w));
                function N(e) {
                    ["next", "throw", "return"].forEach(function (t) {
                        u(e, t, function (e) {
                            return this._invoke(t, e);
                        });
                    });
                }
                function C(e, t) {
                    function n(a, l, i, o) {
                        var s = d(e[a], e, l);
                        if ("throw" !== s.type) {
                            var u = s.arg,
                                c = u.value;
                            return c && "object" == m(c) && r.call(c, "__await")
                                ? t.resolve(c.__await).then(
                                      function (e) {
                                          n("next", e, i, o);
                                      },
                                      function (e) {
                                          n("throw", e, i, o);
                                      },
                                  )
                                : t.resolve(c).then(
                                      function (e) {
                                          (u.value = e), i(u);
                                      },
                                      function (e) {
                                          return n("throw", e, i, o);
                                      },
                                  );
                        }
                        o(s.arg);
                    }
                    var l;
                    a(this, "_invoke", {
                        value: function (e, r) {
                            function a() {
                                return new t(function (t, a) {
                                    n(e, r, t, a);
                                });
                            }
                            return (l = l ? l.then(a, a) : a());
                        },
                    });
                }
                function E(t, n, r) {
                    var a = f;
                    return function (l, i) {
                        if (a === h)
                            throw new Error("Generator is already running");
                        if (a === v) {
                            if ("throw" === l) throw i;
                            return { value: e, done: !0 };
                        }
                        for (r.method = l, r.arg = i; ; ) {
                            var o = r.delegate;
                            if (o) {
                                var s = L(o, r);
                                if (s) {
                                    if (s === g) continue;
                                    return s;
                                }
                            }
                            if ("next" === r.method) r.sent = r._sent = r.arg;
                            else if ("throw" === r.method) {
                                if (a === f) throw ((a = v), r.arg);
                                r.dispatchException(r.arg);
                            } else
                                "return" === r.method &&
                                    r.abrupt("return", r.arg);
                            a = h;
                            var u = d(t, n, r);
                            if ("normal" === u.type) {
                                if (((a = r.done ? v : p), u.arg === g))
                                    continue;
                                return { value: u.arg, done: r.done };
                            }
                            "throw" === u.type &&
                                ((a = v),
                                (r.method = "throw"),
                                (r.arg = u.arg));
                        }
                    };
                }
                function L(t, n) {
                    var r = n.method,
                        a = t.iterator[r];
                    if (a === e)
                        return (
                            (n.delegate = null),
                            ("throw" === r &&
                                t.iterator.return &&
                                ((n.method = "return"),
                                (n.arg = e),
                                L(t, n),
                                "throw" === n.method)) ||
                                ("return" !== r &&
                                    ((n.method = "throw"),
                                    (n.arg = new TypeError(
                                        "The iterator does not provide a '" +
                                            r +
                                            "' method",
                                    )))),
                            g
                        );
                    var l = d(a, t.iterator, n.arg);
                    if ("throw" === l.type)
                        return (
                            (n.method = "throw"),
                            (n.arg = l.arg),
                            (n.delegate = null),
                            g
                        );
                    var i = l.arg;
                    return i
                        ? i.done
                            ? ((n[t.resultName] = i.value),
                              (n.next = t.nextLoc),
                              "return" !== n.method &&
                                  ((n.method = "next"), (n.arg = e)),
                              (n.delegate = null),
                              g)
                            : i
                        : ((n.method = "throw"),
                          (n.arg = new TypeError(
                              "iterator result is not an object",
                          )),
                          (n.delegate = null),
                          g);
                }
                function _(e) {
                    var t = { tryLoc: e[0] };
                    1 in e && (t.catchLoc = e[1]),
                        2 in e && ((t.finallyLoc = e[2]), (t.afterLoc = e[3])),
                        this.tryEntries.push(t);
                }
                function P(e) {
                    var t = e.completion || {};
                    (t.type = "normal"), delete t.arg, (e.completion = t);
                }
                function O(e) {
                    (this.tryEntries = [{ tryLoc: "root" }]),
                        e.forEach(_, this),
                        this.reset(!0);
                }
                function z(t) {
                    if (t || "" === t) {
                        var n = t[i];
                        if (n) return n.call(t);
                        if ("function" == typeof t.next) return t;
                        if (!isNaN(t.length)) {
                            var a = -1,
                                l = function n() {
                                    for (; ++a < t.length; )
                                        if (r.call(t, a))
                                            return (
                                                (n.value = t[a]),
                                                (n.done = !1),
                                                n
                                            );
                                    return (n.value = e), (n.done = !0), n;
                                };
                            return (l.next = l);
                        }
                    }
                    throw new TypeError(m(t) + " is not iterable");
                }
                return (
                    (b.prototype = x),
                    a(S, "constructor", { value: x, configurable: !0 }),
                    a(x, "constructor", { value: b, configurable: !0 }),
                    (b.displayName = u(x, s, "GeneratorFunction")),
                    (t.isGeneratorFunction = function (e) {
                        var t = "function" == typeof e && e.constructor;
                        return (
                            !!t &&
                            (t === b ||
                                "GeneratorFunction" ===
                                    (t.displayName || t.name))
                        );
                    }),
                    (t.mark = function (e) {
                        return (
                            Object.setPrototypeOf
                                ? Object.setPrototypeOf(e, x)
                                : ((e.__proto__ = x),
                                  u(e, s, "GeneratorFunction")),
                            (e.prototype = Object.create(S)),
                            e
                        );
                    }),
                    (t.awrap = function (e) {
                        return { __await: e };
                    }),
                    N(C.prototype),
                    u(C.prototype, o, function () {
                        return this;
                    }),
                    (t.AsyncIterator = C),
                    (t.async = function (e, n, r, a, l) {
                        void 0 === l && (l = Promise);
                        var i = new C(c(e, n, r, a), l);
                        return t.isGeneratorFunction(n)
                            ? i
                            : i.next().then(function (e) {
                                  return e.done ? e.value : i.next();
                              });
                    }),
                    N(S),
                    u(S, s, "Generator"),
                    u(S, i, function () {
                        return this;
                    }),
                    u(S, "toString", function () {
                        return "[object Generator]";
                    }),
                    (t.keys = function (e) {
                        var t = Object(e),
                            n = [];
                        for (var r in t) n.push(r);
                        return (
                            n.reverse(),
                            function e() {
                                for (; n.length; ) {
                                    var r = n.pop();
                                    if (r in t)
                                        return (e.value = r), (e.done = !1), e;
                                }
                                return (e.done = !0), e;
                            }
                        );
                    }),
                    (t.values = z),
                    (O.prototype = {
                        constructor: O,
                        reset: function (t) {
                            if (
                                ((this.prev = 0),
                                (this.next = 0),
                                (this.sent = this._sent = e),
                                (this.done = !1),
                                (this.delegate = null),
                                (this.method = "next"),
                                (this.arg = e),
                                this.tryEntries.forEach(P),
                                !t)
                            )
                                for (var n in this)
                                    "t" === n.charAt(0) &&
                                        r.call(this, n) &&
                                        !isNaN(+n.slice(1)) &&
                                        (this[n] = e);
                        },
                        stop: function () {
                            this.done = !0;
                            var e = this.tryEntries[0].completion;
                            if ("throw" === e.type) throw e.arg;
                            return this.rval;
                        },
                        dispatchException: function (t) {
                            if (this.done) throw t;
                            var n = this;
                            function a(r, a) {
                                return (
                                    (o.type = "throw"),
                                    (o.arg = t),
                                    (n.next = r),
                                    a && ((n.method = "next"), (n.arg = e)),
                                    !!a
                                );
                            }
                            for (
                                var l = this.tryEntries.length - 1;
                                l >= 0;
                                --l
                            ) {
                                var i = this.tryEntries[l],
                                    o = i.completion;
                                if ("root" === i.tryLoc) return a("end");
                                if (i.tryLoc <= this.prev) {
                                    var s = r.call(i, "catchLoc"),
                                        u = r.call(i, "finallyLoc");
                                    if (s && u) {
                                        if (this.prev < i.catchLoc)
                                            return a(i.catchLoc, !0);
                                        if (this.prev < i.finallyLoc)
                                            return a(i.finallyLoc);
                                    } else if (s) {
                                        if (this.prev < i.catchLoc)
                                            return a(i.catchLoc, !0);
                                    } else {
                                        if (!u)
                                            throw new Error(
                                                "try statement without catch or finally",
                                            );
                                        if (this.prev < i.finallyLoc)
                                            return a(i.finallyLoc);
                                    }
                                }
                            }
                        },
                        abrupt: function (e, t) {
                            for (
                                var n = this.tryEntries.length - 1;
                                n >= 0;
                                --n
                            ) {
                                var a = this.tryEntries[n];
                                if (
                                    a.tryLoc <= this.prev &&
                                    r.call(a, "finallyLoc") &&
                                    this.prev < a.finallyLoc
                                ) {
                                    var l = a;
                                    break;
                                }
                            }
                            l &&
                                ("break" === e || "continue" === e) &&
                                l.tryLoc <= t &&
                                t <= l.finallyLoc &&
                                (l = null);
                            var i = l ? l.completion : {};
                            return (
                                (i.type = e),
                                (i.arg = t),
                                l
                                    ? ((this.method = "next"),
                                      (this.next = l.finallyLoc),
                                      g)
                                    : this.complete(i)
                            );
                        },
                        complete: function (e, t) {
                            if ("throw" === e.type) throw e.arg;
                            return (
                                "break" === e.type || "continue" === e.type
                                    ? (this.next = e.arg)
                                    : "return" === e.type
                                    ? ((this.rval = this.arg = e.arg),
                                      (this.method = "return"),
                                      (this.next = "end"))
                                    : "normal" === e.type &&
                                      t &&
                                      (this.next = t),
                                g
                            );
                        },
                        finish: function (e) {
                            for (
                                var t = this.tryEntries.length - 1;
                                t >= 0;
                                --t
                            ) {
                                var n = this.tryEntries[t];
                                if (n.finallyLoc === e)
                                    return (
                                        this.complete(n.completion, n.afterLoc),
                                        P(n),
                                        g
                                    );
                            }
                        },
                        catch: function (e) {
                            for (
                                var t = this.tryEntries.length - 1;
                                t >= 0;
                                --t
                            ) {
                                var n = this.tryEntries[t];
                                if (n.tryLoc === e) {
                                    var r = n.completion;
                                    if ("throw" === r.type) {
                                        var a = r.arg;
                                        P(n);
                                    }
                                    return a;
                                }
                            }
                            throw new Error("illegal catch attempt");
                        },
                        delegateYield: function (t, n, r) {
                            return (
                                (this.delegate = {
                                    iterator: z(t),
                                    resultName: n,
                                    nextLoc: r,
                                }),
                                "next" === this.method && (this.arg = e),
                                g
                            );
                        },
                    }),
                    t
                );
            }
            function sa(e, t, n, r, a, l, i) {
                try {
                    var o = e[l](i),
                        s = o.value;
                } catch (u) {
                    return void n(u);
                }
                o.done ? t(s) : Promise.resolve(s).then(r, a);
            }
            function ua(e, t, n, r, a) {
                return ca.apply(this, arguments);
            }
            function ca() {
                var e;
                return (
                    (e = oa().mark(function e(t, n, r, a, l) {
                        var i, o, s;
                        return oa().wrap(function (e) {
                            for (;;)
                                switch ((e.prev = e.next)) {
                                    case 0:
                                        if (
                                            ((i = t.pathname + t.search),
                                            -1 === kt(n, t.pathname))
                                        ) {
                                            e.next = 15;
                                            break;
                                        }
                                        return (
                                            a(function (e) {
                                                return e + 1;
                                            }),
                                            (e.next = 5),
                                            fetch("/api/v2".concat(i))
                                        );
                                    case 5:
                                        if ((o = e.sent).ok) {
                                            e.next = 10;
                                            break;
                                        }
                                        return (
                                            console.error(
                                                "Network response was not ok.",
                                            ),
                                            a(function (e) {
                                                return e - 1;
                                            }),
                                            e.abrupt("return")
                                        );
                                    case 10:
                                        return (e.next = 12), o.json();
                                    case 12:
                                        (s = e.sent),
                                            l() && r(s),
                                            a(function (e) {
                                                return e - 1;
                                            });
                                    case 15:
                                    case "end":
                                        return e.stop();
                                }
                        }, e);
                    })),
                    (ca = function () {
                        var t = this,
                            n = arguments;
                        return new Promise(function (r, a) {
                            var l = e.apply(t, n);
                            function i(e) {
                                sa(l, r, a, i, o, "next", e);
                            }
                            function o(e) {
                                sa(l, r, a, i, o, "throw", e);
                            }
                            i(void 0);
                        });
                    }),
                    ca.apply(this, arguments)
                );
            }
            var da = function () {
                    return (0, Ze.jsx)("div", {
                        className: "flex justify-center",
                        children: (0, Ze.jsx)(vt, { cls: "w-12 h-12 m-16" }),
                    });
                },
                fa = ["Profil", "Bek\xfcld\xe9sek", "Be\xe1ll\xedt\xe1sok"],
                pa = [_t.profile, _t.profileSubmissions, _t.profileSettings],
                ma = pa;
            var ha = function () {
                var e = je().user,
                    t = ye(),
                    n = c((0, r.useState)(null), 2),
                    a = n[0],
                    l = n[1],
                    i = c((0, r.useState)(0), 2),
                    o = i[0],
                    s = i[1],
                    u = pa.map(function (t) {
                        return t.replace(":user", e);
                    });
                (0, r.useEffect)(
                    function () {
                        var e = t.pathname + t.search;
                        -1 !== kt(ma, t.pathname) && ua(e, l, s);
                    },
                    [t],
                );
                var d = (0, Ze.jsx)(da, {});
                return (
                    0 === o && a && (d = (0, Ze.jsx)(Te, { data: a })),
                    (0, Ze.jsx)("div", {
                        className: "flex justify-center",
                        children: (0, Ze.jsx)("div", {
                            className: "w-full max-w-7xl",
                            children: (0, Ze.jsx)("div", {
                                className: "w-full px-4",
                                children: (0, Ze.jsx)(ia, {
                                    routes: u,
                                    routePatterns: pa,
                                    routeLabels: fa,
                                    children: (0, Ze.jsx)("div", {
                                        className: "w-full",
                                        children: d,
                                    }),
                                }),
                            }),
                        }),
                    })
                );
            };
            function va(e) {
                var t = e.tagName;
                return (0, Ze.jsx)("span", {
                    className:
                        "w-28 text-center truncate whitespace-nowrap cursor-pointer text-sm px-2 py-1 border-1 rounded bg-grey-725 hover:bg-indigo-600 border-grey-650 hover:border-indigo-500 transition-all duration-200",
                    children: t,
                });
            }
            var ga = function (e) {
                var t = e.title,
                    n = e.titleComponent,
                    r = e.tagNames.map(function (e, t) {
                        return (0, Ze.jsx)(
                            "div",
                            {
                                className: "flex m-1",
                                children: (0, Ze.jsx)(va, { tagName: e }, t),
                            },
                            t,
                        );
                    });
                return (0, Ze.jsx)(At, {
                    title: t,
                    titleComponent: n,
                    children: (0, Ze.jsx)("div", {
                        className:
                            "flex flex-col w-full overflow-x-auto rounded-md",
                        children: (0, Ze.jsx)("div", {
                            className: "flex flex-wrap p-4 bg-grey-850",
                            children: r,
                        }),
                    }),
                });
            };
            var ya = function () {
                var e = (0, Ze.jsx)(Ht, {
                        svg: (0, Ze.jsx)(jt, {
                            cls: "w-6 h-6 text-green-500 mr-2",
                        }),
                        title: "Megoldott feladatok",
                    }),
                    t = (0, Ze.jsx)(Ht, {
                        svg: (0, Ze.jsx)(wt, {
                            cls: "w-6 h-6 text-red-500 mr-2",
                        }),
                        title: "Megpr\xf3b\xe1lt feladatok",
                    });
                return (0, Ze.jsxs)("div", {
                    className: "flex flex-col sm:flex-row w-full items-start",
                    children: [
                        (0, Ze.jsxs)("div", {
                            className: "w-full sm:w-80 mb-3 shrink-0",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mb-3",
                                    children: (0, Ze.jsx)(Wt, {
                                        src: "/assets/profile.webp",
                                        username: "dbence",
                                        rating: 2350,
                                    }),
                                }),
                                (0, Ze.jsx)(Qt, {
                                    rating: 2350,
                                    score: 65.4,
                                    solved: 187,
                                }),
                            ],
                        }),
                        (0, Ze.jsxs)("div", {
                            className: "w-full mb-3 sm:ml-3",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mb-3",
                                    children: (0, Ze.jsx)(ga, {
                                        titleComponent: e,
                                        tagNames: [
                                            "KK23_tomjerry",
                                            "KK23_swaps",
                                            "KK23_speeding",
                                            "KK23_snacks",
                                            "KK23_rusco",
                                            "KK23_tomjerry",
                                            "KK23_swaps",
                                            "KK23_speeding",
                                            "KK23_snacks",
                                            "KK23_rusco",
                                        ],
                                    }),
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "mb-3",
                                    children: (0, Ze.jsx)(ga, {
                                        titleComponent: t,
                                        tagNames: [
                                            "KK23_tomjerry",
                                            "KK23_swaps",
                                            "KK23_speeding",
                                            "KK23_snacks",
                                            "KK23_rusco",
                                            "KK23_tomjerry",
                                            "KK23_swaps",
                                            "KK23_speeding",
                                            "KK23_snacks",
                                            "KK23_rusco",
                                        ],
                                    }),
                                }),
                            ],
                        }),
                    ],
                });
            };
            var ba = function (e) {
                var t = e.data;
                return t && G(_t.profileSubmissions, t.route)
                    ? (0, Ze.jsx)("div", {
                          className: "relative",
                          children: (0, Ze.jsxs)("div", {
                              className: "flex flex-col w-full",
                              children: [
                                  (0, Ze.jsx)("div", {
                                      className: "mb-2",
                                      children: (0, Ze.jsx)(Un, {
                                          submissions: t.submissions,
                                      }),
                                  }),
                                  (0, Ze.jsx)(In, {
                                      paginationData: t.paginationData,
                                  }),
                              ],
                          }),
                      })
                    : (0, Ze.jsx)(Ze.Fragment, {});
            };
            var xa = function (e) {
                var t = e.id,
                    n = e.label;
                return (0, Ze.jsxs)("label", {
                    htmlFor: t,
                    className: "flex items-start max-w-fit",
                    children: [
                        (0, Ze.jsx)("input", {
                            id: t,
                            className:
                                "appearance-none bg-grey-850 text-white border-1 border-default rounded w-5 h-5 shrink-0 checked:bg-indigo-600 checked:border-indigo-600 checkmark hover:bg-grey-800 hover:border-grey-600 checked:hover:bg-indigo-500 checked:hover:border-indigo-500 transition duration-200",
                            type: "checkbox",
                        }),
                        (0, Ze.jsx)("span", {
                            className: "text-label ml-3",
                            children: n,
                        }),
                    ],
                });
            };
            function wa() {
                var e = c((0, r.useState)(""), 2),
                    t = e[0],
                    n = e[1],
                    a = c((0, r.useState)(""), 2),
                    l = a[0],
                    i = a[1],
                    o = c((0, r.useState)(""), 2),
                    s = o[0],
                    u = o[1],
                    d = (0, Ze.jsx)(Ht, {
                        svg: (0, Ze.jsx)(yt, { cls: "w-6 h-6 mr-2" }),
                        title: "Jelsz\xf3 megv\xe1ltoztat\xe1sa",
                    });
                return (0, Ze.jsx)(At, {
                    titleComponent: d,
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "mb-4 w-full",
                                children: (0, Ze.jsx)(An, {
                                    id: "oldPassword",
                                    label: "R\xe9gi jelsz\xf3",
                                    type: "password",
                                    initText: t,
                                    onChange: function (e) {
                                        return n(e);
                                    },
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "mb-4 w-full",
                                children: (0, Ze.jsx)(An, {
                                    id: "newPassword",
                                    label: "\xdaj jelsz\xf3",
                                    type: "password",
                                    initText: l,
                                    onChange: function (e) {
                                        return i(e);
                                    },
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "mb-6 w-full",
                                children: (0, Ze.jsx)(An, {
                                    id: "newPasswordConfirm",
                                    label: "\xdaj jelsz\xf3 meger\u0151s\xedt\xe9se",
                                    type: "password",
                                    initText: s,
                                    onChange: function (e) {
                                        return u(e);
                                    },
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "flex justify-center",
                                children: (0, Ze.jsx)("button", {
                                    className: "btn-indigo w-32",
                                    children: "Ment\xe9s",
                                }),
                            }),
                        ],
                    }),
                });
            }
            function ja() {
                var e = (0, Ze.jsx)(Ht, {
                    svg: (0, Ze.jsx)(bt, { cls: "w-5 h-5 mr-2" }),
                    title: "Egy\xe9b be\xe1ll\xedt\xe1sok",
                });
                return (0, Ze.jsx)(At, {
                    titleComponent: e,
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "flex flex-col px-6 py-5 sm:px-10 sm:py-8 w-full",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "mb-2",
                                children: (0, Ze.jsx)(xa, {
                                    id: "showTagsUnsolved",
                                    label: "Megoldatlan feladatok c\xedmk\xe9inek mutat\xe1sa",
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "mb-6",
                                children: (0, Ze.jsx)(xa, {
                                    id: "hideSolved",
                                    label: "Megoldott feladatok elrejt\xe9se",
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "flex justify-center",
                                children: (0, Ze.jsx)("button", {
                                    className: "btn-indigo w-32",
                                    children: "Ment\xe9s",
                                }),
                            }),
                        ],
                    }),
                });
            }
            var ka = function () {
                    return (0, Ze.jsxs)("div", {
                        className:
                            "flex flex-col md:flex-row w-full items-start",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "w-full md:w-96 mb-3 shrink-0",
                                children: (0, Ze.jsx)(wa, {}),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "w-full mb-3 md:ml-3",
                                children: (0, Ze.jsx)(ja, {}),
                            }),
                        ],
                    });
                },
                Sa = [
                    "Le\xedr\xe1s",
                    "Bek\xfcld",
                    "Bek\xfcld\xe9sek",
                    "Eredm\xe9nyek",
                ],
                Na = [
                    _t.problem,
                    _t.problemSubmit,
                    _t.problemSubmissions,
                    _t.problemRanklist,
                ];
            var Ca = function () {
                    var e = je().problem,
                        t = Na.map(function (t) {
                            return t.replace(":problem", e);
                        });
                    return (0, Ze.jsx)("div", {
                        className: "flex justify-center",
                        children: (0, Ze.jsx)("div", {
                            className: "w-full max-w-7xl",
                            children: (0, Ze.jsx)("div", {
                                className: "w-full px-4",
                                children: (0, Ze.jsx)(ia, {
                                    routes: t,
                                    routeLabels: Sa,
                                    routePatterns: Na,
                                    children: (0, Ze.jsx)("div", {
                                        className: "w-full",
                                        children: (0, Ze.jsx)(Te, {}),
                                    }),
                                }),
                            }),
                        }),
                    });
                },
                Ea = {
                    file: (0, Ze.jsx)(Je, {}),
                    description: (0, Ze.jsx)(et, {}),
                };
            function La() {
                var e = (0, Ze.jsx)("div", {
                        className: "flex flex-wrap",
                        children: [
                            "oszd meg \xe9s uralkodj",
                            "dp",
                            "adatszerkezetek",
                            "bin\xe1ris keres\xe9s",
                        ].map(function (e, t) {
                            return (0, Ze.jsx)(
                                "span",
                                { className: "tag", children: e },
                                t,
                            );
                        }),
                    }),
                    t = (0, Ze.jsx)(Ht, {
                        svg: (0, Ze.jsx)(tt, {}),
                        title: "Inform\xe1ci\xf3k",
                    });
                return (0, Ze.jsx)($t, {
                    titleComponent: t,
                    data: [
                        ["Azonos\xedt\xf3", "OKTV23_Szivarvanyszamok"],
                        [
                            "C\xedm",
                            "Az \xf3vodai l\xe9t elviselhetetlen k\xf6nny\u0171s\xe9ge",
                        ],
                        ["Id\u0151limit", "300 ms"],
                        ["Mem\xf3rialimit", "31 MiB"],
                        ["C\xedmk\xe9k", e],
                        ["T\xedpus", "batch"],
                    ],
                });
            }
            function _a() {
                var e = (0, Ze.jsx)(Ht, {
                    svg: (0, Ze.jsx)(nt, {}),
                    title: "Megold\xe1s bek\xfcld\xe9se",
                });
                return (0, Ze.jsx)(At, {
                    titleComponent: e,
                    children: (0, Ze.jsx)("div", {
                        className: "px-6 py-5",
                        children: (0, Ze.jsxs)("div", {
                            className: "flex flex-col",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mb-4",
                                    children: (0, Ze.jsx)(Lt, {
                                        itemNames: [
                                            "C++ 11",
                                            "C++ 14",
                                            "C++ 17",
                                            "Go",
                                            "Java",
                                            "Python 3",
                                        ],
                                    }),
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "mb-2 mx-1 text-label",
                                    children: "Nincs kiv\xe1lasztva f\xe1jl.",
                                }),
                                (0, Ze.jsxs)("div", {
                                    className: "flex justify-center",
                                    children: [
                                        (0, Ze.jsx)("button", {
                                            className: "btn-gray w-1/2",
                                            children: "Tall\xf3z\xe1s",
                                        }),
                                        (0, Ze.jsx)("button", {
                                            className: "ml-2 btn-indigo w-1/2",
                                            children: "Bek\xfcld\xe9s",
                                        }),
                                    ],
                                }),
                            ],
                        }),
                    }),
                });
            }
            function Pa() {
                var e = [
                        ["minta.zip", "file"],
                        ["english", "description"],
                        ["hungarian", "description"],
                    ].map(function (e, t) {
                        var n =
                            "file" === e[1]
                                ? "F\xe1jl"
                                : "description" === e[1]
                                ? "Le\xedr\xe1s"
                                : "Csatolm\xe1ny";
                        return (0, Ze.jsxs)(
                            "li",
                            {
                                className:
                                    "flex items-center cursor-pointer text-indigo-200 hover:text-indigo-100 transition duration-200",
                                children: [
                                    Ea[e[1]],
                                    (0, Ze.jsxs)("span", {
                                        className: "underline",
                                        children: [n, " (", e[0], ")"],
                                    }),
                                ],
                            },
                            t,
                        );
                    }),
                    t = (0, Ze.jsx)(Ht, {
                        svg: (0, Ze.jsx)(rt, {}),
                        title: "Mell\xe9kletek",
                    });
                return (0, Ze.jsx)(At, {
                    titleComponent: t,
                    children: (0, Ze.jsx)("div", {
                        className: "px-6 py-5",
                        children: (0, Ze.jsx)("ul", { children: e }),
                    }),
                });
            }
            var Oa = function () {
                return (0, Ze.jsxs)("div", {
                    className: "flex flex-col lg:flex-row",
                    children: [
                        (0, Ze.jsx)("div", {
                            className: "w-full mb-3",
                            children: (0, Ze.jsx)("object", {
                                data: "/assets/statement.pdf",
                                type: "application/pdf",
                                width: "100%",
                                className: "h-[36rem] lg:h-[52rem]",
                            }),
                        }),
                        (0, Ze.jsxs)("div", {
                            className: "w-full lg:w-96 mb-3 lg:ml-3 shrink-0",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mb-3",
                                    children: (0, Ze.jsx)(La, {}),
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "mb-3",
                                    children: (0, Ze.jsx)(_a, {}),
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "mb-3",
                                    children: (0, Ze.jsx)(Pa, {}),
                                }),
                            ],
                        }),
                    ],
                });
            };
            function za() {
                return (0, Ze.jsx)(At, {
                    children: (0, Ze.jsxs)("div", {
                        className: "px-8 py-6 flex",
                        children: [
                            (0, Ze.jsx)(Lt, {
                                itemNames: [
                                    "C++ 11",
                                    "C++ 14",
                                    "C++ 17",
                                    "Go",
                                    "Java",
                                    "Python 3",
                                ],
                            }),
                            (0, Ze.jsx)("button", {
                                className: "ml-2 btn-indigo",
                                children: "Bek\xfcld\xe9s",
                            }),
                        ],
                    }),
                });
            }
            var Ma = function () {
                return (0, Ze.jsxs)("div", {
                    className: "flex flex-col",
                    children: [
                        (0, Ze.jsx)("div", {
                            className: "mb-2",
                            children: (0, Ze.jsx)(za, {}),
                        }),
                        (0, Ze.jsx)(ra, {
                            className: "border-1 border-default",
                            height: "60vh",
                            theme: "vs-dark",
                            defaultLanguage: "cpp",
                            options: { fontFamily: "JetBrains Mono" },
                        }),
                    ],
                });
            };
            function Ra() {
                return (0, Ze.jsx)(At, {
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "px-6 py-4 flex flex-col sm:flex-row items-start sm:items-center justify-between",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "mb-2 sm:mb-0",
                                children: (0, Ze.jsx)(xa, {
                                    label: "Teljes megold\xe1sok",
                                }),
                            }),
                            (0, Ze.jsx)(xa, {
                                label: "Saj\xe1t bek\xfcld\xe9seim",
                            }),
                        ],
                    }),
                });
            }
            var Ta = function (e) {
                var t = e.data;
                return t && G(_t.problemSubmissions, t.route)
                    ? (0, Ze.jsxs)("div", {
                          className: "relative",
                          children: [
                              (0, Ze.jsx)("div", {
                                  className: "mb-2",
                                  children: (0, Ze.jsx)(Ra, {}),
                              }),
                              (0, Ze.jsx)("div", {
                                  className: "mb-2",
                                  children: (0, Ze.jsx)(Un, {
                                      submissions: t.submissions,
                                  }),
                              }),
                              (0, Ze.jsx)(In, {
                                  paginationData: t.paginationData,
                              }),
                          ],
                      })
                    : (0, Ze.jsx)(Ze.Fragment, {});
            };
            function Fa(e) {
                var t = e.name,
                    n = e.score,
                    r = e.emphasize;
                return (0, Ze.jsxs)("tr", {
                    className: "divide-x divide-grey-700",
                    children: [
                        (0, Ze.jsx)("td", {
                            className: "padding-td-default bg-grey-800 ".concat(
                                r ? "font-medium" : "",
                                " align-top",
                            ),
                            children: (0, Ze.jsx)("span", {
                                className: "link",
                                children: t,
                            }),
                        }),
                        (0, Ze.jsx)("td", {
                            className:
                                "padding-td-default bg-grey-825 text-center whitespace-nowrap",
                            children: (0, Ze.jsx)("span", {
                                className: "link",
                                children: n,
                            }),
                        }),
                    ],
                });
            }
            var Ia = function (e) {
                var t = e.data,
                    n = e.title,
                    r = e.titleComponent,
                    a = e.emphasize;
                null == a && (a = !0);
                var l = t.map(function (e, t) {
                    return (0, Ze.jsx)(
                        Fa,
                        { name: e[0], score: e[1], emphasize: a },
                        t,
                    );
                });
                return (0, Ze.jsx)(Vt, {
                    title: n,
                    titleComponent: r,
                    children: (0, Ze.jsx)("tbody", {
                        className: "divide-y divide-default",
                        children: l,
                    }),
                });
            };
            var Da = function () {
                var e = (0, Ze.jsx)(Ht, {
                    svg: (0, Ze.jsx)(mt, {}),
                    title: "Eredm\xe9nyek",
                });
                return (0, Ze.jsxs)("div", {
                    children: [
                        (0, Ze.jsx)("div", {
                            className: "mb-2",
                            children: (0, Ze.jsx)(Ia, {
                                data: [
                                    ["dbence", "50 / 50", "5669"],
                                    ["dbence", "50 / 50", "5669"],
                                    ["vpeti", "48 / 50", "5669"],
                                    ["vpeti", "48 / 50", "5669"],
                                    ["gonterarmin", "2 / 50", "5669"],
                                    ["gonterarmin", "2 / 50", "5669"],
                                ],
                                titleComponent: e,
                                emphasize: !1,
                            }),
                        }),
                        (0, Ze.jsx)(In, { current: 20, last: 50 }),
                    ],
                });
            };
            function Ua() {
                var e = (0, Ze.jsx)(Ht, {
                    svg: (0, Ze.jsx)(ot, { cls: "w-[1.1rem] h-[1.1rem] mr-2" }),
                    title: "Bel\xe9p\xe9s",
                });
                return (0, Ze.jsx)(At, {
                    titleComponent: e,
                    children: (0, Ze.jsxs)("div", {
                        className: "px-10 py-8",
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "mb-4",
                                children: (0, Ze.jsx)(An, {
                                    id: "userName",
                                    label: "Felhaszn\xe1l\xf3n\xe9v",
                                }),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "mb-6",
                                children: (0, Ze.jsx)(An, {
                                    id: "password",
                                    label: "Jelsz\xf3",
                                    type: "password",
                                }),
                            }),
                            (0, Ze.jsxs)("div", {
                                className: "flex justify-center mb-2",
                                children: [
                                    (0, Ze.jsx)("button", {
                                        className: "btn-indigo mr-2 w-1/2",
                                        children: "Bel\xe9p\xe9s",
                                    }),
                                    (0, Ze.jsxs)("button", {
                                        className:
                                            "relative btn-gray flex items-center justify-between w-1/2",
                                        children: [
                                            (0, Ze.jsx)("div", {
                                                className:
                                                    "h-full flex items-center absolute left-2.5",
                                                children: (0, Ze.jsx)(ht, {}),
                                            }),
                                            (0, Ze.jsx)("div", {
                                                className:
                                                    "w-full flex justify-center",
                                                children: (0, Ze.jsx)("span", {
                                                    children: "Google",
                                                }),
                                            }),
                                        ],
                                    }),
                                ],
                            }),
                        ],
                    }),
                });
            }
            var Ba = function () {
                return (0, Ze.jsx)("div", {
                    className: "text-white",
                    children: (0, Ze.jsx)("div", {
                        className: "w-full flex justify-center",
                        children: (0, Ze.jsx)("div", {
                            className: "flex justify-center w-full sm:max-w-md",
                            children: (0, Ze.jsx)("div", {
                                className: "w-full px-4",
                                children: (0, Ze.jsx)(Ua, {}),
                            }),
                        }),
                    }),
                });
            };
            function Aa() {
                var e = (0, Ze.jsx)(Ht, {
                    svg: (0, Ze.jsx)(ot, { cls: "w-[1.1rem] h-[1.1rem] mr-2" }),
                    title: "Regisztr\xe1ci\xf3",
                });
                return (0, Ze.jsxs)(At, {
                    titleComponent: e,
                    children: [
                        (0, Ze.jsxs)("div", {
                            className:
                                "px-10 pt-8 pb-6 border-b border-default",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mb-4 relative",
                                    children: (0, Ze.jsx)(An, {
                                        id: "username",
                                        label: "Felhaszn\xe1l\xf3n\xe9v",
                                    }),
                                }),
                                (0, Ze.jsx)(An, {
                                    id: "email",
                                    label: "E-mail c\xedm",
                                }),
                            ],
                        }),
                        (0, Ze.jsxs)("div", {
                            className: "px-10 pt-4 pb-8",
                            children: [
                                (0, Ze.jsx)("div", {
                                    className: "mb-4",
                                    children: (0, Ze.jsx)(An, {
                                        id: "password",
                                        label: "Jelsz\xf3",
                                    }),
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "mb-6",
                                    children: (0, Ze.jsx)(An, {
                                        id: "passwordConfirm",
                                        label: "Jelsz\xf3 meger\u0151s\xedt\xe9se",
                                    }),
                                }),
                                (0, Ze.jsx)("div", {
                                    className: "flex justify-center",
                                    children: (0, Ze.jsx)("button", {
                                        className: "btn-indigo w-40",
                                        children: "Regisztr\xe1ci\xf3",
                                    }),
                                }),
                            ],
                        }),
                    ],
                });
            }
            var Va = function () {
                return (0, Ze.jsx)("div", {
                    className: "text-white",
                    children: (0, Ze.jsx)("div", {
                        className: "w-full flex justify-center",
                        children: (0, Ze.jsx)("div", {
                            className: "flex justify-center w-full sm:max-w-md",
                            children: (0, Ze.jsx)("div", {
                                className: "w-full px-4",
                                children: (0, Ze.jsx)(Aa, {}),
                            }),
                        }),
                    }),
                });
            };
            function $a() {
                return (0, Ze.jsx)(At, {
                    title: "Az oldal nem el\xe9rhet\u0151",
                    children: (0, Ze.jsxs)("div", {
                        className:
                            "px-10 py-8 flex flex-col relative justify-between",
                        children: [
                            (0, Ze.jsx)("p", {
                                className: "z-10",
                                children:
                                    "A keresett oldal nem tal\xe1lhat\xf3. Gy\u0151z\u0151dj meg r\xf3la, hogy a megadott link helyes.",
                            }),
                            (0, Ze.jsx)("div", {
                                className:
                                    "flex justify-center absolute inset-0",
                                children: (0, Ze.jsx)(dt, {}),
                            }),
                            (0, Ze.jsx)("div", {
                                className: "flex justify-center mt-8",
                                children: (0, Ze.jsx)(Ge, {
                                    className:
                                        "z-10 w-60 btn-indigo text-center",
                                    to: "/",
                                    children: "Vissza a f\u0151oldalra",
                                }),
                            }),
                        ],
                    }),
                });
            }
            var Ha = function () {
                    return (0, Ze.jsx)("div", {
                        className: "w-full flex justify-center",
                        children: (0, Ze.jsx)("div", {
                            className:
                                "flex justify-center w-full max-w-md px-4",
                            children: (0, Ze.jsx)($a, {}),
                        }),
                    });
                },
                Ka = [
                    _t.main,
                    _t.contests,
                    _t.info,
                    _t.archive,
                    _t.submissions,
                    _t.problems,
                    _t.submission,
                    _t.login,
                    _t.register,
                ];
            var Wa = function () {
                var e = ye(),
                    t = c((0, r.useState)(null), 2),
                    n = t[0],
                    a = t[1],
                    l = c((0, r.useState)(0), 2),
                    i = l[0],
                    o = l[1];
                (0, r.useEffect)(
                    function () {
                        var t = !0;
                        return (
                            ua(e, Ka, a, o, function () {
                                return t;
                            }),
                            function () {
                                t = !1;
                            }
                        );
                    },
                    [e],
                );
                var s = (0, Ze.jsx)(da, {}),
                    u = null;
                return (
                    0 === i &&
                        n &&
                        ((s = null),
                        (u = (0, Ze.jsxs)(De, {
                            children: [
                                (0, Ze.jsx)(Fe, {
                                    path: _t.main,
                                    element: (0, Ze.jsx)(Zt, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.contests,
                                    element: (0, Ze.jsx)(tn, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.info,
                                    element: (0, Ze.jsx)(an, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.archive,
                                    element: (0, Ze.jsx)(sn, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.submissions,
                                    element: (0, Ze.jsx)(Bn, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.problems,
                                    element: (0, Ze.jsx)(Xn, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.submission,
                                    element: (0, Ze.jsx)(aa, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.login,
                                    element: (0, Ze.jsx)(Ba, { data: n }),
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: _t.register,
                                    element: (0, Ze.jsx)(Va, { data: n }),
                                }),
                                (0, Ze.jsxs)(Fe, {
                                    path: _t.profile,
                                    element: (0, Ze.jsx)(ha, {}),
                                    children: [
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.profile,
                                            element: (0, Ze.jsx)(ya, {}),
                                        }),
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.profileSubmissions,
                                            element: (0, Ze.jsx)(ba, {}),
                                        }),
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.profileSettings,
                                            element: (0, Ze.jsx)(ka, {}),
                                        }),
                                    ],
                                }),
                                (0, Ze.jsxs)(Fe, {
                                    path: _t.problem,
                                    element: (0, Ze.jsx)(Ca, {}),
                                    children: [
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.problem,
                                            element: (0, Ze.jsx)(Oa, {}),
                                        }),
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.problemSubmit,
                                            element: (0, Ze.jsx)(Ma, {}),
                                        }),
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.problemSubmissions,
                                            element: (0, Ze.jsx)(Ta, {}),
                                        }),
                                        (0, Ze.jsx)(Fe, {
                                            path: _t.problemRanklist,
                                            element: (0, Ze.jsx)(Da, {}),
                                        }),
                                    ],
                                }),
                                (0, Ze.jsx)(Fe, {
                                    path: "*",
                                    element: (0, Ze.jsx)(Ha, {}),
                                }),
                            ],
                        }))),
                    (0, Ze.jsxs)(Ze.Fragment, {
                        children: [
                            (0, Ze.jsx)("div", {
                                className: "pb-20",
                                children: (0, Ze.jsx)(Bt, {}),
                            }),
                            (0, Ze.jsx)("div", {
                                className:
                                    "transition-opacity ease-linear duration-250 ".concat(
                                        0 === i ? "opacity-0" : "opacity-100",
                                    ),
                                children: s,
                            }),
                            (0, Ze.jsx)("div", {
                                className:
                                    "transition-opacity ease-linear duration-250 ".concat(
                                        0 === i ? "opacity-100" : "opacity-0",
                                    ),
                                children: u,
                            }),
                        ],
                    })
                );
            };
            var Qa = function () {
                return (0, Ze.jsx)("div", {
                    className: "text-white h-full min-h-screen pb-4",
                    children: (0, Ze.jsx)(We, {
                        children: (0, Ze.jsx)(Wa, {}),
                    }),
                });
            };
            l.createRoot(document.getElementById("root")).render(
                (0, Ze.jsx)(Qa, {}),
            );
        })();
})();
//# sourceMappingURL=main.ec9516f8.js.map
