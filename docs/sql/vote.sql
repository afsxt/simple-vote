SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for vote_theme
-- ----------------------------
DROP TABLE IF EXISTS `vote_theme`;
CREATE TABLE `vote_theme` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) UNIQUE COMMENT '主题名称',
  `description` varchar(255) DEFAULT '' COMMENT '简述',
  `flag` tinyint(3) unsigned DEFAULT '0' COMMENT '选举是否开始标记',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='选举主题';

-- ----------------------------
-- Table structure for vote_candidates
-- ----------------------------
DROP TABLE IF EXISTS `vote_candidates`;
CREATE TABLE `vote_candidates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COMMENT '候选人名称',
  `description` varchar(255) DEFAULT '' COMMENT '简述',
  `theme_id` int(10) unsigned COMMENT '选举主题id',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  UNIQUE KEY(`name`, `theme_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='候选人';

-- ----------------------------
-- Table structure for vote_user
-- ----------------------------
DROP TABLE IF EXISTS `vote_user`;
CREATE TABLE `vote_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(100) COMMENT '用户邮箱',
  `id_card` varchar(255) UNIQUE COMMENT '身份证号',
  `verify` tinyint(3) unsigned DEFAULT '0' COMMENT '是否通过较验',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='普通用户';

-- ----------------------------
-- Table structure for vote_vote
-- ----------------------------
DROP TABLE IF EXISTS `vote_vote`;
CREATE TABLE `vote_vote` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `theme_id` int(10) unsigned NOT NULL COMMENT '主题id',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `candidate_id` int(10) unsigned NOT NULL COMMENT '候选人id',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='选举表';