import type { Language } from "../models/Settings";

export const Persian: Language = {
  Errors: {
    LogoutFailed: ".خروج ناموفق بود",
    FileDeleteFailed: "حذف فایل ناموفق بود.",
    FileDownloadFailed: "دانلود فایل ناموفق بود.",
    FileShareFailed: "اشتراک گذاری فایل ناموفق بود.",
    FileUploadFailed: "بارگذاری فایل ناموفق بود.",
    UserListFetchFailed: "دریافت لیست کاربران ناموفق بود.",
    UserFetchFailed: "دریافت اطلاعات کاربر ناموفق بود.",
    UserFilesFetchFailed: "دریافت فایل های کاربر ناموفق بود.",
    UserExists: "کاربری با همین ایمیل وجود دارد.",
    UsernameUpdateFailed: "تغییر نام کاربری ناموفق بود.",
    WrongEmailOrPassword: "ایمیل یا رمز عبور اشتباه است.",
    UnknownError: "خطای ناشناخته",
    CheckCredentials: "مشخصات خود را بررسی کنید.",
  },
  Warnings: {
    NoUserSelected: "هیچ کاربری انتخاب نکردید.",
    NoFileSelected: "هیچ فایلی انتخاب نکردید.",
    CheckRegisterInfo: "مشخصات عضویت خود را بررسی کنید.",
  },
  Success: {
    SuccessfulLogout: "خروج موفق",
    FileDeleted: "فایل حذف شد.",
    FileShared: "فایل به اشتراک گذاشته شد.",
    FileUploaded: "فایل بارگذاری شد.",
    UsernameUpdated: "نام کاربری تغییر داده شد.",
  },
  Strings: {
    Logout: "خروج",
    Delete: "حذف",
    AreYouSure: "آیا اطمینان دارید؟",
    Yes: "بله",
    No: "خیر",
    Share: "اشتراک گذاری",
    RecentlyShared: "اخیرا به اشتراک گذاشته شده",
    Selected: "انتخاب شد",
    FileAlreadySharedWithEveryone: "این فایل قبلا با همه به اشتراک گذاشته شده.",
    Upload: "بارگذاری",
    Username: "نام کاربری",
    Email: "ایمیل",
    Loading: "در حال بارگیری...",
    NothingToSeeHere: "چیزی برای دیدن موجود نیست.",
    FilesSharedWithYou: "فایل های به اشتراک گذاشته شده با شما",
    NoOneHasSharedAnyFileWithYou: "کسی فایلی را با شما به اشتراک نگذاشته.",
    WelcomeToNimbusCloud: "به فضای ابری نیمبوس خوش آمدید.",
    Register: "عضویت",
    Login: "ورود",
    OpenDashboard: "باز کردن داشبورد",
    DontHaveAnAccountYet: "آیا هنوز عضو نشده اید؟ عضویت",
    AlreadyHaveAnAccount: "آیا قبلا عضو شده اید؟ ورود",
  },
};

export const English: Language = {
  Errors: {
    LogoutFailed: "Logout failed",
    FileDeleteFailed: "Failed to delete the file.",
    FileDownloadFailed: "Failed to download the file.",
    FileShareFailed: "Failed to share the file.",
    FileUploadFailed: "Failed to upload the file.",
    UserListFetchFailed: "Failed to fetch users list.",
    UserFetchFailed: "Failed to fetch user info.",
    UserFilesFetchFailed: "Failed to fetch user files.",
    UserExists: "User with the same email already exists.",
    UsernameUpdateFailed: "Failed to update the username.",
    WrongEmailOrPassword: "Wrong email or password",
    UnknownError: "Unknown error",
    CheckCredentials: "Check your credentials.",
  },
  Warnings: {
    NoUserSelected: "No user selected.",
    NoFileSelected: "No file selected.",
    CheckRegisterInfo: "Check you register info.",
  },
  Success: {
    SuccessfulLogout: "Successful logout",
    FileDeleted: "File deleted.",
    FileShared: "File shared.",
    FileUploaded: "File uploaded.",
    UsernameUpdated: "Username updated.",
  },
  Strings: {
    Logout: "Logout",
    Delete: "Delete",
    AreYouSure: "Are you sure?",
    Yes: "Yes",
    No: "No",
    Share: "Share",
    RecentlyShared: "Recently Shared",
    Selected: "Selected",
    FileAlreadySharedWithEveryone: "This file is shared with everyone already.",
    Upload: "Upload",
    Username: "Username",
    Email: "Email",
    Loading: "Loading...",
    NothingToSeeHere: "Nothing to see here.",
    FilesSharedWithYou: "Files shared with you",
    NoOneHasSharedAnyFileWithYou: "No one has shared any file with you.",
    WelcomeToNimbusCloud: "Welcome to Nimbus Cloud",
    Register: "Register",
    Login: "Login",
    OpenDashboard: "Open dashboard",
    DontHaveAnAccountYet: "Don't have an account? Register",
    AlreadyHaveAnAccount: "Already have an account? Login",
  },
};
