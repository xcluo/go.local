#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>

"""
/learning_design
    /plan
        /lists
        /detail
        /create
        /edit
        /release
    /course
        /lists
    /project
        /lists
    /learning_publish
        /lists
        /detail
        /publish_setup
        /publish
/learning_market
    /lists      # Default
    /detail     # Default
    /favorites  # Default
    /add_favorite
    /to_buy
/team_management
    /lists
    /detail
    /edit
"""


PERMISSIONS = {
    'learning_center': {
        'alias': '课程中心',
        'description': '',
        'subs': {
            'learn_design': {
                'alias': '课程设计',
                'description': '',
                'subs': {
                    'lists': {
                        'alias': '列出',
                        'description': ''
                    },
                    'detail': {
                        'alias': '详情',
                        'description': ''
                    },
                    'create': {
                        'alias': '创建',
                        'description': ''
                    },
                    'update': {
                        'alias': '更新',
                        'description': ''
                    },
                    'edit': {
                        'alias': '编辑',
                        'description': ''
                    },
                    'release': {
                        'alias': '版本',
                        'description': ''
                    },
                }
            },
            'my_release': {
                'alias': '我的发布',
                'description': '',
                'subs': {
                    'lists': {
                        'alias': '列出',
                    },
                    'detail': {
                        'alias': '详情',
                    },
                    'publish': {
                        'alias': '发布',
                    }
                }
            },
            'market_learning': {
                'alias': '课程市场',
                'subs': {
                    # lists               #
                    # detail              #
                    'buy': {
                        'alias': '购买'
                    }
                }
            },
            'market_favorite': {
                'alias': '我的收藏',
                'subs': {
                    'lists': {
                        'alias': '列出',
                    },
                    'remove': {
                        'alias': '移除',
                    }
                }
            }
        }
    },

    'user_support': {
        'alias': '用户支持',
        'subs': {
            'learning_qa': {
                'alias': '学习问答',
                'subs': {
                    'lists': {
                        'alias': '列出',
                    },
                    'detail': {
                        'alias': '详情',
                    },
                    'accept': {
                        'alias': '解答',
                    },
                }
            }
        }
    },

    'reports_statistic': {
        'alias': '报表统计',
        'subs': {
            'reports': {
                'alias': '报表'
            },
            'statistics': {
                'alias': '统计'
            }
        }
    },

    'team_management': {
        'alias': '团队管理',
        'subs': {
            'lists': {
                'alias': '列出',
            },
            'detail': {
                'alias': '详情',
            },
            'edit': {
                'alias': '编辑',
            },
            'create': {
                'alias': '创建',
            }
        }
    },

    # persona_center

    'system_management': {
        'alias': '系统管理'
    }
}


if __name__ == '__main__':
    for k1, v1 in PERMISSIONS.items():
        print 'vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv'
        if 'subs' in v1 and v1['subs']:
            for k2, v2 in v1['subs'].items():
                if 'subs' in v2 and v2['subs']:
                    for k3, v3 in v2['subs'].items():
                        if 'subs' in v3 and v3['subs']:
                            for k4, v4 in v3['subs'].items():
                                if 'subs' in v4 and v4['subs']:
                                    for k5, v5 in v4['subs'].items():
                                        if 'subs' in v5 and v5['subs']:
                                            pass
                                        else:
                                            print "%s:%s:%s"% (v1.get('alias'), v2.get('alias'), v3.get('alias'), v4.get('alias'), v5.get('alias')),
                                            print '\t\t' + k1 + ':' + k2 + ':' + k3 + ':' + k4 + ':' + k5
                                else:
                                    print "%s:%s:%s"% (v1.get('alias'), v2.get('alias'), v3.get('alias'), v4.get('alias')),
                                    print '\t\t' + k1 + ':' + k2 + ':' + k3 + ':' + k4
                        else:
                            print "%s:%s:%s"% (v1.get('alias'), v2.get('alias'), v3.get('alias')),
                            print '\t\t' + k1 + ':' + k2 + ':' + k3
                else:
                    print "%s:%s"% (v1.get('alias'), v2.get('alias')),
                    print '\t\t' + k1 + ':' + k2
                print '-------------------------------------------------------------------------'
        else:
            print v1.get('alias'),
            print '\t\t' + k1
        print '\n'
