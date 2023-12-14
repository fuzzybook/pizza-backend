export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Any: { input: any; output: any; }
  Map: { input: any; output: any; }
  Time: { input: any; output: any; }
  Upload: { input: any; output: any; }
};

export type DeleteCategory = {
  id: Scalars['Int']['input'];
};

export type DeleteCondiment = {
  id: Scalars['Int']['input'];
};

export type DeleteDough = {
  id: Scalars['Int']['input'];
};

export type DeleteIngredient = {
  id: Scalars['Int']['input'];
};

/** The `File` type, represents the response of uploading a file. */
export type File = {
  __typename?: 'File';
  content: Scalars['String']['output'];
  contentType: Scalars['String']['output'];
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type Language = {
  __typename?: 'Language';
  flag: Scalars['String']['output'];
  label: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export type LogoutResult = {
  __typename?: 'LogoutResult';
  ok?: Maybe<Scalars['Boolean']['output']>;
};

export type Menu = {
  __typename?: 'Menu';
  Categories: Array<MenuCategory>;
  Condiments: Array<MenuCondiment>;
  Doughs: Array<MenuDough>;
  Ingredients: Array<MenuIngredient>;
};

export type MenuCategory = {
  __typename?: 'MenuCategory';
  category: Scalars['String']['output'];
  id: Scalars['Int']['output'];
  image: Scalars['String']['output'];
  items: Array<MenuItem>;
  title: Scalars['String']['output'];
  uuid: Scalars['String']['output'];
};

export type MenuCondiment = {
  __typename?: 'MenuCondiment';
  available: Scalars['Int']['output'];
  categories: Array<MenuDataCategory>;
  id: Scalars['Int']['output'];
  priority: Scalars['Int']['output'];
  text: Scalars['String']['output'];
};

export type MenuDataCategory = {
  __typename?: 'MenuDataCategory';
  priority: Scalars['Int']['output'];
  text: Scalars['String']['output'];
};

export type MenuDough = {
  __typename?: 'MenuDough';
  available: Scalars['Int']['output'];
  categories: Array<MenuDataCategory>;
  id: Scalars['Int']['output'];
  priority: Scalars['Int']['output'];
  text: Scalars['String']['output'];
};

export type MenuIngredient = {
  __typename?: 'MenuIngredient';
  available: Scalars['Int']['output'];
  categories: Array<MenuDataCategory>;
  id: Scalars['Int']['output'];
  priority: Scalars['Int']['output'];
  text: Scalars['String']['output'];
};

export type MenuItem = {
  __typename?: 'MenuItem';
  categoryRefer: Scalars['Int']['output'];
  data: MenuItemData;
  dirty: Scalars['Boolean']['output'];
  id: Scalars['Int']['output'];
  promo: Scalars['Boolean']['output'];
  uuid: Scalars['String']['output'];
};

export type MenuItemData = {
  __typename?: 'MenuItemData';
  condiments?: Maybe<Array<Scalars['String']['output']>>;
  doughs?: Maybe<Array<Scalars['String']['output']>>;
  extra?: Maybe<Array<Maybe<MenuExtra>>>;
  image: Scalars['String']['output'];
  ingredients?: Maybe<Array<Scalars['String']['output']>>;
  price: Scalars['String']['output'];
  text: Scalars['String']['output'];
  title: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createUser: User;
  deleteMenuCategory: Scalars['Boolean']['output'];
  deleteMenuCategoryItem: Scalars['Boolean']['output'];
  deleteMenuCondiment: Scalars['Boolean']['output'];
  deleteMenuDough: Scalars['Boolean']['output'];
  deleteMenuIngredient: Scalars['Boolean']['output'];
  login?: Maybe<Session>;
  saveMenuImage?: Maybe<Scalars['String']['output']>;
  updateMenuCategory: MenuCategory;
  updateMenuCategoryItem: MenuItem;
  updateMenuCondiment: MenuCondiment;
  updateMenuCondiments: Array<MenuCondiment>;
  updateMenuDough: MenuDough;
  updateMenuDoughs: Array<MenuDough>;
  updateMenuIngredient: MenuIngredient;
  updateMenuIngredients: Array<MenuIngredient>;
  updateMenuItemPromo: Scalars['Boolean']['output'];
  updateUserPassword: User;
  updateUserRoles: User;
  uploadMenuImage: Scalars['Boolean']['output'];
};


export type MutationCreateUserArgs = {
  input: NewUser;
};


export type MutationDeleteMenuCategoryArgs = {
  input: DeleteCategory;
};


export type MutationDeleteMenuCategoryItemArgs = {
  input?: InputMaybe<UpdateCategoryItem>;
};


export type MutationDeleteMenuCondimentArgs = {
  input: DeleteCondiment;
};


export type MutationDeleteMenuDoughArgs = {
  input: DeleteDough;
};


export type MutationDeleteMenuIngredientArgs = {
  input: DeleteIngredient;
};


export type MutationLoginArgs = {
  input: UserLogin;
};


export type MutationSaveMenuImageArgs = {
  input: SaveImage;
};


export type MutationUpdateMenuCategoryArgs = {
  input?: InputMaybe<UpdateCategory>;
};


export type MutationUpdateMenuCategoryItemArgs = {
  input?: InputMaybe<UpdateCategoryItem>;
};


export type MutationUpdateMenuCondimentArgs = {
  input: UpdateCondiment;
};


export type MutationUpdateMenuCondimentsArgs = {
  input?: InputMaybe<Array<InputMaybe<UpdateCondiment>>>;
};


export type MutationUpdateMenuDoughArgs = {
  input: UpdateDough;
};


export type MutationUpdateMenuDoughsArgs = {
  input?: InputMaybe<Array<InputMaybe<UpdateDough>>>;
};


export type MutationUpdateMenuIngredientArgs = {
  input: UpdateIngredient;
};


export type MutationUpdateMenuIngredientsArgs = {
  input?: InputMaybe<Array<InputMaybe<UpdateIngredient>>>;
};


export type MutationUpdateMenuItemPromoArgs = {
  id: Scalars['Int']['input'];
  promo: Scalars['Boolean']['input'];
};


export type MutationUpdateUserPasswordArgs = {
  input: UpdateUserPassword;
};


export type MutationUpdateUserRolesArgs = {
  input: UpdateUserRoles;
};


export type MutationUploadMenuImageArgs = {
  file: Scalars['Upload']['input'];
};

export type NewUser = {
  email: Scalars['String']['input'];
  name: Scalars['String']['input'];
  password: Scalars['String']['input'];
  roles: Array<UserRole>;
  type: UserType;
};

export type Query = {
  __typename?: 'Query';
  getUser: User;
  logout?: Maybe<LogoutResult>;
  me: User;
  menu: Menu;
  sessions: Array<Session>;
  system: SystemInfo;
  users: Array<User>;
  usersPaginate: UserPagesResponse;
};


export type QueryGetUserArgs = {
  userId: Scalars['ID']['input'];
};


export type QueryUsersPaginateArgs = {
  input: UserPages;
};

export type SaveImage = {
  id: Scalars['Int']['input'];
  image: Scalars['String']['input'];
  uuid: Scalars['String']['input'];
};

export type Session = {
  __typename?: 'Session';
  firedAt?: Maybe<Scalars['Time']['output']>;
  recoveryToken: Scalars['String']['output'];
  user: User;
};

export type SystemInfo = {
  __typename?: 'SystemInfo';
  avatars: Scalars['String']['output'];
  languages: Array<Language>;
  roles: Array<Scalars['String']['output']>;
  version: Scalars['String']['output'];
};

export type UpdateCategory = {
  category: Scalars['String']['input'];
  id: Scalars['Int']['input'];
  image: Scalars['String']['input'];
  title: Scalars['String']['input'];
};

export type UpdateCategoryItem = {
  categoryRefer: Scalars['Int']['input'];
  data: Scalars['String']['input'];
  id: Scalars['Int']['input'];
};

export type UpdateCondiment = {
  available: Scalars['Int']['input'];
  categories: Scalars['String']['input'];
  id: Scalars['Int']['input'];
  priority: Scalars['Int']['input'];
  text: Scalars['String']['input'];
};

export type UpdateDough = {
  available: Scalars['Int']['input'];
  categories: Scalars['String']['input'];
  id: Scalars['Int']['input'];
  priority: Scalars['Int']['input'];
  text: Scalars['String']['input'];
};

export type UpdateIngredient = {
  available: Scalars['Int']['input'];
  categories: Scalars['String']['input'];
  id: Scalars['Int']['input'];
  priority: Scalars['Int']['input'];
  text: Scalars['String']['input'];
};

export type UpdateUserPassword = {
  password: Scalars['String']['input'];
  userId: Scalars['Int']['input'];
};

export type UpdateUserRoles = {
  roles: Array<UserRole>;
  userId: Scalars['Int']['input'];
};

/** The `UploadFile` type, represents the request for uploading a file with certain payload. */
export type UploadFile = {
  file: Scalars['Upload']['input'];
  id: Scalars['Int']['input'];
};

export type User = {
  __typename?: 'User';
  activatedAt?: Maybe<Scalars['Time']['output']>;
  avatar?: Maybe<Scalars['String']['output']>;
  details?: Maybe<UserDetails>;
  email: Scalars['String']['output'];
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
  preferences?: Maybe<UserPreferences>;
  roles: Array<UserRole>;
  status: UserStatus;
  types: UserType;
  uuid?: Maybe<Scalars['String']['output']>;
};

export type UserDetails = {
  __typename?: 'UserDetails';
  address?: Maybe<Scalars['String']['output']>;
  city?: Maybe<Scalars['String']['output']>;
  country?: Maybe<Scalars['String']['output']>;
  firstName: Scalars['String']['output'];
  id: Scalars['Int']['output'];
  lastName: Scalars['String']['output'];
  phone?: Maybe<Scalars['String']['output']>;
  title: Scalars['String']['output'];
  zipCode?: Maybe<Scalars['String']['output']>;
};

export type UserLogin = {
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
};

export type UserPages = {
  descending?: InputMaybe<Scalars['Boolean']['input']>;
  page: Scalars['Int']['input'];
  pageSize: Scalars['Int']['input'];
  sortBy?: InputMaybe<Scalars['String']['input']>;
};

export type UserPagesResponse = {
  __typename?: 'UserPagesResponse';
  count: Scalars['Int']['output'];
  page: Scalars['Int']['output'];
  pageSize: Scalars['Int']['output'];
  users?: Maybe<Array<Maybe<User>>>;
};

export type UserPreferences = {
  __typename?: 'UserPreferences';
  idlePin?: Maybe<Scalars['String']['output']>;
  idleTimeout?: Maybe<Scalars['Int']['output']>;
  language?: Maybe<Scalars['String']['output']>;
  sendNoticesMail?: Maybe<Scalars['Boolean']['output']>;
  useDirectLogin?: Maybe<Scalars['Boolean']['output']>;
  useIdle?: Maybe<Scalars['Boolean']['output']>;
  useIdlePassword?: Maybe<Scalars['Boolean']['output']>;
  useQuadcodeLogin?: Maybe<Scalars['Boolean']['output']>;
};

export enum UserRole {
  Admin = 'ADMIN',
  Musicmanager = 'MUSICMANAGER',
  User = 'USER'
}

export type UserSocials = {
  __typename?: 'UserSocials';
  facebook?: Maybe<Scalars['String']['output']>;
  instagram?: Maybe<Scalars['String']['output']>;
  linkedin?: Maybe<Scalars['String']['output']>;
  twitter?: Maybe<Scalars['String']['output']>;
  website?: Maybe<Scalars['String']['output']>;
  youtube?: Maybe<Scalars['String']['output']>;
};

export enum UserStatus {
  Active = 'ACTIVE',
  Awaiting = 'AWAITING',
  Banned = 'BANNED',
  Blocked = 'BLOCKED',
  Pending = 'PENDING'
}

export enum UserType {
  Author = 'AUTHOR',
  Site = 'SITE',
  System = 'SYSTEM'
}

export type MenuExtra = {
  __typename?: 'menuExtra';
  alternative?: Maybe<Scalars['String']['output']>;
  price: Scalars['String']['output'];
  take: Scalars['Boolean']['output'];
  text: Scalars['String']['output'];
};
