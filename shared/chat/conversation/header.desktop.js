// @flow
import React from 'react'
import {Box, Icon, Usernames} from '../../common-adapters'
import {globalStyles, globalColors, globalMargins} from '../../styles'
import {participantFilter, usernamesToUserListItem} from '../../constants/chat'

import type {Props} from './header'

const Header = ({participants, onOpenFolder, onToggleSidePanel, sidePanelOpen, you, metaDataMap, followingMap, onShowProfile}: Props) => (
  <Box style={containerStyle}>
    <Usernames colorFollowing={true} inline={false} commaColor={globalColors.black_40} type='BodyBig' users={usernamesToUserListItem(participantFilter(participants, you).toArray(), you, metaDataMap, followingMap)}
      containerStyle={{flex: 1, textAlign: 'center', justifyContent: 'center'}} onUsernameClicked={onShowProfile} />
    <Icon type='iconfont-folder-private' style={{marginLeft: globalMargins.tiny}} onClick={onOpenFolder} />
    <Icon type={sidePanelOpen ? 'iconfont-close' : 'iconfont-info'} style={{marginLeft: globalMargins.tiny}} onClick={onToggleSidePanel} />
  </Box>
)

const containerStyle = {
  ...globalStyles.flexBoxRow,
  minHeight: 32,
  borderBottom: `solid 1px ${globalColors.black_05}`,
  justifyContent: 'center',
  alignItems: 'center',
  padding: globalMargins.tiny,
}

export default Header
