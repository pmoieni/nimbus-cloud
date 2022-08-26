export interface Language {
  Errors: {
    LogoutFailed: string;
    FileDeleteFailed: string;
    FileDownloadFailed: string;
    FileShareFailed: string;
    FileUploadFailed: string;
    UserListFetchFailed: string;
    UserFetchFailed: string;
    UserFilesFetchFailed: string;
    UserExists: string;
    UsernameUpdateFailed: string;
    WrongEmailOrPassword: string;
    UnknownError: string;
    CheckCredentials: string;
  };
  Warnings: {
    NoUserSelected: string;
    NoFileSelected: string;
    CheckRegisterInfo: string;
  };
  Success: {
    SuccessfulLogout: string;
    FileDeleted: string;
    FileShared: string;
    FileUploaded: string;
    UsernameUpdated: string;
  };
  Strings: {
    Logout: string;
    Delete: string;
    AreYouSure: string;
    Yes: string;
    No: string;
    Share: string;
    RecentlyShared: string;
    Selected: string;
    FileAlreadySharedWithEveryone: string;
    Upload: string;
    Username: string;
    Email: string;
    Loading: string;
    NothingToSeeHere: string;
    FilesSharedWithYou: string;
    NoOneHasSharedAnyFileWithYou: string;
    WelcomeToNimbusCloud: string;
    Register: string;
    Login: string;
    OpenDashboard: string;
    DontHaveAnAccountYet: string;
    AlreadyHaveAnAccount: string;
  };
}
