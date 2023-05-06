import { observer } from 'mobx-react-lite';
import Image from 'next/image';
import { FC } from 'react';

import { currentChat } from '@/store/CurrentChat';
import noAvatarWhite from './assets/no-avatar-white.png';
import { ChatItemProps } from './models';
import styles from './styles.module.scss';

const ChatItem: FC<ChatItemProps> = ({ companionAvatar, companionName, id, lastMessage }) => {
  return (
    <li
      onClick={() => currentChat.setCurrentChat(id)}
      className={`${styles.chatItem} ${currentChat.id === id ? "color-bg hover:color-bg " : "hover:chats-bg-hover"}`}>
      <Image
        src={noAvatarWhite}
        alt='avatar'
        className="rounded-[50%] h-[45px] min-h-[45px] w-[45px] min-w-[45px] "
      />
      <div className="h-full flex w-full flex-col justify-around ">
        <span className="font-bold">
          {companionName}
        </span>
        <span>
          {lastMessage}
        </span>
      </div>
    </li>
  )
}

export default observer(ChatItem);