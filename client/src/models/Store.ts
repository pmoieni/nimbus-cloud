export interface File {
  id: number;
  name: string;
  object_name: string;
  owner: number;
}

export interface FileList {
  userFiles: File[];
  permittedFiles: File[];
}

export interface FileShareReq {
  users: string[];
}
